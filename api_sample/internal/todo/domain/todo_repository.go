package domain

type TodoRepository interface {
	FindAll() []Todo
}
