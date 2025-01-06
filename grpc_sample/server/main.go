package main

import (
	"context"
	"fmt"
	"net"

	"example.com/aaa/greet"
	"google.golang.org/grpc"
)

type GreetServiceImpl struct {
	greet.UnimplementedGreetServiceServer
}

func (s *GreetServiceImpl) SayHello(_ context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	return &greet.GreetResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	greetService := &GreetServiceImpl{}

	s := grpc.NewServer()
	greet.RegisterGreetServiceServer(s, greetService)

	lis, _ := net.Listen("tcp", ":50051")
	fmt.Printf("Server started at %s\n", lis.Addr().String())
	s.Serve(lis)

}
