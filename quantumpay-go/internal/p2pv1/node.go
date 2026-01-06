package p2pv1

import (
	"log"
	"net"
	"sync"
	"time"
)

/*
Node = orchestrator utama blockchain node

Tanggung jawab:
- wire P2P
- wire tx gossip
- wire state execution
- wire finality v1 & v2
- TIDAK mengandung logic finality / crypto langsung (ANTI-CYCLE)
*/

// ===============================
// Node structure
// ===============================

type Node struct {
	cfg Config

	listener net.Listener

	peers   map[*Peer]struct{}
	peersMu sync.Mutex

	headerPool *HeaderPool
	blockStore *BlockStore
	state      *State

	txPool   *TxPool
	txGossip *TxGossip

	finality   *FinalityGadget // v1 (checkpoint lokal)
	finalityV2 *FinalityV2     // v2 (validator vote)

	handler *Handler
}

// ===============================
// Constructor
// ===============================

func NewNode(cfg Config) *Node {
	n := &Node{
		cfg:        cfg,
		peers:      make(map[*Peer]struct{}),
		headerPool: NewHeaderPool(),
		blockStore: NewBlockStore(),
		state:      NewState(),
		txPool:     NewTxPool(10_000),
		txGossip:   NewTxGossip(30 * time.Second),
		finality:   NewFinalityGadget(),
		finalityV2: NewFinalityV2(cfg.Validators),
	}

	// one-way dependency: Node → Handler
	n.handler = NewHandler(n)
	return n
}

// ===============================
// Start / Stop
// ===============================

func (n *Node) Start(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	n.listener = ln

	log.Printf("[NODE] listening on %s", addr)
	go n.acceptLoop()
	return nil
}

func (n *Node) Stop() {
	if n.listener != nil {
		_ = n.listener.Close()
	}

	n.peersMu.Lock()
	defer n.peersMu.Unlock()

	for p := range n.peers {
		p.Close()
	}
	log.Printf("[NODE] stopped")
}

// ===============================
// Peer handling
// ===============================

func (n *Node) acceptLoop() {
	for {
		conn, err := n.listener.Accept()
		if err != nil {
			log.Printf("[NODE] accept error: %v", err)
			return
		}

		p := NewPeer(conn)
		n.addPeer(p)

		log.Printf("[NODE] inbound peer connected")
		go p.ReceiveLoop(n.handler)
	}
}

func (n *Node) Dial(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	p := NewPeer(conn)
	n.addPeer(p)

	log.Printf("[NODE] outbound peer connected: %s", addr)
	go p.ReceiveLoop(n.handler)
	return nil
}

func (n *Node) addPeer(p *Peer) {
	n.peersMu.Lock()
	defer n.peersMu.Unlock()
	n.peers[p] = struct{}{}
}

// ===============================
// TX GOSSIP
// ===============================

func (n *Node) BroadcastTxExcept(tx Transaction, except *Peer) {
	n.peersMu.Lock()
	defer n.peersMu.Unlock()

	for p := range n.peers {
		if p == except {
			continue
		}
		_ = p.SendTx(tx)
	}
}

func (n *Node) InjectTx(tx Transaction) error {
	if !n.txGossip.Allow(tx.Hash) {
		return nil
	}
	if err := n.txPool.Add(tx, n.state); err != nil {
		return err
	}
	n.BroadcastTxExcept(tx, nil)
	return nil
}

// ===============================
// Block / Header helpers
// ===============================

func (n *Node) HasBlock(hash string) bool {
	return n.blockStore.Has(hash)
}

func (n *Node) GetBlock(hash string) (BlockMsg, bool) {
	return n.blockStore.Get(hash)
}

func (n *Node) StoreBlock(b BlockMsg) error {
	return n.blockStore.Add(b)
}

// ===============================
// Header + Finality wiring
// ===============================

func (n *Node) AddHeader(h HeaderMsg) {
	// fork-choice
	n.headerPool.Add(h)

	// finality v1 (checkpoint lokal)
	n.finality.OnHeader(h)
	n.finality.TryFinalize()

	// finality v2 TIDAK auto vote (anti-cycle)
}

// ===============================
// STATE EXECUTION + STATE ROOT
// ===============================

func (n *Node) TryExecuteBlocks() {
	for {
		tip, err := n.headerPool.Tip()
		if err != nil {
			return
		}

		if n.state.Height() >= tip.Height {
			return
		}

		b, ok := n.blockStore.NextUnexecuted("")
		if !ok {
			return
		}

		snap := n.state.Snapshot()
		if err := n.state.ApplyBlock(b); err != nil {
			n.state = snap
			return
		}

		// compute deterministic state root
		root := n.state.RootHash()

		// mark executed
		n.blockStore.MarkExecuted(b.Hash)

		// bind header with state root
		h := HeaderMsg{
			Hash:       b.Hash,
			ParentHash: b.ParentHash,
			Height:     b.Height,
			Timestamp:  b.Timestamp,
			StateRoot:  root,
		}
		n.headerPool.Add(h)

		// cleanup mempool
		for _, tx := range b.Txs {
			n.txPool.Remove(tx.Hash)
		}

		log.Printf(
			"[EXEC] block height=%d hash=%s stateRoot=%x",
			b.Height,
			b.Hash,
			root,
		)
	}
}

// ===============================
// FINALITY V2 HELPERS
// ===============================

// SubmitVote dipakai validator lokal / RPC
func (n *Node) SubmitVote(v Vote) error {
	return n.finalityV2.OnVote(v)
}

// LastFinalizedCheckpoint (v2)
func (n *Node) LastFinalizedCheckpoint() uint64 {
	return n.finalityV2.LastFinalized()
}
