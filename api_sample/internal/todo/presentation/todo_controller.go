package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
}

/*
curl -X GET http://localhost:8080/todos/
*/
func (tc TodoController) listTodos(c echo.Context) error {
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
func (tc TodoController) showTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

/*
curl -X POST http://localhost:8080/todos/ -H "Content-Type: application/json" -d '{"userId":1, "title":"title", "status":"waiting"}'
*/
func (tc TodoController) createTodo(c echo.Context) error {
	return c.String(http.StatusOK, "post todos")
}

/*
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"id":1, "userId":1, "title":"title", "status":"waiting"}'
*/
func (tc TodoController) updateTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

/*
curl -X DELETE http://localhost:8080/todos/1
*/
func (tc TodoController) deleteTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func (tc TodoController) RegisterTodoController(e *echo.Echo) {
	e.GET("/todos", tc.listTodos)
	e.GET("/todos:id", tc.showTodo)
	e.POST("/todos", tc.createTodo)
	e.PUT("/todos/:id", tc.updateTodo)
	e.DELETE("/todos/:id", tc.deleteTodo)
}
