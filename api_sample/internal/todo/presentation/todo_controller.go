package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
curl -X GET http://localhost:8080/todos/
*/
func listTodos(c echo.Context) error {
	response := TodoListResponse{Todos: []TodoResponse{
		{Id: 1, UserId: 1, Title: "title", Status: "waiting"},
		{Id: 2, UserId: 2, Title: "title2", Status: "doing"},
	}}
	// return c.String(http.StatusOK, "aa")
	return c.JSON(http.StatusOK, response)
}

/*
curl -X GET http://localhost:8080/todos/1
*/
func showTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

/*
curl -X POST http://localhost:8080/todos/ -H "Content-Type: application/json" -d '{"userId":1, "title":"title", "status":"waiting"}'
*/
func createTodo(c echo.Context) error {
	return c.String(http.StatusOK, "post todos")
}

/*
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"id":1, "userId":1, "title":"title", "status":"waiting"}'
*/
func updateTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

/*
curl -X DELETE http://localhost:8080/todos/1
*/
func deleteTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func RegisterTodoController(e *echo.Echo) {
	e.GET("/todos", listTodos)
	e.GET("/todos:id", showTodo)
	e.POST("/todos", createTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)
}
