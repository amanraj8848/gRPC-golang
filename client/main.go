package main

const (
	PORT = ":8080"
)

import(
	"log"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost" + PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(" failed to connect: %v", err)
	}	
	defer conn.Close() 

	client := pb.NewGreetServiceClient(conn)
}