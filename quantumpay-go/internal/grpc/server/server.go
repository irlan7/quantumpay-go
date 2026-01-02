package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr string
	srv  *grpc.Server
}

func New(addr string) *GRPCServer {
	return &GRPCServer{
		addr: addr,
		srv:  grpc.NewServer(),
	}
}

func (g *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", g.addr)
	if err != nil {
		return err
	}

	log.Println("🧩 gRPC server listening on", g.addr)
	return g.srv.Serve(lis)
}

func (g *GRPCServer) Stop() {
	log.Println("🛑 gRPC server stopped")
	g.srv.GracefulStop()
}
