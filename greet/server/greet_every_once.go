package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/saeedjhn/go-grpc/greet/goproto"
)

func (s *Server) GreetEveryOnce(stream pb.GreetService_GreetEveryOnceServer) error {
	log.Println("GreetEveryOnce was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client streaming %v\n", err)
		}

		res := fmt.Sprintf("Hello %s", req.GetFirstName())
		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("Error while sending data to client %v\n", err)
		}
	}
}
