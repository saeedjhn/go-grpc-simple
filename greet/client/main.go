package main

import (
	"log"
	"time"

	pb "github.com/saeedjhn/go-grpc/greet/goproto"
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

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	doGreetManyTimes(c)
	doLongGreet(c)
	doGreetEveryOnce(c)
	doGreetWithDeadline(c, 4*time.Second)
}
