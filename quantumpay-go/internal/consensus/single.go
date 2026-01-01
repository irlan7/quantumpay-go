package consensus

// Single proposer consensus (dev / testnet)

type Single struct {
	nodeID string
}

func NewSingle(nodeID string) *Single {
	return &Single{nodeID: nodeID}
}

// IsProposer selalu true (single node mode)
func (s *Single) IsProposer() bool {
	return true
}

// ID mengembalikan identifier proposer
func (s *Single) ID() string {
	return s.nodeID
}
