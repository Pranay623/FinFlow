package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Portfolio Service Starting...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// We will register services here later
	
	log.Println("Portfolio gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
