package main

import (
	"context"
	"fmt"

	"example.com/aaa/greet"
)

type GreetServiceImpl struct {
	greet.UnimplementedGreetServiceServer
}

func (s *GreetServiceImpl) SayHello(_ context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	return &greet.GreetResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	service := &GreetServiceImpl{}
	fmt.Println(service.SayHello(context.Background(), &greet.GreetRequest{Name: "Bob"}))
}
