package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/greet/goproto"
	"io"
	"log"
	"time"
)

func doGreetEveryOnce(c pb.GreetServiceClient) {
	log.Println("doGreetEveryOnce was invoked")

	stream, err := c.GreetEveryOnce(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream %v", err)
	}

	waitC := make(chan struct{})

	go func() {
		reqs := []*pb.GreetRequest{
			{FirstName: "John"},
			{FirstName: "Happy"},
			{FirstName: "Marie"},
			{FirstName: "Test"},
		}

		for _, req := range reqs {
			log.Printf("Send data: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving %v", err)
				break
			}

			log.Printf("Receiving %v\n", res.Result)
		}

		close(waitC)
	}()

	<-waitC
}
