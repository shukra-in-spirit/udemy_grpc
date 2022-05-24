package main

import (
	"context"
	"time"

	pb "github.com/shukra-in-spirit/udemy_grpc/calculator/proto"
	log "github.com/sirupsen/logrus"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.AverageRequest{
		{Num: 1},
		{Num: 1},
		{Num: 1},
		{Num: 1},
		{Num: 1},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending num: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving response from Average: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)
}
