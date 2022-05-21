package main

import (
	"context"
	
	log "github.com/sirupsen/logrus"
	pb "github.com/shukra-in-spirit/udemy_grpc/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Result: in.FirstNum + in.SecondNum,
	}, nil
}