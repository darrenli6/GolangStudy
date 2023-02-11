package main

import (
	pb "github.com/darrenli6/GolangStudy/GoBlog/tag-service/proto"
	"github.com/darrenli6/GolangStudy/GoBlog/tag-service/server"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
}
