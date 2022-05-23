package main

import (
	pb "github.com/shukra-in-spirit/udemy_grpc/calculator/proto"
	log "github.com/sirupsen/logrus"
)

func (s *Server) Factors(in *pb.FactorRequest, stream pb.CalculatorService_FactorsServer) error {
	log.Printf("Factors was invoked: %v\n", in)

	N := in.Num
	var k int64
	k = 2
	for N > 1 {
		if N%k == 0 {
			stream.Send(&pb.CalculatorResponse{
				Result: k,
			})
			N = N / k
		} else {
			k++
		}
	}

	return nil
}
