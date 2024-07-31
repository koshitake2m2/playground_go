package main

import (
	"net/http"

	"api_sample/cmd/api/di"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
func main() {
	e := echo.New()
	dependencies, _ := di.Initialize()
	dependencies.TodoController.RegisterController(e)

	e.GET("/hello", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
