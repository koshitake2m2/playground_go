package domain

import (
	baseDomain "api_sample/internal/base/domain"
)

type TodoId struct {
	value int
}

func (v TodoId) ToInt() int {
	return v.value
}

func NewTodoId(value int) TodoId {
	return TodoId{value: value}
}

type TodoTitle struct {
	value string
}

func (v TodoTitle) ToString() string {
	return v.value
}

func NewTodoTitle(value string) TodoTitle {
	return TodoTitle{value: value}
}

type TodoStatus struct {
	value string // waiting, doing, done
}

func (v TodoStatus) ToString() string {
	return v.value
}

func NewTodoStatus(value string) TodoStatus {
	return TodoStatus{value: value}
}

type Todo struct {
	Id     TodoId
	UserId baseDomain.UserId
	Title  TodoTitle
	Status TodoStatus
}
