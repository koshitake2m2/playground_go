package main

import (
	"context"

	"connectrpc.com/connect"
	"example.com/aaa/gen/proto/greet"
)

type GreetServiceImpl struct {
}

func (s *GreetServiceImpl) SayHello(_ context.Context, req *connect.Request[greet.GreetRequest]) (*connect.
	Response[greet.GreetResponse], error) {
	r := &greet.GreetResponse{Message: "" + req.Msg.Name}
	res := connect.NewResponse(r)
	return res, nil
}
