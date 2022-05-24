package main

import (
	"io"
	"log"

	pb "github.com/shukra-in-spirit/udemy_grpc/calculator/proto"
)

func (*Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average was invoked")

	var sum float32 = 0
	var i float32 = 0
	var res float32 = 0.0

	for {
		i++
		req, err := stream.Recv()
		if err == io.EOF {
			res = sum / i
			return stream.SendAndClose(&pb.AverageResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Number %f: %v\n", i, req)
		sum += req.Num
	}
}
