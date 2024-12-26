package main

import (
	"context"
	"log"

	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{N: n})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			return
		}

		log.Fatalf("A non GRPC error: %v\n", err)
	}

	log.Printf("Sqrt %d is: %f\n", n, res.Result)
}
