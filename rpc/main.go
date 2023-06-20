package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "rpc/pb/greeter" // Update with your module name
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// Prepare the request
	request := &pb.HelloRequest{
		Name: "John",
	}

	// Call the gRPC method
	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	log.Printf("Response: %s", response.Message)
}
