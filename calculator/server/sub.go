package main

import (
	"context"
	"log"

	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
)

func (s *Server) Sub(ctx context.Context, in *pb.SubRequest) (*pb.SubResponse, error) {
	log.Printf("Sub function was invoked with %v\n", in)

	return &pb.SubResponse{
		Result: in.GetFirstNum() - in.GetSecondNum(),
	}, nil
}
