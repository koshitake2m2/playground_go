package presentation

import (
	basePresentation "api_sample/internal/base/presentation"
	"api_sample/internal/todo/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	AuthenticationHelper basePresentation.AuthenticationHelper
	TodoUsecase          usecase.TodoUsecase
}

/*
curl -X GET http://localhost:8080/todos/ -b "SESSION_ID=SESSION_ID_USER_1"
*/
func (tc TodoController) listTodos(c echo.Context) error {
	user, err := tc.AuthenticationHelper.Authenticate(c)
	if err != nil {
		return err
	}
	fmt.Println("login user id:", user.UserId)

	todos := tc.TodoUsecase.ListTodos()
	response := ConvertTodoListResponse(todos)
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

func (tc TodoController) RegisterController(e *echo.Echo) {
	e.GET("/todos", tc.listTodos)
	e.GET("/todos/:id", tc.showTodo)
	e.POST("/todos", tc.createTodo)
	e.PUT("/todos/:id", tc.updateTodo)
	e.DELETE("/todos/:id", tc.deleteTodo)
}
