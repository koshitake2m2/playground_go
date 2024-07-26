package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

/*
curl -X GET 'http://localhost:1323/show?ids=3&ids=4'
*/
func show(c echo.Context) error {
	name := c.QueryParams().Get("name")
	ids := c.Request().URL.Query()["ids"]
	return c.String(http.StatusOK, "name:"+name+", ids:"+fmt.Sprintf("%v", ids))
}

/*
curl -d "name=Bob" -d "age=10" http://localhost:1323/save
*/
func saveForm(c echo.Context) error {
	name := c.FormValue("name")
	age := c.FormValue("age")
	return c.String(http.StatusOK, "name:"+name+", age:"+age)
}

type User struct {
	Name string
	Age  int
}

/*
curl -X POST -H "Content-Type: application/json" -d '{"name":"Bob", "age":10}' http://localhost:1323/users
*/
func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println(u)
	return c.JSON(http.StatusCreated, u)
}

/*
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Bob", "age":10}' http://localhost:1323/users/3
*/
func updateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println(id, u)
	return c.JSON(http.StatusOK, u)
}

/*
curl -X DELETE http://localhost:1323/users/
*/
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.GET("/show", show)
	e.GET("/users/:id", getUser)
	e.POST("/users", saveUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.POST("/save", saveForm)
	e.Logger.Fatal(e.Start(":1323"))
}
