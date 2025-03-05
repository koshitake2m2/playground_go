package main

import (
	"fmt"
)

type (
	ValueObject[T any] interface {
		Value() T
		// Equals(other S) bool
		String() string
		// Show() string
	}

	valueObject[T any] struct {
		value T
	}
)

func NewValueObject[T any](v T) ValueObject[T] {
	return valueObject[T]{value: v}
}

func (v valueObject[T]) Value() T {
	return v.value
}

// NOTE: You can use `==` operator instead of `Equals` method.
// func (v valueObject[T, S]) Equals(other S) bool {
// 	return reflect.DeepEqual(v.Value(), other.Value())
// }

func (v valueObject[T]) String() string {
	return fmt.Sprintf("%v", v.value)
}
