package main

import (
	"context"
	"fmt"
	"grpc/user"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	client := user.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	authResp, err := client.CheckUserCredo(ctx, &user.AuthRequest{
		Email:    "ivanov@example.com",
		Password: "123",
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(authResp.Answer)
}
