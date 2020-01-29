package main

import (
	"context"
	"fmt"
	"gin-server-api/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Message: fmt.Sprintln("%s", in.Name)}, nil
}

func (s *server) LotsOfReplies(in *helloworld.HelloRequest, stream helloworld.HelloWorld_LotsOfRepliesServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&helloworld.HelloResponse{Message: fmt.Sprintf("%s %s %d", in.Name, "Reply", i)})
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterHelloWorldServer(s, &server{})
	s.Serve(listen)

}
