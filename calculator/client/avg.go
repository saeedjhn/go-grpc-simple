package main

import (
	"context"
	"log"

	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	reqs := []int32{11, 52, 1, 82, 62, 101, 15}
	//reqs := []*pb.AvgRequest{
	//	{N: 22},
	//	{N: 52},
	//	{N: 101},
	//}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending: %v", req)
		stream.Send(&pb.AvgRequest{N: req})
		//stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg %v", err)
	}

	log.Printf("Avg %f\n", res.Result)
}
