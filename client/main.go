package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "www.github.com/amanraj8848/gRPC-golang/proto"
)

const (
	PORT = ":2500"
)



func main() {
	conn, err := grpc.Dial("localhost" + PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(" failed to connect: %v", err)
	}	
	defer conn.Close() 
	client := proto.NewGreetServiceClient(conn) // registring the client with the server	
	names := &proto.NameList{
		Name: []string{"Buddha","Lao Tzu","Nirvana","Socrates","Rama","Krishna"},
	}
	// callSayHelloServerStreaming(client,names);
	// callSayHello(client)
	callSayHelloClientStreaming(client,names)
}