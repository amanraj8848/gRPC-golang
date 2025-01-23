package main

import (
	"log"
	"io"
	"time"
	proto "www.github.com/amanraj8848/gRPC-golang/proto"
)


func (s *server) SayHelloBiStreaming (stream proto.GreetService_SayHelloBiStreamingServer) error{
	log.Printf("Bidirectional Streaming Request Started")
	for{
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			log.Fatalf("Error while receiving data from client: %v", err)
		}
		res := &proto.HelloReply{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil{
			log.Fatalf("Error while sending data to client: %v", err)
		}
		log.Printf("Data Sent to Client: %v", req.Name)
		time.Sleep(2000 * time.Millisecond)
	}
	log.Printf("Bidirectional Streaming Finished")
	return nil
}