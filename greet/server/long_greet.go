package main

import (
	"fmt"
	pb "github.com/saeedjhn/go-grpc/greet/goproto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		res += fmt.Sprintf("Hello %s\n", req.GetFirstName())
		log.Printf("Sending %v\n", req.GetFirstName())
	}
}
