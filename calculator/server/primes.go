package main

import (
	"log"
	"time"

	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
)

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
