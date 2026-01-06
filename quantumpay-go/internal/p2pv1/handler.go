package p2pv1

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
)

// ===============================
// Wire protocol message types
// ===============================

const (
	MsgSignedHeader = "SIGNED_HEADER"
	MsgGetBlocks    = "GET_BLOCKS"
	MsgBlocks       = "BLOCKS"
	MsgTx           = "TX"
)

// ===============================
// Envelope
// ===============================

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// ===============================
// Block sync messages
// ===============================

type GetBlocksMsg struct {
	Hashes []string `json:"hashes"`
}

type BlocksMsg struct {
	Blocks []BlockMsg `json:"blocks"`
}

// ===============================
// Handler
// ===============================

type Handler struct {
	node *Node
}

// Constructor (Node → Handler one-way, anti-cycle)
func NewHandler(n *Node) *Handler {
	return &Handler{node: n}
}

// ===============================
// Entry point from Peer
// ===============================

func (h *Handler) HandleMessage(p *Peer, msg *Message) error {
	if msg == nil {
		return errors.New("nil message")
	}

	switch msg.Type {

	case MsgSignedHeader:
		return h.handleSignedHeader(p, msg.Data)

	case MsgGetBlocks:
		return h.handleGetBlocks(p, msg.Data)

	case MsgBlocks:
		return h.handleBlocks(p, msg.Data)

	case MsgTx:
		return h.handleTx(p, msg.Data)

	default:
		log.Printf("[HANDLER] unknown msg type: %s", msg.Type)
		return nil
	}
}

// ===============================
// SIGNED HEADER
// ===============================

func (h *Handler) handleSignedHeader(p *Peer, raw json.RawMessage) error {
	var sh SignedHeaderMsg
	if err := json.Unmarshal(raw, &sh); err != nil {
		return err
	}

	// cryptographic verification
	if err := VerifySignedHeader(&sh); err != nil {
		return err
	}

	// semantic validation (height / parent)
	if err := ValidateHeaderBasic(sh.Header, h.node.headerPool); err != nil {
		return err
	}

	h.node.AddHeader(sh.Header)

	log.Printf(
		"[HEADER] accepted height=%d hash=%s",
		sh.Header.Height,
		sh.Header.Hash,
	)

	return nil
}

// ===============================
// TX RECEIVE + RATE LIMIT
// ===============================

func (h *Handler) handleTx(p *Peer, raw json.RawMessage) error {
	var tx Transaction
	if err := json.Unmarshal(raw, &tx); err != nil {
		return err
	}

	// rate-limit & anti-rebroadcast
	if !h.node.txGossip.Allow(tx.Hash) {
		return nil
	}

	if err := h.node.txPool.Add(tx, h.node.state); err != nil {
		return nil
	}

	h.node.BroadcastTxExcept(tx, p)
	return nil
}

// ===============================
// GET_BLOCKS
// ===============================

func (h *Handler) handleGetBlocks(p *Peer, raw json.RawMessage) error {
	var req GetBlocksMsg
	if err := json.Unmarshal(raw, &req); err != nil {
		return err
	}

	// hard bound anti-spam
	if len(req.Hashes) == 0 || len(req.Hashes) > 128 {
		return nil
	}

	blocks := make([]BlockMsg, 0, len(req.Hashes))
	for _, hash := range req.Hashes {
		if b, ok := h.node.GetBlock(hash); ok {
			blocks = append(blocks, b)
		}
	}

	if len(blocks) == 0 {
		return nil
	}

	return p.SendBlocks(BlocksMsg{Blocks: blocks})
}

// ===============================
// BLOCKS (VALIDATE STATE ROOT + FINALITY LOG)
// ===============================

func (h *Handler) handleBlocks(p *Peer, raw json.RawMessage) error {
	var msg BlocksMsg
	if err := json.Unmarshal(raw, &msg); err != nil {
		return err
	}

	if len(msg.Blocks) == 0 {
		return nil
	}

	for _, b := range msg.Blocks {

		// idempotent
		if h.node.HasBlock(b.Hash) {
			continue
		}

		// parent-first rule
		if b.ParentHash != "" && !h.node.HasBlock(b.ParentHash) {
			_ = p.SendGetBlocks(GetBlocksMsg{
				Hashes: []string{b.ParentHash},
			})
			continue
		}

		// ===============================
		// STATE ROOT VALIDATION
		// ===============================

		if h.node.state.Height()+1 == b.Height && len(b.StateRoot) > 0 {
			expected := h.node.state.RootHash()
			if !bytes.Equal(expected, b.StateRoot) {
				log.Printf(
					"[BLOCK] reject invalid state root height=%d hash=%s",
					b.Height,
					b.Hash,
				)
				continue
			}
		}

		// store block body
		if err := h.node.StoreBlock(b); err != nil {
			continue
		}

		// store header (state root will be re-attached after execution)
		h.node.AddHeader(HeaderMsg{
			Hash:       b.Hash,
			ParentHash: b.ParentHash,
			Height:     b.Height,
			Timestamp:  b.Timestamp,
			StateRoot:  b.StateRoot,
		})

		log.Printf(
			"[BLOCK] accepted height=%d hash=%s",
			b.Height,
			b.Hash,
		)
	}

	// trigger execution
	h.node.TryExecuteBlocks()

	// ===============================
	// OPTIONAL FINALITY LOG
	// ===============================

	if f := h.node.finality.LastFinalized(); f > 0 {
		log.Printf("[FINALITY] finalized checkpoint height=%d", f)
	}

	return nil
}
