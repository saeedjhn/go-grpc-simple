package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"log"
	"time"
)

type Server struct {
	pb.CalculatorServiceServer
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.GetFirstNum() + in.GetSecondNum(),
	}, nil
}

func (s *Server) Sub(ctx context.Context, in *pb.SubRequest) (*pb.SubResponse, error) {
	log.Printf("Sub function was invoked with %v\n", in)

	return &pb.SubResponse{
		Result: in.GetFirstNum() - in.GetSecondNum(),
	}, nil
}

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	// running the for loop from 1 to N
	for i := 2; i <= int(in.N); i++ {
		time.Sleep(1 * time.Second)

		// flag which will confirm that the number is Prime or not
		isPrime := true

		// This for loop is checking that the current number have
		// any divisible other than 1
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		// if the value of the flag is false then the number is not prime
		// and we are not printing it.
		if isPrime {
			stream.Send(&pb.PrimesResponse{
				Result: int32(i),
			})
		}
	}

	return nil
}
