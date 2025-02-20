package bbb

import (
	"fmt"

	"example.com/aaa/aa2"

	"example.com/aaa"
	// could not import example.com/aaa/internal/internalaaa (invalid use of internal package "example.com/aaa/internal/internalaaa")compilerBrokenImport
	// "example.com/aaa/internal/internalaaa"
)

func PrintAB() {
	aaa.PrintA()
	aa2.PrintA3()
	// internalaaa.PrintInternalAaa()
	fmt.Println("B")
}
