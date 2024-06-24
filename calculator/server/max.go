package main

import (
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function invoked")

	var maximum int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client streaming %v", err)
		}

		if number := req.N; number > maximum {
			maximum = number
			err = stream.Send(&pb.MaxResponse{Result: maximum})
			if err != nil {
				log.Fatalf("Error while sending data to client %v", err)
			}
		}

	}
}
