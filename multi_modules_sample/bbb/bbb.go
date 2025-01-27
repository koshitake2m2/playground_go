package bbb

import (
	"example.com/aaa/aa2"
	"fmt"

	"example.com/aaa"
)

func PrintAB() {
	aaa.PrintA()
	aa2.PrintA3()
	fmt.Println("B")
}
