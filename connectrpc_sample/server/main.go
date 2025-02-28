package main

import (
	"fmt"
	"net/http"

	"example.com/aaa/gen/proto/greet/greetconnect"
	"example.com/aaa/gen/proto/pdf/pdfconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewGreetServiceImpl() greetconnect.GreetServiceHandler {
	return &GreetServiceImpl{}
}

func main() {

	gs := NewGreetServiceImpl()
	ps := NewPdfService()
	mux := http.NewServeMux()
	path, handler := greetconnect.NewGreetServiceHandler(gs)
	mux.Handle(path, handler)
	path2, hanlder2 := pdfconnect.NewPdfServiceHandler(ps)
	mux.Handle(path2, hanlder2)
	fmt.Printf("Server started at %s\n", "localhost:50051")
	http.ListenAndServe(
		"localhost:50051",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
