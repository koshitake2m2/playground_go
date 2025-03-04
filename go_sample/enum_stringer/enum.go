package main

import "fmt"

//go:generate stringer -type TodoStatus
type TodoStatus int

const (
	New TodoStatus = iota
	InProgress
	Done
)

var todoStatusLabel map[TodoStatus]string = map[TodoStatus]string{
	New:        "NEW",
	InProgress: "IN_PROGRESS",
	Done:       "DONE",
}

func TodoStatusByLabel(label string) TodoStatus {
	for k, v := range todoStatusLabel {
		if v == label {
			return k
		}
	}
	panic(fmt.Sprintf("Invalid label: %s", label))
}

func (t TodoStatus) Label() string {
	if label, ok := todoStatusLabel[t]; ok {
		return label
	}
	panic(fmt.Sprintf("Invalid TodoStatus: %d", t))
}

//go:generate stringer -type AnimalType
type AnimalType int

const (
	Dog AnimalType = iota
	Cat
	Bird
)
