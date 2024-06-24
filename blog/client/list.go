package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"io"
	"log"
)

func doList(c pb.BlogServiceClient) {
	log.Println("doList was invoked")

	stream, err := c.List(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("Error happend was ")
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error happend %v", err)
		}

		log.Println(res)
	}
}
