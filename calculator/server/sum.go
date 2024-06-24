package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.GetFirstNum() + in.GetSecondNum(),
	}, nil
}
