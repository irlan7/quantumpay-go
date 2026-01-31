package p2pv1

const (
	MsgHeader = 1
)

// Message adalah envelope tunggal
type Message struct {
	Type uint8
	Data []byte
}

// HeaderMsg adalah payload untuk gossip header
type HeaderMsg struct {
	Height uint64
	Hash   []byte
}
