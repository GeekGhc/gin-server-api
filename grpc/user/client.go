package main

import (
	"context"
	"gin-server-api/pb/user"
	"google.golang.org/grpc"
	"log"
)

const (
	ClientAddress = "0.0.0.0:50051"
)

func main() {
	conn, err := grpc.Dial(ClientAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//初始化客户端
	c := user.NewUserServiceClient(conn)
	//调用Register方法
	res, err := c.Register(context.Background(), &user.RegisterUserRequest{
		Username: "geekghc",
		Email:    "geekghc@gmail.com",
		Password: "geek123",
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("username: ",res.Username," email: ",res.Email," created at: ",res.CreatedAt)
}
