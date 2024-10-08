package main

import (
	"io"
	"log"

	pb "github.com/LOVESH1808/grpc-one/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	log.Printf("in server side")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{
				Message: messages,
			})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with names: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}

}
