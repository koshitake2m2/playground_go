package main

import (
	"log"
	"oapi_codegen_sample/api"
	_ "oapi_codegen_sample/docs"
	"oapi_codegen_sample/petstore"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"

	// middleware "github.com/oapi-codegen/echo-middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title                      Recommend Swaggo API
// @version                    1.0
// @host                       localhost:8080
// @BasePath                   /
func main() {
	server := api.NewServer()
	swagger, _ := api.GetSwagger()
	swagger.Servers = []*openapi3.Server{{URL: "http://localhost:8080"}}

	e := echo.New()

	// OpenAPIに基づいたバリデーションを行うミドルウェアを追加
	// TODO: 以下を追加すると, swagger-uiやhelloが404になってしまう
	// e.Use(middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
	// 	SilenceServersWarning: true,
	// }))

	petstore.RegisterHandlers(e, server)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/hello", hello)

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start("0.0.0.0:8080"))
}

// hello godoc
// @Summary Return "Hello, World!"
// @Description Return "Hello, World!"
// @Produce plain
// @Success 200 {string} string
// @Router /hello [get]
func hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
