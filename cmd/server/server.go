package main

import (
	"log"
	"net"

	"github.com/rvgoncalves/fs2-grpc/pb"
	"github.com/rvgoncalves/fs2-grpc/services"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not connect %v", err)
	}

}
