package main

import (
	"fmt"
	pb "github.com/saeedjhn/go-grpc/greet/goproto"
	"log"
	"time"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v", in)

	for i := 0; i < 15; i++ {
		time.Sleep(1 * time.Second)

		res := fmt.Sprintf("Hello %s, number %d", in.GetFirstName(), i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
