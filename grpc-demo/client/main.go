package main

import (
	"context"
	pb "github.com/darrenli6/grpc-demo/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, _ := grpc.Dial(":8080", grpc.WithInsecure())

	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	_ = SayHello(client)
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Darren"})

	log.Panicf("client sayhello resp : %v", resp.Message)
	return nil
}
