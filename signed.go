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
		return fmt.Errorf("%#v is not positive", v)
	})
}

// Negative returns a [Validator] that checks if the value is negative.
func Negative[T Signed]() Validator[T] {
	return WithStringFunc(func() string { return "Negative" }, func(v T) error {
		if v < 0 {
			return nil
		}
		return fmt.Errorf("%#v is not negative", v)
	})
}
