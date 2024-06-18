package main

import (
	"context"
	pb "github.com/saeedjhn/go-grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "saeed",
		LastName:  "jahanakbari",
		Email:     "saeedjahanakbari@gmail.com",
	})
	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Greeting %v", res.Result)
}
