package node

import (
	"fmt"
	"net"

	pb "github.com/irlan/quantumpay-go/internal/grpc/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server
	port   string
}

func NewServer(port string) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
		port:   port,
	}
}

// Start sekarang menerima adapter sebagai dependency
func (s *GRPCServer) Start(adapter ChainAdapter) error {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// PERBAIKAN UTAMA DISINI:
	// Sebelumnya: service := NewNodeService() -> ERROR
	// Sekarang: Kita masukkan adapter ke dalam constructor
	nodeService := NewNodeService(adapter)

	// Register service ke gRPC server
	pb.RegisterNodeServer(s.server, nodeService)

	fmt.Printf("📡 gRPC Server listening on %s\n", s.port)
	return s.server.Serve(lis)
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}
