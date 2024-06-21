package main

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"log"
	"time"
)

func (s *Server) List(in *empty.Empty, stream pb.BlogService_ListServer) error {
	log.Println("List was invoked")

	for _, blogItem := range BlogList {
		log.Printf("Sending blogItem: %v", blogItem)
		stream.Send(mapGetBlogResponseToProtobuf(blogItem))

		time.Sleep(1 * time.Second)
	}

	return nil
}
