package main

import (
	"context"
	"log"

	pb "github.com/saeedjhn/go-grpc/blog/goproto"
)

func doRead(c pb.BlogServiceClient, id string) {
	log.Println("doRead was invoked")

	req := &pb.BlogID{Id: id}

	res, err := c.Read(context.Background(), req)
	if err != nil {
		log.Fatalf("Error happend while reading: %v", err)
	}

	log.Printf("Information blog is: %v", res)
}
