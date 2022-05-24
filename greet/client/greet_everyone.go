package main

import (
	"context"
	"io"
	"time"

	pb "github.com/shukra-in-spirit/udemy_grpc/greet/proto"
	log "github.com/sirupsen/logrus"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Suporno"},
		{FirstName: "Test1"},
		{FirstName: "Test2"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)

			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Recieved: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
