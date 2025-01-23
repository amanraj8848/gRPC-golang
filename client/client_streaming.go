package main

import(
	"log"
	"context"
	"time"
	proto "www.github.com/amanraj8848/gRPC-golang/proto"
)

func callSayHelloClientStreaming(client proto.GreetServiceClient, names *proto.NameList){
	log.Printf("Client streaming Request: %v", names.Name)
	stream,err:= client.SayHelloClientStreaming(context.Background())
	if err!= nil{
		log.Fatalf("Error while calling SayHelloClientStreaming: %v", err);
	}
	for _,name := range names.Name{
		if err := stream.Send(&proto.HelloRequest{
			Name: name,
		}); err != nil{
			log.Fatalf("Error while sending data to server: %v", err)
		}
		time.Sleep(2000 * time.Millisecond)
	}
	res,err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("Error while receiving response from server: %v", err)
	}
	log.Printf("Response from Server: %v", res.Message)
}