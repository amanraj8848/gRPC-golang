package main


import (
	"context"
	proto "www.github.com/amanraj8848/gRPC-golang/proto"
)

func (s *server) SayHello(ctx context.Context, req *proto.NoParam) ( *proto.HelloReply,error) {
	return &proto.HelloReply{
		Message: "Hello from Server",
	},nil
}