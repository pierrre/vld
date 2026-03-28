package vld

import (
	"fmt"
)

// TypeOptional creates a [TypeOptionalValidator].
func TypeOptional[In, Out any](vr Validator[Out]) *TypeOptionalValidator[In, Out] {
	return &TypeOptionalValidator[In, Out]{
		Validator: vr,
	}
}

// TypeOptionalValidator is a [Validator] that validates the value if it can be converted to the output type.
type TypeOptionalValidator[In, Out any] struct {
	Validator Validator[Out]
}

// Validate implements [Validator].
func (vr *TypeOptionalValidator[In, Out]) Validate(v In) error {
	vOut, ok := any(v).(Out)
	if !ok {
		return nil
	}
	return vr.Validator.Validate(vOut) //nolint:wrapcheck // Not needed.
}

func (vr *TypeOptionalValidator[In, Out]) String() string {
	return fmt.Sprintf("TypeOptional[%T](%v)", *new(Out), vr.Validator)
}

// TypeRequired creates a [TypeRequiredValidator].
func TypeRequired[In, Out any](vr Validator[Out]) *TypeRequiredValidator[In, Out] {
	return &TypeRequiredValidator[In, Out]{
		Validator: vr,
	}
}

// TypeRequiredValidator is a [Validator] that checks if the value can be converted to the output type, and validates the value.
type TypeRequiredValidator[In, Out any] struct {
	Validator Validator[Out]
}

// Validate implements [Validator].
func (vr *TypeRequiredValidator[In, Out]) Validate(v In) error {
	vOut, ok := any(v).(Out)
	if !ok {
		return &TypeRequiredError[In, Out]{
			Value: v,
		}
	}
	return vr.Validator.Validate(vOut) //nolint:wrapcheck // Not needed.
}

func (vr *TypeRequiredValidator[In, Out]) String() string {
	return fmt.Sprintf("TypeRequired[%T](%v)", *new(Out), vr.Validator)
}

// TypeRequiredError is the error type returned by [TypeRequiredValidator] when conversion fails.
type TypeRequiredError[In, Out any] struct {
	Value In
}

func (e *TypeRequiredError[In, Out]) Error() string {
	return fmt.Sprintf("%T cannot be converted to %T", e.Value, *new(Out))
}

// Localization implements [LocalizableError].
func (e *TypeRequiredError[In, Out]) Localization() (key string, args []any) {
	return "TypeRequiredError", []any{e.Value, *new(Out)}
}
