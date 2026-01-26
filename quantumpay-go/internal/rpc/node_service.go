package rpc

import (
	"context"

	pb "github.com/irlan/quantumpay-go/internal/rpc/nodepb"
)

// === Interfaces (Dependency Inversion) ===

type ChainReader interface {
	Height() uint64
	GetBlockByHeight(uint64) (BlockView, bool)
}

type AccountReader interface {
	GetBalance(address string) string
}

type BlockView struct {
	Height uint64
	Hash   string
}

// === gRPC Service ===

type NodeService struct {
	pb.UnimplementedNodeServiceServer
	chain   ChainReader
	account AccountReader
}

func NewNodeService(c ChainReader, a AccountReader) *NodeService {
	return &NodeService{
		chain:   c,
		account: a,
	}
}

// ===== RPCs =====

func (s *NodeService) GetHeight(
	ctx context.Context,
	_ *pb.GetHeightRequest,
) (*pb.GetHeightResponse, error) {
	return &pb.GetHeightResponse{
		Height: s.chain.Height(),
	}, nil
}

func (s *NodeService) Health(
	ctx context.Context,
	_ *pb.HealthRequest,
) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{
		Status: "ok",
	}, nil
}
