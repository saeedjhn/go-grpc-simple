package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"log"
)

func (s *Server) Sub(ctx context.Context, in *pb.SubRequest) (*pb.SubResponse, error) {
	log.Printf("Sub function was invoked with %v\n", in)

	return &pb.SubResponse{
		Result: in.GetFirstNum() - in.GetSecondNum(),
	}, nil
}
