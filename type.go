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
			return &TypeRequiredError[In, Out]{
				Value: v,
			}
		}
		return vr.Validate(vOut)
	})
}

// TypeRequiredError is the error type returned by [TypeRequired] when conversion fails.
type TypeRequiredError[In, Out any] struct {
	Value In
}

// Error implements [error].
func (e *TypeRequiredError[In, Out]) Error() string {
	return fmt.Sprintf("%T cannot be converted to %T", e.Value, *new(Out))
}

// Localization implements [LocalizableError].
func (e *TypeRequiredError[In, Out]) Localization() (key string, args []any) {
	return "TypeRequiredError", []any{e.Value, *new(Out)}
}
