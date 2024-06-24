package main

import (
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
)

type Server struct {
	pb.CalculatorServiceServer
}
