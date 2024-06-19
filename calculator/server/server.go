package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/calculator/proto"
	"log"
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

func (s *Server) Mines(ctx context.Context, in *pb.MinesRequest) (*pb.MinesResponse, error) {
	return &pb.MinesResponse{
		Result: in.GetFirstNum() - in.GetSecondNum(),
	}, nil
}
