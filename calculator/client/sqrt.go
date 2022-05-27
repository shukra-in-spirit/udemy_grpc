package main

import (
	"context"
	"log"

	pb "github.com/shukra-in-spirit/udemy_grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Num: n})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("You probably sent a negative number")
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
		return
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
