package p2pv1

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Node struct {
	cfg   Config
	peers map[string]*Peer
	mu    sync.Mutex
}

func NewNode(cfg Config) *Node {
	return &Node{
		cfg:   cfg,
		peers: make(map[string]*Peer),
	}
}

// ===== Lifecycle =====

func (n *Node) Start() error {
	go n.listen()

	if n.cfg.PeerAddr != "" {
		go n.dialPeer(n.cfg.PeerAddr)
	}

	return nil
}

// ===== Network =====

func (n *Node) listen() {
	addr := fmt.Sprintf(":%d", n.cfg.ListenPort)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("[P2P] listen error:", err)
		return
	}

	log.Println("[P2P] listening on", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go n.handleConn(conn)
	}
}

func (n *Node) dialPeer(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("[P2P] dial error:", err)
		return
	}
	n.addPeer(addr, conn)
	go n.handlePeer(n.peers[addr])
}

func (n *Node) handleConn(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	n.addPeer(addr, conn)
	n.handlePeer(n.peers[addr])
}

// ===== Peer =====

func (n *Node) addPeer(addr string, conn net.Conn) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if _, ok := n.peers[addr]; ok {
		return
	}

	n.peers[addr] = &Peer{conn: conn, addr: addr}
	log.Println("[P2P] peer connected:", addr)
}

func (n *Node) handlePeer(p *Peer) {
	for {
		msg, err := readMessage(p.conn)
		if err != nil {
			return
		}

		if msg.Type == MsgHeader && msg.Header != nil {
			log.Printf(
				"[P2P] header gossip from %s height=%d hash=%x\n",
				p.addr,
				msg.Header.Height,
				msg.Header.Hash,
			)
		}
	}
}

// ===== Gossip =====

func (n *Node) BroadcastHeader(h HeaderMsg) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, p := range n.peers {
		_ = sendMessage(p.conn, &Message{
			Type:   MsgHeader,
			Header: &h,
		})
	}
}
