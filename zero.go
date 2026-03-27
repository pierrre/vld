package vld

import (
	"fmt"
)

func isZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

// Zero returns a [Validator] that checks if the value is the zero value.
func Zero[T comparable]() Validator[T] {
	return WithStringFunc(func() string { return "Zero" }, func(v T) error {
		if !isZero(v) {
			return &ZeroError[T]{
				Value: v,
			}
		}
		return nil
	})
}

// ZeroError is the error type returned by [Zero].
type ZeroError[T comparable] struct {
	Value T
}

// Error implements [error].
func (e *ZeroError[T]) Error() string {
	return fmt.Sprintf("%#v is not zero", e.Value)
}

// Localization implements [LocalizableError].
func (e *ZeroError[T]) Localization() (key string, args []any) {
	return "ZeroError", []any{e.Value}
}

// NotZero returns a [Validator] that checks if the value is not the zero value.
func NotZero[T comparable]() Validator[T] {
	return WithStringFunc(func() string { return "NotZero" }, func(v T) error {
		if isZero(v) {
			return &NotZeroError{}
		}
		return nil
	})
}

// NotZeroError is the error type returned by [NotZero].
type NotZeroError struct{}

// Error implements [error].
func (e *NotZeroError) Error() string {
	return "is zero"
}

// Localization implements [LocalizableError].
func (e *NotZeroError) Localization() (key string, args []any) {
	return "NotZeroError", nil
}

// Optional returns a [Validator] that validates the value if it's not the zero value.
func Optional[T comparable](vr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Optional(%v)", vr) }, func(v T) error {
		if isZero(v) {
			return nil
		}
		return vr.Validate(v)
	})
}

// Required returns a [Validator] that checks if the value is not the zero value, and validates the value.
func Required[T comparable](vr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Required(%v)", vr) }, func(v T) error {
		if isZero(v) {
			return &RequiredError{}
		}
		return vr.Validate(v)
	})
}

// RequiredError is the error type returned by [Required].
type RequiredError struct{}

// Error implements [error].
func (e *RequiredError) Error() string {
	return "required"
}

// Localization implements [LocalizableError].
func (e *RequiredError) Localization() (key string, args []any) {
	return "RequiredError", nil
}
