package p2pv1

type Message struct {
	Type string
	Data []byte
}

const (
	MsgHandshake = "handshake"
	MsgPing      = "ping"
	MsgPong      = "pong"
)
