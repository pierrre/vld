package vld

import (
	"errors"
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
			return fmt.Errorf("%#v is not zero", v)
		}
		return nil
	})
}

// NotZero returns a [Validator] that checks if the value is not the zero value.
func NotZero[T comparable]() Validator[T] {
	return WithStringFunc(func() string { return "NotZero" }, func(v T) error {
		if isZero(v) {
			return errors.New("is zero")
		}
		return nil
	})
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
			return errors.New("required")
		}
		return vr.Validate(v)
	})
}
