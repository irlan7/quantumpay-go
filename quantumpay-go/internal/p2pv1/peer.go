package p2pv1

import (
	"net"
)

type Peer struct {
	conn net.Conn
	addr string
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
		addr: conn.RemoteAddr().String(),
	}
}
