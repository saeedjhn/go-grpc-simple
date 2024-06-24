package main

import (
	"context"
	"fmt"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function was invoked with %v\n", in)

	number := in.N

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number))
	}

	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
