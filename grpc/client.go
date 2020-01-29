package main

import (
	"context"
	"gin-server-api/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	Address1 = "0.0.0.0:9090"
)
func main() {
	conn, err := grpc.Dial(Address1, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	//初始化客户端
	c := helloworld.NewHelloWorldClient(conn)
	//调用SayHello 方法
	res, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "GeekGhc"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Message)
}
