package vld

import (
	"fmt"
)

// Equal returns a [Validator] that checks if the value is equal to a specific value.
func Equal[T comparable](v T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Equal(%#v)", v) }, equalFunc(v, func(a, b T) bool { return a == b }))
}

// EqualFunc returns a [Validator] that checks if the value is equal to a specific value using a custom equality function.
func EqualFunc[T any](v T, eqFunc func(a, b T) bool) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("EqualFunc(%#v)", v) }, equalFunc(v, eqFunc))
}

func equalFunc[T any](expected T, eqFunc func(a, b T) bool) func(T) error {
	return func(v T) error {
		if !eqFunc(v, expected) {
			return &EqualError[T]{
				Value:    v,
				Expected: expected,
			}
		}
		return nil
	}
}

// EqualError is the error type returned by validators using [equalFunc].
type EqualError[T any] struct {
	Value    T
	Expected T
}

// Error implements [error].
func (e *EqualError[T]) Error() string {
	return fmt.Sprintf("%#v is not equal to %#v", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *EqualError[T]) Localization() (key string, args []any) {
	return "EqualError", []any{e.Value, e.Expected}
}

// EqualCmpFunc returns a [Validator] that checks if the value is equal to a specific value using a custom comparison function.
func EqualCmpFunc[T any](v T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("EqualCmpFunc(%#v)", v) }, equalCmpFunc(v, cmpFunc))
}

func equalCmpFunc[T any](v T, cmpFunc func(a, b T) int) func(T) error {
	return equalFunc(v, func(a, b T) bool { return cmpFunc(a, b) == 0 })
}

// NotEqual returns a [Validator] that checks if the value is not equal to a specific value.
func NotEqual[T comparable](v T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("NotEqual(%#v)", v) }, notEqualFunc(v, func(a, b T) bool { return a == b }))
}

// NotEqualFunc returns a [Validator] that checks if the value is not equal to a specific value using a custom equality function.
func NotEqualFunc[T any](v T, eqFunc func(a, b T) bool) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("NotEqualFunc(%#v)", v) }, notEqualFunc(v, eqFunc))
}

func notEqualFunc[T any](expected T, eqFunc func(a, b T) bool) func(T) error {
	return func(v T) error {
		if eqFunc(v, expected) {
			return &NotEqualError[T]{
				Value:    v,
				Expected: expected,
			}
		}
		return nil
	}
}

// NotEqualError is the error type returned by validators using [notEqualFunc].
type NotEqualError[T any] struct {
	Value    T
	Expected T
}

// Error implements [error].
func (e *NotEqualError[T]) Error() string {
	return fmt.Sprintf("%#v is equal to %#v", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *NotEqualError[T]) Localization() (key string, args []any) {
	return "NotEqualError", []any{e.Value, e.Expected}
}

// NotEqualCmpFunc returns a [Validator] that checks if the value is not equal to a specific value using a custom comparison function.
func NotEqualCmpFunc[T any](v T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("NotEqualCmpFunc(%#v)", v) }, notEqualCmpFunc(v, cmpFunc))
}

func notEqualCmpFunc[T any](v T, cmpFunc func(a, b T) int) func(T) error {
	return notEqualFunc(v, func(a, b T) bool { return cmpFunc(a, b) == 0 })
}
