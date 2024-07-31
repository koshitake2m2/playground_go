package presentation

type TodoResponse struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Status string `json:"status"` // waiting, doing, done
}

type TodoListResponse struct {
	Todos []TodoResponse `json:"todos"`
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
