package main

import "fmt"

type Id[T any, V any] struct {
	value V
}

func NewId[T any, V any](value V) Id[T, V] {
	return Id[T, V]{value: value}
}

func (id *Id[T, V]) Value() V {
	return id.value
}

type User struct {
	id Id[User, string]
}

func (u *User) Id() Id[User, string] {
	return u.id
}

type Animal struct{}
type Shape struct{}

type AppleId string
type BananaId string

func main() {
	userId1 := NewId[User, string]("1")
	userId1_2 := NewId[User]("1")
	userId2 := NewId[User]("2")
	animalId1 := NewId[Animal]("1")
	shapeId1 := NewId[Shape, uint64](1)

	fmt.Printf("userId1: %v\n", userId1.Value())
	fmt.Printf("animalId1: %v\n", animalId1.Value())
	fmt.Printf("shapeId1: %v\n", shapeId1.Value())
	fmt.Println("userId1 == userId1_2:", userId1 == userId1_2) // true
	fmt.Println("userId1 == userId2:", userId1 == userId2)     // false
	// fmt.Println("userId1 == animalId1:", userId1 == animalId1) // compile error
	// fmt.Println("userId1 == shapeId1:", userId1 == shapeId1) // compile error

	appleId1 := AppleId("1")
	appleId1_2 := AppleId("1")
	bananaId1 := BananaId("1")
	fmt.Printf("appleId1: %v\n", appleId1)
	fmt.Printf("bananaId1: %v\n", bananaId1)
	fmt.Println("appleId1 == appleId1_2:", appleId1 == appleId1_2) // true
	// fmt.Println("appleId == bananaId:", appleId1 == bananaId1) // compile error

}
