package p2pv1

import (
	"log"
	"net"
)

type Node struct {
	cfg   Config
	peers []*Peer
}

func NewNode(cfg Config) *Node {
	return &Node{
		cfg:   cfg,
		peers: make([]*Peer, 0),
	}
}

func (n *Node) Start() error {
	ln, err := net.Listen("tcp", n.cfg.ListenAddress)
	if err != nil {
		return err
	}

	log.Printf("[p2p] listening on %s", n.cfg.ListenAddress)

	go n.acceptLoop(ln)
	return nil
}

func (n *Node) acceptLoop(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("[p2p] accept error: %v", err)
			continue
		}

		peer := NewPeer(conn)
		n.peers = append(n.peers, peer)

		log.Printf("[p2p] new inbound peer %s", conn.RemoteAddr())
	}
}
