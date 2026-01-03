package p2pv1

import (
	"log"
)

type Node struct {
	listenAddr string
	networkID  string
	peerAddr   string

	server *Server
}

func NewNode(listenAddr, networkID, peerAddr string) *Node {
	n := &Node{
		listenAddr: listenAddr,
		networkID:  networkID,
		peerAddr:   peerAddr,
	}
	n.server = NewServer(n, listenAddr)
	return n
}

func (n *Node) Run() {
	log.Printf("[P2P] node starting (listen=%s, net=%s)", n.listenAddr, n.networkID)

	go n.server.Run()

	if n.peerAddr != "" {
		log.Printf("[P2P] dialing peer %s", n.peerAddr)
		go DialPeer(n, n.peerAddr)
	}
}

func (n *Node) handleConn(p *Peer) {
	log.Printf("[P2P] connected peer %s", p.addr)

	// handshake
	if err := SendHandshake(p, n.networkID); err != nil {
		log.Println("[P2P] handshake error:", err)
		return
	}

	go HandlePeer(p)
}
