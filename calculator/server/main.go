package main

import (
	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const addr = "0.0.0.0:50001"

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening  on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}
