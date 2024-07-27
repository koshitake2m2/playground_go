package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + ": Woof!"
}

type MySingleton struct {
	Name string
}

type AnimalService struct {
	Animal      Animal
	MySingleton MySingleton
}

func (a AnimalService) Run() {
	fmt.Println(a.MySingleton.Name)
	fmt.Println(a.Animal.Speak())
}
