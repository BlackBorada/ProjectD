package server

import (
	"context"
	"log"

	pb "github.com/BlackBorada/ProjectD/pkg/proto"
)

type TestServer struct {
	pb.UnimplementedTestServiceServer
}

func (s *TestServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
