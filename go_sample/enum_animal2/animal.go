package main

import "fmt"

type AnimalType struct {
	value string
}

type Animal struct {
}

var animalEnum = Animal{}

func AnimalEnum() Animal {
	return animalEnum
}

var dog = AnimalType{"Dog"}
var cat = AnimalType{"Cat"}
var bird = AnimalType{"Bird"}
var animals map[string]AnimalType = map[string]AnimalType{
	dog.value:  dog,
	cat.value:  cat,
	bird.value: bird,
}

func (a Animal) Dog() AnimalType {
	return dog
}
func (a Animal) Cat() AnimalType {
	return cat
}
func (a Animal) Bird() AnimalType {
	return bird
}
func (a Animal) Values() []AnimalType {
	return []AnimalType{dog, cat, bird}
}
func ValueOfAnimalType(value string) AnimalType {
	if v, ok := animals[value]; ok {
		return v
	}
	panic(fmt.Sprintf("Invalid AnimalType: %s", value))
}

func main() {
	a := AnimalEnum().Dog()
	fmt.Println(a)
	fmt.Println(AnimalEnum().Values())
}
