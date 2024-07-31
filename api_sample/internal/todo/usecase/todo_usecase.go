package usecase

import (
	baseDomain "api_sample/internal/base/domain"
	"api_sample/internal/todo/domain"
)

type TodoUsecase struct {
	IdGenerator    baseDomain.IdGenerator
	TodoRepository domain.TodoRepository
}

func (tu TodoUsecase) ListTodos() []domain.Todo {
	return tu.TodoRepository.FindAll()
}
