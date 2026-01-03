package p2pv1

import (
	"errors"
)

func SendHandshake(p *Peer, network string) error {
	return Send(p.conn, Message{
		Type: MsgHandshake,
		Data: []byte(network),
	})
}

func HandlePeer(p *Peer) {
	for {
		var msg Message
		if err := Receive(p.conn, &msg); err != nil {
			return
		}

		switch msg.Type {
		case MsgHandshake:
			// ok
		case MsgPing:
			Send(p.conn, Message{Type: MsgPong})
		}
	}
}

func ValidateNetwork(remote, local string) error {
	if remote != local {
		return errors.New("network mismatch")
	}
	return nil
}
