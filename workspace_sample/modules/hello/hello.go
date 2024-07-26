package hello

import (
	"fmt"
	"playground_go/calc"
)

func Hello() {
	fmt.Println("hello!!")
}

func UseAdd() {
	fmt.Println(calc.Add(1, 2))
}
