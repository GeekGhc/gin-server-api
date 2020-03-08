package main

import (
	"context"
	"gin-server-api/pb/user"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const PORT = ":50051"

type server struct{}

//注册用户
func (server *server) Register(ctx context.Context, req *user.RegisterUserRequest) (*user.UserResponse, error) {
	log.Printf("receive from client: %s", req.Username)
	createdAt := &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   0,
	}
	return &user.UserResponse{
		Username:  "fff",
		Email:     "aaa",
		CreatedAt: createdAt,
	}, nil
}

//用户列表
func (server *server) UserList(ctx context.Context, empty *user.Empty) (*user.UserListResponse, error) {
	createdAt := &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   0,
	}
	user1 := &user.User{
		Username:  "name1",
		Email:     "email1",
		CreatedAt: createdAt,
	}
	return &user.UserListResponse{
		User: []*user.User{user1},
	}, nil
}

//用户流
func (server *server) UserListStream(ctx context.Context, stream *user.UserStreamRequest) (*user.UserResponse, error) {
	createdAt := &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   0,
	}
	return &user.UserResponse{
		Username:  "name2",
		Email:     "email2",
		CreatedAt: createdAt,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
