package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/greet/goproto"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "John"}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			return
		}

		log.Fatalf("A non GRPC error: %v\n", err)
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
