package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"log"
)

func (s *Server) Create(ctx context.Context, in *pb.Blog) (*pb.BlogID, error) {
	log.Printf("Create was invoked with %v", in)

	BlogList = append(
		BlogList,
		&Blog{
			ID:       uuid.New().String(),
			AuthorID: in.AuthorId,
			Title:    in.Title,
			Content:  in.Content,
		},
	)

	return &pb.BlogID{
		Id: BlogList[len(BlogList)-1].ID, // Give last item
	}, nil
}
