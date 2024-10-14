package main

import (
	"context"
	"log"
	"time"

	pb "github.com/BlackBorada/ProjectD/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTestServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Client"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for i := 0; i < 10; i++ {
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Client"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
