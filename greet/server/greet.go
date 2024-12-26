package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/saeedjhn/go-grpc/greet/goproto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", in.GetFirstName()),
	}, nil
}
