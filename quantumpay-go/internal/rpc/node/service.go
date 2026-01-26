package node

import (
	"context"

	pb "github.com/irlan/quantumpay-go/internal/rpc/nodepb"
)

type NodeService struct {
	pb.UnimplementedNodeServiceServer
}

func NewNodeService() *NodeService {
	return &NodeService{}
}

func (s *NodeService) ChainHeight(
	ctx context.Context,
	req *pb.Empty,
) (*pb.ChainHeight, error) {

	// SAFE STUB (tidak sentuh blockchain)
	return &pb.ChainHeight{
		Height: 0,
	}, nil
}
