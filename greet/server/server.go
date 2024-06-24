package main

import pb "github.com/saeedjhn/go-grpc/greet/goproto"

type Server struct {
	pb.GreetServiceServer
}
