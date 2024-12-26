package main

import (
	"log"

	pb "github.com/saeedjhn/go-grpc/calculator/goproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "0.0.0.0:50001"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	log.Printf("Connection is successful on %s\n", addr)

	c := pb.NewCalculatorServiceClient(conn)
	doSum(c)
	doSub(c)
	doPrimes(c)
	doAvg(c)
	doMax(c)
	doSqrt(c, 10)
}
