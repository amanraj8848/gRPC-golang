package main

import(
	"log"
	"context"
	"io"
	"time"
	proto "www.github.com/amanraj8848/gRPC-golang/proto"
)

func callSayHelloBiStreaming(client proto.GreetServiceClient, names *proto.NameList){
	log.Printf("Client Streaming started %v", names.Name)
	stream,err := client.SayHelloBiStreaming(context.Background())
	if err!= nil{
		log.Fatalf("Error while streaming %v",err)
	}
	waitc := make(chan struct{})
	go func(){
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving data from SayHelloBiStreaming: %v", err)
			}
			log.Printf("Streaming Response from Server: %v", res.Message)
		}
		close(waitc)
		log.Printf("Streaming Finished")
	}()

	for _,name := range names.Name{
		err :=stream.Send(&proto.HelloRequest{
			Name: name,
		})
		if err != nil{
			log.Fatalf("Error while sending data to server: %v", err)
		}
		time.Sleep(2500 * time.Millisecond)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")
	
}