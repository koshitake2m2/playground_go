package main

import (
	"log"
	"oapi_codegen_sample/api"
	"oapi_codegen_sample/petstore"

	"github.com/labstack/echo/v4"
)

func main() {
	server := api.NewServer()

	e := echo.New()

	petstore.RegisterHandlers(e, server)

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start("0.0.0.0:8080"))
}
