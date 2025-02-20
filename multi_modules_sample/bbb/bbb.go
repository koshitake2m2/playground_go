package bbb

import (
	"fmt"

	"example.com/aaa/aa2"

	"example.com/aaa"
	// could not import example.com/aaa/internal/internalaaa (invalid use of internal package "example.com/aaa/internal/internalaaa")compilerBrokenImport
	// "example.com/aaa/internal/internalaaa"

	// "example.com/aaa/cmd" imported as main and not used compiler UnusedImport
	// "example.com/aaa/cmd"
	"example.com/aaa/cmd/cmdutil"

	// "example.com/ccc" // import cycle not allowedgo list
	"example.com/ccc/cc2"
)

func PrintAB() {
	aaa.PrintA()
	aa2.PrintA3()
	fmt.Println("B")

	// internalaaa.PrintInternalAaa()
	// cmd.Hello()

	// Oh, we can import package in cmd which is not main package.
	cmdutil.Hello()

	// ccc.PrintC()
	cc2.PrintC2()
}

func PrintB() {
	fmt.Println("B")
}
