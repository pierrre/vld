package vld

import (
	"errors"
	"fmt"
)

// PointerOptional returns a [Validator] that validates the dereferenced value if the pointer is not nil.
func PointerOptional[T any](vr Validator[T]) Validator[*T] {
	return WithStringFunc(func() string { return fmt.Sprintf("PointerOptional(%v)", vr) }, func(v *T) error {
		if v == nil {
			return nil
		}
		return vr.Validate(*v)
	})
}

// PointerRequired returns a [Validator] that checks if the pointer is not nil, and validates the dereferenced value.
func PointerRequired[T any](vr Validator[T]) Validator[*T] {
	return WithStringFunc(func() string { return fmt.Sprintf("PointerRequired(%v)", vr) }, func(v *T) error {
		if v == nil {
			return errors.New("pointer is nil")
		}
		return vr.Validate(*v)
	})
}
