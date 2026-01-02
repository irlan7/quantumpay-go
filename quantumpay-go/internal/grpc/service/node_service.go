package service

import (
	"context"

	pb "github.com/irlan/quantumpay-go/internal/grpc/proto"
)

type NodeService struct {
	pb.UnimplementedNodeServiceServer
	chain ChainAPI
}

func NewNodeService(chain ChainAPI) *NodeService {
	return &NodeService{chain: chain}
}

func (s *NodeService) GetStatus(ctx context.Context, _ *pb.Empty) (*pb.NodeStatus, error) {
	return &pb.NodeStatus{
		Height: s.chain.Height(),
	}, nil
}

func (s *NodeService) GetBlockByHeight(
	ctx context.Context,
	req *pb.BlockRequest,
) (*pb.BlockResponse, error) {

	block, ok := s.chain.GetBlockByHeight(req.Height)
	if !ok {
		return &pb.BlockResponse{Found: false}, nil
	}

	data, err := block.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return &pb.BlockResponse{
		Found: true,
		Data:  data,
	}, nil
}
