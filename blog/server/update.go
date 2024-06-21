package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"log"
)

func (s *Server) Update(ctx context.Context, in *pb.Blog) (*empty.Empty, error) {
	log.Printf("Update was invoked with %v", in)

	for _, blog := range BlogList {
		if blog.ID == in.GetId() {
			blog.AuthorID = in.GetAuthorId()
			blog.Title = in.GetTitle()
			blog.Content = in.GetContent()
		}

		break
	}

	return nil, nil
}
