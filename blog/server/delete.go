package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"log"
)

func (s *Server) Delete(ctx context.Context, in *pb.BlogID) (*empty.Empty, error) {
	log.Printf("Delete was invoked with %v", in)

	for index, blog := range BlogList {
		if blog.ID == in.GetId() {
			BlogList = append(BlogList[:index], BlogList[index+1:]...)
		}

		break
	}

	return nil, nil

}
