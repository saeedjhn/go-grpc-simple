package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"log"
)

func doSub(c pb.CalculatorServiceClient) {
	log.Println("doSub was invoked")

	res, err := c.Sub(context.Background(), &pb.SubRequest{
		FirstNum:  8,
		SecondNum: 5,
	})

	if err != nil {
		log.Fatalf("Could not sub %v\n", err)
	}

	log.Printf("Sub is: %d\n", res.Result)
}
