package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer - Start a gRPC based server
func StartGrpcServer() {

	svAddress := "localhost:50051"

	lis, err := net.Listen("tcp", svAddress)
	if err != nil {
		log.Fatalf("Could not connect", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	log.Printf("gRPC Server started on => %v\n", svAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not connect", err)
	}

	return
}
