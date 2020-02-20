package main

import (
	"context"
	"gin-server-api/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	ClientAddress = "0.0.0.0:9090"
)

func main() {
	conn, err := grpc.Dial(ClientAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//初始化客户端
	c := helloworld.NewHelloWorldClient(conn)
	//调用SayHello 方法
	res, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "gin-server"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Message)

	//调用 LotsOfReplies 方法
	stream, err := c.LotsOfReplies(context.Background(), &helloworld.HelloRequest{Name: "gin-server"})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv: %v", err)
		}
		log.Printf("%s", res.Message)
	}
}
