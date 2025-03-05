package main

import (
	"fmt"
	"reflect"
)

type (
	ValueObject[T any] interface {
		Value() T
		Equals(other ValueObject[T]) bool
		String() string
	}

	valueObject[T any] struct {
		value T
	}
)

func NewValueObject[T any](v T) ValueObject[T] {
	return &valueObject[T]{value: v}
}

func (v *valueObject[T]) Value() T {
	return v.value
}

// FIXME: You can compare two different type value objects.
func (v *valueObject[T]) Equals(other ValueObject[T]) bool {
	return reflect.DeepEqual(v.Value(), other.Value())
}

func (v *valueObject[T]) String() string {
	return fmt.Sprintf("%v", v.value)
}
