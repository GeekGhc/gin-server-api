package main

import (
	"context"
	"fmt"
	"gin-server-api/pb/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	Address = "0.0.0.0:9090"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Message: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func (s *server) LotsOfReplies(in *helloworld.HelloRequest, stream helloworld.HelloWorld_LotsOfRepliesServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&helloworld.HelloResponse{Message: fmt.Sprintf("%s %s %d", in.Name, "Reply", i)})
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	//服务注册
	helloworld.RegisterHelloWorldServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
