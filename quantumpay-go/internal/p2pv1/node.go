package p2pv1

import (
	"log"
	"net"
	"sync"
)

type Node struct {
	cfg     Config
	headers map[uint64]BlockHeader
	mu      sync.Mutex
}

func NewNode(cfg Config) *Node {
	return &Node{
		cfg:     cfg,
		headers: make(map[uint64]BlockHeader),
	}
}

func (n *Node) Start() error {
	addr := ":" + n.cfg.ListenPort
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Println("[P2P] listening on", addr)

	if n.cfg.PeerAddr != "" {
		go n.dialPeer(n.cfg.PeerAddr)
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go n.handleConn(c)
	}
}

func (n *Node) dialPeer(addr string) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("[P2P] dial failed:", err)
		return
	}
	n.handleConn(c)
}

func (n *Node) handleConn(c net.Conn) {
	p := NewPeer(c)

	buf := make([]byte, 4096)
	for {
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		msg := Message{
			Type: buf[0],
			Data: buf[1:nr],
		}
		n.handleMessage(p, msg)
	}
}

// === HEADER STORAGE (MEMORY ONLY) ===

func (n *Node) storeHeader(h HeaderMsg) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.headers[h.Height] = BlockHeader{
		Height: h.Height,
		Hash:   h.Hash,
	}
}
