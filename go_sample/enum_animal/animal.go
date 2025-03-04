package main

import "fmt"

type AnimalType struct {
	value string
}

var dog = AnimalType{"Dog"}
var cat = AnimalType{"Cat"}
var bird = AnimalType{"Bird"}
var animalTypes map[string]AnimalType = map[string]AnimalType{
	dog.value:  dog,
	cat.value:  cat,
	bird.value: bird,
}

func Dog() AnimalType {
	return dog
}
func Cat() AnimalType {
	return cat
}
func Bird() AnimalType {
	return bird
}

func Values() []AnimalType {
	vs := make([]AnimalType, 0, len(animalTypes))
	for _, v := range animalTypes {
		vs = append(vs, v)
	}
	return vs
}

var animals map[string]AnimalType = map[string]AnimalType{
	dog.value:  dog,
	cat.value:  cat,
	bird.value: bird,
}

func ValueOf(value string) AnimalType {
	if v, ok := animals[value]; ok {
		return v
	}
	panic(fmt.Sprintf("Invalid AnimalType: %s", value))
}

func main() {
	fmt.Println(Values())
	fmt.Println(ValueOf("Dog"))
	// fmt.Println(ValueOf("Dog2")) // panic
}
