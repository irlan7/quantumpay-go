package p2p

import (
	"bufio"
	"encoding/json"
	"net"
)

type Client struct {
	ChainID string
	NodeID  string
}

func (c *Client) Connect(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	hs, _ := Encode(Handshake{
		Version: ProtocolVersion,
		ChainID: c.ChainID,
		NodeID:  c.NodeID,
	})

	SendMessage(conn, Message{
		Type: MsgHandshake,
		Data: hs,
	})

	reader := bufio.NewReader(conn)
	line, _ := reader.ReadBytes('\n')

	var ack Message
	_ = json.Unmarshal(line, &ack)

	if ack.Type != MsgHandshakeAck {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
