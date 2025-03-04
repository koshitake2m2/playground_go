package main

import "fmt"

func run(fn func(a int, b int) error) {
	err := fn(1, 2)
	if err != nil {
		panic("error")
	}
}

func main() {
	var res int
	run(func(a int, b int) error {
		res = a + b
		return nil
	})
	fmt.Println(res)
}
