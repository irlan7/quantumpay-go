package p2pv1

import (
	"log"
	"net"
)

type Server struct {
	node *Node
	addr string
}

func NewServer(node *Node, addr string) *Server {
	return &Server{
		node: node,
		addr: addr,
	}
}

func (s *Server) Run() {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Println("[P2P] listen error:", err)
		return
	}
	log.Println("[P2P] listening on", s.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		peer := NewPeer(conn)
		go s.node.handleConn(peer)
	}
}
