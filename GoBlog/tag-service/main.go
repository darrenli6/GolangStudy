package main

import (
	pb "github.com/darrenli6/GolangStudy/GoBlog/tag-service/proto"
	"github.com/darrenli6/GolangStudy/GoBlog/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()

	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("net listen err %v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("s.Serve error  %v", err)
	}
}
