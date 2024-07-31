package presentation

import (
	"api_sample/internal/todo/domain"
)

type TodoResponse struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TodoListResponse struct {
	Todos []TodoResponse `json:"todos"`
}

func ConvertTodoResponse(todo domain.Todo) TodoResponse {
	return TodoResponse{
		Id:     todo.Id.ToInt(),
		UserId: todo.UserId.ToInt(),
		Title:  todo.Title.ToString(),
		Status: todo.Status.ToString(),
	}
}

func ConvertTodoListResponse(todos []domain.Todo) TodoListResponse {
	var todoResponses []TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, ConvertTodoResponse(todo))
	}
	return TodoListResponse{Todos: todoResponses}
}

type TodoRequest struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type NewTodoRequest struct {
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
