package main

import (
	"context"
	"log"
	"time"

	pb "github.com/saeedjhn/go-grpc/greet/goproto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Happy", LastName: "Doe"},
		{FirstName: "Smith", LastName: "Doe"},
		{FirstName: "Marie", LastName: "Doe"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet %v", err)
	}

	log.Printf("LongGreet %s\n", res.Result)
}
