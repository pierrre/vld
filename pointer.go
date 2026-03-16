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
		err := vr.Validate(*v)
		if err != nil {
			return ErrorWrapPathElement(err, &PointerPathElem{})
		}
		return nil
	})
}

// PointerRequired returns a [Validator] that checks if the pointer is not nil, and validates the dereferenced value.
func PointerRequired[T any](vr Validator[T]) Validator[*T] {
	return WithStringFunc(func() string { return fmt.Sprintf("PointerRequired(%v)", vr) }, func(v *T) error {
		if v == nil {
			return errors.New("pointer is nil")
		}
		err := vr.Validate(*v)
		if err != nil {
			return ErrorWrapPathElement(err, &PointerPathElem{})
		}
		return nil
	})
}
