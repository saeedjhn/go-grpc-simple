package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"log"
)

func doUpdate(c pb.BlogServiceClient, id string) {
	log.Println("doUpdate was invoked")

	req := &pb.Blog{
		Id:       id,
		AuthorId: "123",
		Title:    "Update title",
		Content:  "Update content",
	}

	_, err := c.Update(context.Background(), req)
	if err != nil {
		log.Fatalf("Error happend while updating %v", err)
	}

	log.Println("Blog was updated")
}
