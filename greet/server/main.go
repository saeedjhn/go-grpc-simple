package main

import (
	pb "github.com/saeedjhn/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const addr = ":50000"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening  on %s\n", addr)

	s := grpc.NewServer()
	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}
