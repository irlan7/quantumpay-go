package p2p

import (
	"encoding/json"
	"net"
)

func SendMessage(conn net.Conn, msg Message) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(append(b, '\n'))
	return err
}
