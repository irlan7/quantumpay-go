package blockchain

type ChainView struct {
	lastHash []byte
	height   uint64
}

func NewChainView(chain *Blockchain) *ChainView {
	return &ChainView{
		lastHash: chain.LastHash(),
		height:   chain.Height(),
	}
}

func (v *ChainView) LastHash() []byte {
	return v.lastHash
}

func (v *ChainView) Height() uint64 {
	return v.height
}
