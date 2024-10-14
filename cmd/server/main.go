package main

import (
	"log"
	"net"

	"github.com/BlackBorada/ProjectD/internal/server"
	pb "github.com/BlackBorada/ProjectD/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server.TestServer{})

	log.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
