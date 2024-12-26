package main

import (
	"context"
	"log"

	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNum:  1,
		SecondNum: 5,
	})

	if err != nil {
		log.Fatalf("Could not sum %v\n", err)
	}

	log.Printf("Sum is: %d\n", res.Result)
}
