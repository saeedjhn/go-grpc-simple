package main

import (
	"context"
	"fmt"
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
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

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function invoked")

	var maximum int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client streaming %v", err)
		}

		if number := req.N; number > maximum {
			maximum = number
			err = stream.Send(&pb.MaxResponse{Result: maximum})
			if err != nil {
				log.Fatalf("Error while sending data to client %v", err)
			}
		}

	}
}

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function was invoked with %v\n", in)

	number := in.N

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number))
	}

	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
