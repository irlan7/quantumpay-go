package rpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/irlan/quantumpay-go/internal/rpc/nodepb"
)

func StartGRPC(
	addr string,
	service *NodeService,
) {
	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("[gRPC] listen error: %v", err)
		}

		server := grpc.NewServer()
		pb.RegisterNodeServiceServer(server, service)

		log.Println("[gRPC] listening on", addr)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("[gRPC] serve error: %v", err)
		}
	}()
}
