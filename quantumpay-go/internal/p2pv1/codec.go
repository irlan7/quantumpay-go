package p2pv1

import (
	"encoding/json"
	"net"
)

func Send(conn net.Conn, msg Message) error {
	return json.NewEncoder(conn).Encode(msg)
}

func Receive(conn net.Conn, msg *Message) error {
	return json.NewDecoder(conn).Decode(msg)
}
