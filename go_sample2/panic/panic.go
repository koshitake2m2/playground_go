package main

import "fmt"

func main() {
	catchPanic()
}

func willPanic() {
	panic("This function will panic")
}

func catchPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()
	willPanic()
}
