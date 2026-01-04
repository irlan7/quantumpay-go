package p2pv1

import "net"

type Peer struct {
	conn net.Conn
	addr string
}
