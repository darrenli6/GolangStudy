package main

import (
	"context"
	"flag"
	pb "github.com/darrenli6/GolangStudy/GoAdvanced/grpc/proto"
	"google.golang.org/grpc"
	"log"
)

var port string

func init() {
	flag.StringVar(&port, "p", "9090", "启动端口号")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	err := SayHello(client)
	if err != nil {
		log.Fatalf("client.say hello err %v", err)
	}
}

func SayHello(client pb.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "darren",
	})
	if err != nil {
		return err
	}
	log.Printf("client.say hello resp %s", resp.Message)
	return nil
}
