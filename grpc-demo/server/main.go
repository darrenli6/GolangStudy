package main

import (
	"context"
	pb "github.com/darrenli6/grpc-demo/proto"
	"google.golang.org/grpc"
	"net"
)

type GreeterServer struct {
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + r.Name}, nil
}

func main() {

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":8080")

	server.Serve(lis)

}
