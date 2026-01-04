package p2pv1

import (
	"encoding/gob"
	"net"
)

func sendMessage(conn net.Conn, msg *Message) error {
	enc := gob.NewEncoder(conn)
	return enc.Encode(msg)
}

func readMessage(conn net.Conn) (*Message, error) {
	dec := gob.NewDecoder(conn)
	var msg Message
	err := dec.Decode(&msg)
	return &msg, err
}
