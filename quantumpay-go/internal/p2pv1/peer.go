package p2pv1

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"
)

/*
Peer merepresentasikan 1 koneksi P2P.
Tanggung jawab:
- kirim message (thread-safe)
- terima message dan lempar ke Handler
- TIDAK menyimpan state, tx, header, atau block (ANTI CYCLE)
*/

// ===============================
// Peer struct
// ===============================

type Peer struct {
	conn net.Conn
	enc  *json.Encoder
	dec  *json.Decoder

	mu     sync.Mutex // proteksi write (send)
	closed bool
}

// ===============================
// Constructor
// ===============================

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
		enc:  json.NewEncoder(conn),
		dec:  json.NewDecoder(bufio.NewReader(conn)),
	}
}

// ===============================
// Lifecycle
// ===============================

// Close menutup koneksi peer secara idempotent
func (p *Peer) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}
	p.closed = true
	_ = p.conn.Close()
}

// ===============================
// SEND HELPERS
// ===============================

// SendSignedHeader dipakai untuk Signed Header Gossip
func (p *Peer) SendSignedHeader(sh SignedHeaderMsg) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return io.ErrClosedPipe
	}

	env := Message{
		Type: MsgSignedHeader,
		Data: sh,
	}
	return p.enc.Encode(&env)
}

// SendBlock dipakai untuk block sync
func (p *Peer) SendBlock(b BlockMsg) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return io.ErrClosedPipe
	}

	env := Message{
		Type: MsgBlocks, // kirim sebagai BLOCKS (single block dibolehkan)
		Data: BlocksMsg{Blocks: []BlockMsg{b}},
	}
	return p.enc.Encode(&env)
}

// SendGetBlocks dipakai untuk request block (batch)
func (p *Peer) SendGetBlocks(req GetBlocksMsg) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return io.ErrClosedPipe
	}

	env := Message{
		Type: MsgGetBlocks,
		Data: req,
	}
	return p.enc.Encode(&env)
}

// ===============================
// TX SEND (PATCH UTAMA)
// ===============================

// SendTx mengirim transaction ke peer (Tx Gossip)
func (p *Peer) SendTx(tx Transaction) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return io.ErrClosedPipe
	}

	env := Message{
		Type: MsgTx,
		Data: tx,
	}
	return p.enc.Encode(&env)
}

// ===============================
// RECEIVE LOOP
// ===============================

// ReceiveLoop membaca message terus-menerus dan menyerahkan ke handler
func (p *Peer) ReceiveLoop(handler *Handler) {
	defer p.Close()

	for {
		var env Message
		if err := p.dec.Decode(&env); err != nil {
			if err != io.EOF {
				log.Printf("[P2P] peer read error: %v", err)
			}
			return
		}

		// Handler adalah satu-satunya pintu masuk logic
		if err := handler.HandleMessage(p, &env); err != nil {
			log.Printf("[P2P] handler error: %v", err)
			return
		}
	}
}
