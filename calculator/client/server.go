package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/proto"
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

func doMines(c pb.CalculatorServiceClient) {
	log.Println("doMines was invoked")

	res, err := c.Mines(context.Background(), &pb.MinesRequest{
		FirstNum:  8,
		SecondNum: 5,
	})

	if err != nil {
		log.Fatalf("Could not mines %v\n", err)
	}

	log.Printf("Mines is: %d\n", res.Result)
}
