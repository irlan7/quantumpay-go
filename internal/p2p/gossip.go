package p2p

import "net"

func GossipBlock(peers []net.Conn, block []byte) {
	for _, p := range peers {
		SendMessage(p, Message{
			Type: MsgBlock,
			Data: block,
		})
	}
}
