package vld

import (
	"fmt"
)

// Equal returns a [Validator] that checks if the value is equal to a specific value.
func Equal[T comparable](v T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Equal(%#v)", v) }, equalFunc(v, func(a, b T) bool { return a == b }))
}

// EqualFunc returns a [Validator] that checks if the value is equal to a specific value using a custom equality function.
func EqualFunc[T any](expected T, eqFunc func(a, b T) bool) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("EqualFunc(%#v)", expected) }, equalFunc(expected, eqFunc))
}

func equalFunc[T any](expected T, eqFunc func(a, b T) bool) func(T) error {
	return func(v T) error {
		if !eqFunc(v, expected) {
			return fmt.Errorf("%#v is not equal to %#v", v, expected)
		}
		return nil
	}
}

// EqualCmpFunc returns a [Validator] that checks if the value is equal to a specific value using a custom comparison function.
func EqualCmpFunc[T any](expected T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("EqualCmpFunc(%#v)", expected) }, equalCmpFunc(expected, cmpFunc))
}

func equalCmpFunc[T any](expected T, cmpFunc func(a, b T) int) func(T) error {
	return equalFunc(expected, func(a, b T) bool { return cmpFunc(a, b) == 0 })
}

// NotEqual returns a [Validator] that checks if the value is not equal to a specific value.
func NotEqual[T comparable](v T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("NotEqual(%#v)", v) }, notEqualFunc(v, func(a, b T) bool { return a == b }))
}

// NotEqualFunc returns a [Validator] that checks if the value is not equal to a specific value using a custom equality function.
func NotEqualFunc[T any](expected T, eqFunc func(a, b T) bool) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("NotEqualFunc(%#v)", expected) }, notEqualFunc(expected, eqFunc))
}

func notEqualFunc[T any](expected T, eqFunc func(a, b T) bool) func(T) error {
	return func(v T) error {
		if eqFunc(v, expected) {
			return fmt.Errorf("%#v is equal to %#v", v, expected)
		}
		return nil
	}
}

// NotEqualCmpFunc returns a [Validator] that checks if the value is not equal to a specific value using a custom comparison function.
func NotEqualCmpFunc[T any](expected T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("NotEqualCmpFunc(%#v)", expected) }, notEqualCmpFunc(expected, cmpFunc))
}

func notEqualCmpFunc[T any](expected T, cmpFunc func(a, b T) int) func(T) error {
	return notEqualFunc(expected, func(a, b T) bool { return cmpFunc(a, b) == 0 })
}
