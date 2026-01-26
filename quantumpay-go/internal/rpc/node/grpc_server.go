package node

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/irlan/quantumpay-go/internal/rpc/nodepb"
)

func StartGRPC(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	pb.RegisterNodeServiceServer(srv, NewNodeService())

	log.Println("[gRPC] listening on", addr)
	return srv.Serve(lis)
}
