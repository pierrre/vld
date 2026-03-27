package vld

import (
	"fmt"
)

// Signed represents all signed numeric types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Positive returns a [Validator] that checks if the value is positive.
func Positive[T Signed]() Validator[T] {
	return WithStringFunc(func() string { return "Positive" }, func(v T) error {
		if v > 0 {
			return nil
		}
		return &PositiveError[T]{
			Value: v,
		}
	})
}

// PositiveError is the error type returned by [Positive].
type PositiveError[T Signed] struct {
	Value T
}

// Error implements [error].
func (e *PositiveError[T]) Error() string {
	return fmt.Sprintf("%#v is not positive", e.Value)
}

// Localization implements [LocalizableError].
func (e *PositiveError[T]) Localization() (key string, args []any) {
	return "PositiveError", []any{e.Value}
}

// Negative returns a [Validator] that checks if the value is negative.
func Negative[T Signed]() Validator[T] {
	return WithStringFunc(func() string { return "Negative" }, func(v T) error {
		if v < 0 {
			return nil
		}
		return &NegativeError[T]{
			Value: v,
		}
	})
}

// NegativeError is the error type returned by [Negative].
type NegativeError[T Signed] struct {
	Value T
}

// Error implements [error].
func (e *NegativeError[T]) Error() string {
	return fmt.Sprintf("%#v is not negative", e.Value)
}

// Localization implements [LocalizableError].
func (e *NegativeError[T]) Localization() (key string, args []any) {
	return "NegativeError", []any{e.Value}
}
