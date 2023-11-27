package main

import (
	"log"
	"net"
	pb "productinfo/ecommerce"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on Port: " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to start server, %v", err)
	}
}
