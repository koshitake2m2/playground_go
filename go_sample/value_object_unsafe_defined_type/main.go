package main

import "fmt"

func main() {
	age := NewUserAge(20)
	age2 := NewUserAge(20)
	age3 := NewUserAge(30)

	fmt.Println("age.Value(): ", (age.Value()))
	fmt.Println("age.String(): ", age.String())
	fmt.Println("age: ", age)
	fmt.Println("age.Show(): ", age.Show())

	// fmt.Println("age.Equals(age2): ", age.Equals(age2)) // true
	// fmt.Println("age.Equals(age3): ", age.Equals(age3)) // false
	fmt.Println("age == age2: ", age == age2) // true
	fmt.Println("age == age3: ", age == age3) // false

	year := NewYear(20)
	// fmt.Println("age.Equals(year): ", age.Equals(year)) // compile error
	fmt.Println("age == year: ", age == year) // OMG! compile success
	nextAge := NextAge(year)                  // OMG! compile success
	fmt.Println("nextAge: ", nextAge)

}
