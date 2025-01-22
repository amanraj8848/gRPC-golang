package main

import (
	"log"
	"time"

	proto "www.github.com/amanraj8848/gRPC-golang/proto"
)


func (s *server) SayHelloServerStreaming (req *proto.NameList,stream proto.GreetService_SayHelloServerStreamingServer) error{
	log.Printf("Server streaming Request Started for : %v", req.Name)
	for i,names := range req.Name{
		res := &proto.HelloReply{
			Message: "Hello " + names,
		}
		if err:= stream.Send(res); err != nil{	
			log.Fatalf("Error while sending data to client: %v", err) 
		}
		log.Printf("Data Sent for %v", i+1)
		time.Sleep(2000 * time.Millisecond)
	}
	log.Printf("Streaming Finished for %v", req.Name)
	return nil
}
