package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"io"
	"log"
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

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimesRequest{N: 25}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Primes is: %d\n", res.Result)
	}
}
