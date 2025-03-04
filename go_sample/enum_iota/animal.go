package main

import "fmt"

type Animal int64

const (
	Dog Animal = iota + 1
	Cat
	Bird
)

var m = map[string]Animal{
	"DOG":  Dog,
	"CAT":  Cat,
	"BIRD": Bird,
}

func ForceValueOf(s string) Animal {
	for k, v := range m {
		if k == s {
			return v
		}
	}
	panic(fmt.Sprintf("Invalid Animal: %s", s))
}

func PointerValueOf(s string) *Animal {
	for k, v := range m {
		if k == s {
			return &v
		}
	}
	return nil
}

func (a Animal) String() string {
	for k, v := range m {
		if a == v {
			return k
		}
	}
	panic(fmt.Sprintf("Invalid Animal: %d", a))
}

func main() {
	fmt.Println(Dog)
	fmt.Println(Cat)
	fmt.Println(Bird)

	fmt.Println(ForceValueOf("DOG"))
	fmt.Println(ForceValueOf("CAT"))
	fmt.Println(ForceValueOf("BIRD"))

	dog1 := ForceValueOf("DOG")
	dog2 := ForceValueOf("DOG")
	fmt.Println(dog1 == dog2) // true

	pDog1 := PointerValueOf("DOG")
	pDog2 := PointerValueOf("DOG")
	fmt.Println(pDog1 == pDog2)   // false
	fmt.Println(*pDog1 == *pDog2) // true

	fmt.Println(ForceValueOf("X")) // panic
}
