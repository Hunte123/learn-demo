package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "rpc/pb/greeter" // Update with your module name
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := "Hello, " + req.Name
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := &greeterServer{}
	pb.RegisterGreeterServer(grpcServer, server)

	log.Println("Server started on port :50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
