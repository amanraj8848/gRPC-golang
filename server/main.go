package main

import(
	"context"
	proto "github.com/amanraj8848/gRPC-golang/proto"
	"errors"
	"net"
	"log"
	"google.golang.org/grpc"
)

type server struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp",":2500")
	if err != nil {
		panic(err);
	}
	log.Printf("Server starting at PORT 2500 and address %v",lis.Addr())
	grpcServer := grpc.NewServer()
	
	proto.RegisterGreetServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		panic(err);
	}

}
// func handler
func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply,error){
	return &proto.HelloReply{
		message: "Hello from Server"
	}, nil;
}