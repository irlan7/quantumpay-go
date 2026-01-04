package p2pv1

import (
	"net"
)

type Peer struct {
	conn net.Conn
}

func NewPeer(c net.Conn) *Peer {
	return &Peer{conn: c}
}

func (p *Peer) Send(msg Message) error {
	_, err := p.conn.Write(append([]byte{msg.Type}, msg.Data...))
	return err
}
