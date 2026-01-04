package p2pv1

type MessageType uint8

const (
	MsgHeader MessageType = 1
)

type HeaderMsg struct {
	Height uint64
	Hash   []byte
}

type Message struct {
	Type   MessageType
	Header *HeaderMsg
}
