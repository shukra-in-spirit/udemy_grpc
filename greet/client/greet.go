package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "github.com/shukra-in-spirit/udemy_grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Suporno",
	})

	if err != nil {
		log.Fatalf("Could not greet %v,\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}