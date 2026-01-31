package node

import (
	"context"

	// Pastikan path ini sesuai
	pb "github.com/irlan/quantumpay-go/internal/grpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChainAdapter interface {
	Height() uint64
}

type NodeService struct {
	pb.UnimplementedNodeServer // Error ini akan hilang setelah perintah protoc
	chain ChainAdapter
}

func NewNodeService(chain ChainAdapter) *NodeService {
	return &NodeService{chain: chain}
}

func (s *NodeService) GetHeight(ctx context.Context, req *emptypb.Empty) (*pb.ChainHeight, error) {
	return &pb.ChainHeight{
		Height: s.chain.Height(),
	}, nil
}

func (s *NodeService) SubmitTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	return &pb.TransactionResponse{
		Success: true,
		Hash:    "dummy_tx_hash",
	}, nil
}
