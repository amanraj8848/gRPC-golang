package main

import (
	"io"
	"log"
	proto "www.github.com/amanraj8848/gRPC-golang/proto"
)

func( s *server) SayHelloClientStreaming (stream proto.GreetService_SayHelloClientStreamingServer) error{
	var messages []string
	for{
		req,err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&proto.MessageList{
				Message: messages,
			})
		}
		if err != nil{
			log.Fatalf("Error while receiving data from client: %v", err)
		}
		messages = append(messages, "Hello " + req.Name)
		log.Printf("Data Received from Client: %v", req.Name)
	}

}