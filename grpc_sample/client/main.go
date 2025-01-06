package main

import (
	"context"
	"fmt"
	"time"

	"example.com/aaa/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, _ := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	c := greet.NewGreetServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, _ := c.SayHello(ctx, &greet.GreetRequest{Name: "Alice"})

	fmt.Println(res.Message)
}
