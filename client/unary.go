package main

import(
	proto "www.github.com/amanraj8848/gRPC-golang/proto"
	"context"
	"log"
	"time"
	
)

func callSayHello(client proto.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v", err)
	}
	log.Printf("Response from Server: %v", res.Message)
}