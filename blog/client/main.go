package main

import (
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const addr = "0.0.0.0:50001"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	log.Printf("Connection is successful on %s\n", addr)

	c := pb.NewBlogServiceClient(conn)

	b := &pb.Blog{
		AuthorId: "1",
		Title:    "This is my title",
		Content:  "This is my content",
	}

	id := doCreate(c, b)
	doCreate(c, b)
	doCreate(c, b)

	doRead(c, id)
	doUpdate(c, id)
	doRead(c, id)
	doDelete(c, id)
	doList(c)
}
