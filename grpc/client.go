package main

import (
	"context"
	"gin-server-api/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
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
}
