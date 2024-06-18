package main

import (
	"context"
	"fmt"
	pb "github.com/saeedjhn/go-grpc/greet/proto"
	"log"
)

type Server struct {
	pb.GreetServiceServer
}

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", in.GetFirstName()),
	}, nil
}
