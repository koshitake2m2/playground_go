package main

import (
	"fmt"
	"reflect"
)

type (
	ValueObject[T any, S any] interface {
		Value() T
		// Equals(other S) bool
		String() string
		Show() string
	}

	valueObject[T any, S ValueObject[T, S]] struct {
		value T
	}
)

func NewValueObject[T any, S ValueObject[T, S]](v T) ValueObject[T, S] {
	return valueObject[T, S]{value: v}
}

func (v valueObject[T, S]) Value() T {
	return v.value
}

// NOTE: You can use `==` operator instead of `Equals` method.
// func (v valueObject[T, S]) Equals(other S) bool {
// 	return reflect.DeepEqual(v.Value(), other.Value())
// }

func (v valueObject[T, S]) String() string {
	return fmt.Sprintf("%v", v.value)
}

func (v valueObject[T, S]) Show() string {
	return fmt.Sprintf("%v(%v)", reflect.TypeOf((*S)(nil)).Elem(), v.value)
}
