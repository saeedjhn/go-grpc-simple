package main

import (
	pb "github.com/saeedjhn/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const addr = ":50000"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	log.Printf("Connection is successful on %s\n", addr)

	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
}
