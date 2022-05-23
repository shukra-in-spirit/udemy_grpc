package main

import (
	"context"
	"io"

	pb "github.com/shukra-in-spirit/udemy_grpc/calculator/proto"
	log "github.com/sirupsen/logrus"
)

func doFactorize(c pb.CalculatorServiceClient) {
	log.Println("Factorize was invoked")

	req := &pb.FactorRequest{
		Num: 120,
	}

	stream, err := c.Factors(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Factors: %v\n", err)
	}
	var d int
	d = 0
	for {
		d++
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Factor %d: %d\n", d, msg.Result)
	}
}
