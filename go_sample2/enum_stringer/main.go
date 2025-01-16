package main

import (
	"fmt"
)

func SayInt(a int) {
	fmt.Println(a)
}
func Say(a AnimalType) {
	fmt.Println(a)
}

func main() {
	s := New
	fmt.Println(s.String())
	fmt.Println(New == 0) // true
	// fmt.Println(New == Dog) // compile error

	Say(Dog)
	// Say(New) // compile error
	// SayInt(Dog) // compile error

	fmt.Println(TodoStatus(0))  // New
	fmt.Println(TodoStatus(-1)) // TodoStatus(-1)
	fmt.Println(TodoStatus(5))  // TodoStatus(5)

	fmt.Println(TodoStatusByLabel("NEW")) // New
	fmt.Println(New.Label())              // NEW

	var emptyTodoStatus TodoStatus
	// emptyTodoStatus = Dog // compile error
	emptyTodoStatus = 5
	fmt.Println(emptyTodoStatus) // TodoStatus(5)

}
