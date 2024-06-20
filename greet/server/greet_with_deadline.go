package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/saeedjhn/go-grpc/greet/goproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v", in)

	//for i := 0; i < 3; i++ {
	time.Sleep(3 * time.Second)

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		log.Println("The client canceled the request!.")
		return nil, status.Error(codes.Canceled, "The client canceled the request!")
	}

	//}

	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", in.GetFirstName()),
	}, nil
}
