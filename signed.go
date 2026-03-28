package vld

import (
	"fmt"
)

// Signed represents all signed numeric types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Positive creates a [PositiveValidator].
func Positive[T Signed]() *PositiveValidator[T] {
	return &PositiveValidator[T]{}
}

// PositiveValidator is a [Validator] that checks if the value is positive.
type PositiveValidator[T Signed] struct{}

// Validate implements [Validator].
func (vr *PositiveValidator[T]) Validate(v T) error {
	if v > 0 {
		return nil
	}
	return &PositiveError[T]{
		Value: v,
	}
}

func (vr *PositiveValidator[T]) String() string {
	return "Positive"
}

// PositiveError is the error type returned by [PositiveValidator].
type PositiveError[T Signed] struct {
	Value T
}

func (e *PositiveError[T]) Error() string {
	return fmt.Sprintf("%#v is not positive", e.Value)
}

// Localization implements [LocalizableError].
func (e *PositiveError[T]) Localization() (key string, args []any) {
	return "PositiveError", []any{e.Value}
}

// Negative creates a [NegativeValidator].
func Negative[T Signed]() *NegativeValidator[T] {
	return &NegativeValidator[T]{}
}

// NegativeValidator is a [Validator] that checks if the value is negative.
type NegativeValidator[T Signed] struct{}

// Validate implements [Validator].
func (vr *NegativeValidator[T]) Validate(v T) error {
	if v < 0 {
		return nil
	}
	return &NegativeError[T]{
		Value: v,
	}
}

func (vr *NegativeValidator[T]) String() string {
	return "Negative"
}

// NegativeError is the error type returned by [NegativeValidator].
type NegativeError[T Signed] struct {
	Value T
}

func (e *NegativeError[T]) Error() string {
	return fmt.Sprintf("%#v is not negative", e.Value)
}

// Localization implements [LocalizableError].
func (e *NegativeError[T]) Localization() (key string, args []any) {
	return "NegativeError", []any{e.Value}
}
