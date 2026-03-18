package vld

import (
	"fmt"
)

// TypeOptional returns a [Validator] that validates the value if it can be converted to the output type.
func TypeOptional[In, Out any](vr Validator[Out]) Validator[In] {
	return WithStringFunc(func() string { return fmt.Sprintf("TypeOptional[%T](%v)", *new(Out), vr) }, func(v In) error {
		vOut, ok := any(v).(Out)
		if !ok {
			return nil
		}
		return vr.Validate(vOut)
	})
}

// TypeRequired returns a [Validator] that checks if the value can be converted to the output type, and validates the value.
func TypeRequired[In, Out any](vr Validator[Out]) Validator[In] {
	return WithStringFunc(func() string { return fmt.Sprintf("TypeRequired[%T](%v)", *new(Out), vr) }, func(v In) error {
		vOut, ok := any(v).(Out)
		if !ok {
			err := fmt.Errorf("%T cannot be converted to %T", v, *new(Out))
			err = ErrorWrapLocalization(err, "TypeRequired", v, *new(Out))
			return err
		}
		return vr.Validate(vOut)
	})
}
