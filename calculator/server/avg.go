package main

import (
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{Result: float32(sum) / float32(count)})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		sum += req.N
		count += 1
	}
}
