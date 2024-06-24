package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"log"
)

func doCreate(c pb.BlogServiceClient, blog *pb.Blog) string {
	log.Println("doCreate function was invoked")

	//req := &pb.Blog{
	//	AuthorId: "1",
	//	Title:    "This is my title",
	//	Content:  "This is my content",
	//}

	res, err := c.Create(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
