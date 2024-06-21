package main

import (
	"context"
	"fmt"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) Read(ctx context.Context, in *pb.BlogID) (*pb.Blog, error) {
	log.Printf("Read was invoked %v", in)

	for _, blog := range BlogList {
		if blog.ID == in.GetId() {
			return &pb.Blog{
				Id:       blog.ID,
				AuthorId: blog.AuthorID,
				Title:    blog.Title,
				Content:  blog.Content,
			}, nil
		}
	}

	return nil, status.Errorf(
		codes.InvalidArgument,
		fmt.Sprintf("Received blog from id: %s `not found`", in.GetId()),
	)
}
