package main

import (
	"context"
	"flag"
	pb "github.com/darrenli6/GolangStudy/GoAdvanced/grpc/proto"
	"google.golang.org/grpc"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "9090", "启动端口号")
	flag.Parse()
}

type GreeterServer struct {
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello, world"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
