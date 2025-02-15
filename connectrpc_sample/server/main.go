package main

import (
	"connectrpc.com/connect"
	"context"
	"example.com/aaa/gen/greet"
	"example.com/aaa/gen/greet/greetconnect"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

type GreetServiceImpl struct {
}

func (s *GreetServiceImpl) SayHello(_ context.Context, req *connect.Request[greet.GreetRequest]) (*connect.
	Response[greet.GreetResponse], error) {
	r := &greet.GreetResponse{Message: "" + req.Msg.Name}
	res := connect.NewResponse(r)
	return res, nil
}

func NewGreetServiceImpl() greetconnect.GreetServiceHandler {
	return &GreetServiceImpl{}
}

func main() {

	gs := NewGreetServiceImpl()
	mux := http.NewServeMux()
	path, handler := greetconnect.NewGreetServiceHandler(gs)
	mux.Handle(path, handler)
	fmt.Printf("Server started at %s\n", "localhost:50051")
	http.ListenAndServe(
		"localhost:50051",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
