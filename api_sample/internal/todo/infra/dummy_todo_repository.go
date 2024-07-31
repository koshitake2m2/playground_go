package infra

import (
	baseDomain "api_sample/internal/base/domain"
	"api_sample/internal/todo/domain"
)

type DummyTodoRepository struct {
}

func (dtr DummyTodoRepository) FindAll() []domain.Todo {
	dummyUsers := []domain.Todo{
		{
			Id:     domain.NewTodoId(1),
			UserId: baseDomain.NewUserId(1),
			Title:  domain.NewTodoTitle("title"),
			Status: domain.NewTodoStatus("waiting"),
		},
		{
			Id:     domain.NewTodoId(2),
			UserId: baseDomain.NewUserId(2),
			Title:  domain.NewTodoTitle("title2"),
			Status: domain.NewTodoStatus("doing"),
		},
	}
	return dummyUsers
}
