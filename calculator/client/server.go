package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
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

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream %v", err)
	}

	waitC := make(chan struct{})

	go func() {
		reqs := []int32{1, 5, 7, 2, 99, 101, 11, 202, 500, 440, 60}

		for _, req := range reqs {
			log.Printf("Sending number %d", req)
			stream.Send(&pb.MaxRequest{N: req})
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving %v", err)
				break
			}

			log.Printf("Receiving new maximum %d\n", res.Result)
		}

		close(waitC)
	}()

	<-waitC
}

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
