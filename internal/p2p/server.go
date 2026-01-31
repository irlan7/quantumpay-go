package p2p

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

type Server struct {
	Addr    string
	ChainID string
	NodeID  string
	OnBlock func([]byte)
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	log.Println("🌐 P2P listening on", s.Addr)

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			go s.handleConn(conn)
		}
	}()

	return nil
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Handshake
	line, _ := reader.ReadBytes('\n')
	var msg Message
	_ = json.Unmarshal(line, &msg)

	if msg.Type != MsgHandshake {
		return
	}

	var hs Handshake
	_ = json.Unmarshal(msg.Data, &hs)

	if hs.ChainID != s.ChainID || hs.Version != ProtocolVersion {
		SendMessage(conn, Message{Type: MsgError})
		return
	}

	ack, _ := Encode(HandshakeAck{OK: true})
	SendMessage(conn, Message{Type: MsgHandshakeAck, Data: ack})

	// Listen messages
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		var m Message
		if err := json.Unmarshal(line, &m); err != nil {
			continue
		}

		if m.Type == MsgBlock && s.OnBlock != nil {
			s.OnBlock(m.Data)
		}
	}
}
