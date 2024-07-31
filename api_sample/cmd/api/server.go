package main

import (
	"api_sample/internal/todo/presentation"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("hh334554")
	return c.String(http.StatusOK, id)
}
func main() {
	e := echo.New()
	tc := presentation.TodoController{}
	tc.RegisterTodoController(e)

	e.GET("/hello", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
