package service

import (
	"context"
	"encoding/json" 
	"fmt"
	pb "github.com/irlan/quantumpay-go/internal/grpc/proto"
)

// Hapus definisi ChainAPI dari sini agar tidak "Redeclared"
// Kita pakai ChainAPI dari interface.go

type NodeService struct {
	pb.UnimplementedNodeServiceServer
	chain ChainAPI
}

func NewNodeService(chain ChainAPI) *NodeService {
	return &NodeService{chain: chain}
}

// 🛡️ Handshake 77077
func (s *NodeService) Handshake(ctx context.Context, req *pb.HandshakeRequest) (*pb.HandshakeResponse, error) {
	const officialID = 77077
	const officialGenesis = "0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a"

	if req.ChainId != officialID || req.GenesisHash != officialGenesis {
		return &pb.HandshakeResponse{
			Success: false,
			Message: "Rejected: Network Identity Mismatch!",
		}, nil
	}

	return &pb.HandshakeResponse{
		Success: true,
		Message: "Connected to QuantumPay Alpha-Mainnet",
	}, nil
}

// 📊 Get Status
func (s *NodeService) GetStatus(ctx context.Context, _ *pb.Empty) (*pb.NodeStatus, error) {
	return &pb.NodeStatus{
		Height:   s.chain.Height(),
		LastHash: "active-sync", 
	}, nil
}

// 📦 Get Block (JSON Marshal)
func (s *NodeService) GetBlockByHeight(ctx context.Context, req *pb.BlockRequest) (*pb.BlockResponse, error) {
	// Di sini kita terima 'any', bukan *core.Block
	block, ok := s.chain.GetBlockByHeight(req.Height)
	if !ok {
		return &pb.BlockResponse{Found: false}, nil
	}

	// json.Marshal bisa memproses 'any' dengan mulus
	data, err := json.Marshal(block)
	if err != nil {
		return nil, fmt.Errorf("failed to encode block: %v", err)
	}

	return &pb.BlockResponse{
		Found: true,
		Data:  data,
	}, nil
}
