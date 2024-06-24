package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream %v", err)
	}

	waitC := make(chan struct{})

	go func() {
		reqs := []int32{1, 5, 7, 2, 99, 101, 11, 202, 500, 440, 60}

		for _, req := range reqs {
			log.Printf("Sending number %d", req)
			stream.Send(&pb.MaxRequest{N: req})
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

			log.Printf("Receiving new maximum %d\n", res.Result)
		}

		close(waitC)
	}()

	<-waitC
}
