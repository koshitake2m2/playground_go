package main

import (
	"fmt"
	"playground_go/calc"

	// use of internal package playground_go/calc/internal not allowed
	// "playground_go/calc/internal" // Not working
	"playground_go/hello"
)

func main() {
	hello.Hello()
	fmt.Println(calc.Add(1, 2))
	// internal.Minus(1, 2)
	fmt.Println("hello!!")
}
