package main

import "fmt"

func main() {
	age := NewUserAge(20)
	age2 := NewUserAge(20)
	age3 := NewUserAge(30)

	fmt.Println((age.Value()))
	fmt.Println(age.String())

	fmt.Println("age.Equals(age2): ", age.Equals(age2)) // true
	fmt.Println("age.Equals(age3): ", age.Equals(age3)) // false
	fmt.Println("age == age2: ", age == age2)           // false
	fmt.Println("age == age3: ", age == age3)           // false

	// FIXME
	year := NewYear(20)
	fmt.Println("age.Equals(age22): ", age.Equals(year)) // true. OMG!!

	// This is a compile error.
	// nextAge := NextAge(year) // compile error

}
