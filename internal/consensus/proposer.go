package consensus

type Proposer interface {
	IsProposer() bool
	ID() string
}
