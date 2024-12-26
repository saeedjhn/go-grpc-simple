package main

import (
	"context"
	"log"

	pb "github.com/saeedjhn/go-grpc/blog/goproto"
)

func doDelete(c pb.BlogServiceClient, id string) {
	log.Println("doDelete was invoked")

	req := &pb.BlogID{Id: id}

	_, err := c.Delete(context.Background(), req)
	if err != nil {
		log.Fatalf("Error happend while deleting %v", err)
	}

	log.Println("Blog was deleted")
}
