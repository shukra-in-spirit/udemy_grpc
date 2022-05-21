package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "github.com/shukra-in-spirit/udemy_grpc/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invoked")

	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		FirstNum: 10,
		SecondNum: 4,
	})

	if err != nil {
		log.Fatalf("Could not calculate %v,\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}