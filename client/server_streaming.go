package main

import (
	"context"
	"io"
	"log"

	proto "www.github.com/amanraj8848/gRPC-golang/proto"
)

func callSayHelloServerStreaming(client proto.GreetServiceClient,names *proto.NameList){
	log.Printf("Client streaming Request: %v", names.Name)
	stream, err := client.SayHelloServerStreaming(context.Background(),names)
	if err != nil {
		log.Fatalf("Error while calling SayHelloServerStreaming: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving data from SayHelloServerStreaming: %v", err)
		}
		log.Println(res)
	}

	log.Printf("Streaming Finished")
}

