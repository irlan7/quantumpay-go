package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/irlan/quantumpay-go/internal/grpc/proto"
)

// Server struct membungkus gRPC server
type Server struct {
	listenAddr string
	grpcServer *grpc.Server
}

// NewServer membuat instance server baru
func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		grpcServer: grpc.NewServer(),
	}
}

// Start menjalankan server pada port yang ditentukan
func (s *Server) Start(nodeService pb.NodeServiceServer) error {
	lis, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Daftarkan service kita ke server gRPC
	pb.RegisterNodeServiceServer(s.grpcServer, nodeService)

	log.Printf("✅ gRPC Server listening on %s", s.listenAddr)
	return s.grpcServer.Serve(lis)
}

// Stop mematikan server dengan aman
func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
