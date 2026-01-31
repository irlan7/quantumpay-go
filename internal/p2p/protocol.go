package p2p

type Handshake struct {
	Version int    `json:"version"`
	ChainID string `json:"chain_id"`
	NodeID  string `json:"node_id"`
}

type HandshakeAck struct {
	OK bool `json:"ok"`
}
