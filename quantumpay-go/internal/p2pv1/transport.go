package p2pv1

import (
	"net"
)

func DialPeer(n *Node, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	p := NewPeer(conn)
	n.handleConn(p)
}
