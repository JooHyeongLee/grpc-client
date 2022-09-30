package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-client/grpcapi"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:19001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	conn.Connect()
	defer conn.Close()

	hello := pb.NewGreeterServiceClient(conn)
	request := pb.SayGoodbyeRequest{RequestMessage: "good bye"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// r, err := c.Get(ctx, &pb.ConfigRequest{Profile: "dev"})
	r, err := hello.SayGoodbye(ctx, &request)

	if err != nil {
		log.Fatalf("could not request: %v", err)
	}

	log.Printf("Config: %v", r)
}