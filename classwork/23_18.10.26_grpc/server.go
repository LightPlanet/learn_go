package main

import (
	"context"
	"fmt"
	"grpc/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

type authServiceServer struct {
	user.UnimplementedAuthServiceServer
}

func (s *authServiceServer) CheckUserCredo(ctx context.Context, req *user.AuthRequest) (*user.AuthResponse, error) {
	log.Println("I got message!")
	if req.Email == "ivanov@example.com" && req.Password == "123" {
		return &user.AuthResponse{
			Answer: "YES!",
		}, nil
	} else {
		return &user.AuthResponse{
			Answer: "NO!",
		}, nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer := grpc.NewServer()
	user.RegisterAuthServiceServer(grpcServer, &authServiceServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
		return
	}
}
