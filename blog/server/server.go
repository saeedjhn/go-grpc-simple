package main

import (
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
)

type Server struct {
	pb.BlogServiceServer
}
