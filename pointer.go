package vld

import (
	"fmt"
)

// PointerOptional creates a [PointerOptionalValidator].
func PointerOptional[T any](vr Validator[T]) *PointerOptionalValidator[T] {
	return &PointerOptionalValidator[T]{
		Validator: vr,
	}
}

// PointerOptionalValidator is a [Validator] that validates the dereferenced value if the pointer is not nil.
type PointerOptionalValidator[T any] struct {
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *PointerOptionalValidator[T]) Validate(v *T) error {
	if v == nil {
		return nil
	}
	err := vr.Validator.Validate(*v)
	if err != nil {
		return ErrorWrapPathElem(err, &PointerPathElem{})
	}
	return nil
}

func (vr *PointerOptionalValidator[T]) String() string {
	return fmt.Sprintf("PointerOptional(%v)", vr.Validator)
}

// PointerRequired creates a [PointerRequiredValidator].
func PointerRequired[T any](vr Validator[T]) *PointerRequiredValidator[T] {
	return &PointerRequiredValidator[T]{
		Validator: vr,
	}
}

// PointerRequiredValidator is a [Validator] that checks if the pointer is not nil, and validates the dereferenced value.
type PointerRequiredValidator[T any] struct {
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *PointerRequiredValidator[T]) Validate(v *T) error {
	if v == nil {
		return &PointerRequiredError{}
	}
	err := vr.Validator.Validate(*v)
	if err != nil {
		return ErrorWrapPathElem(err, &PointerPathElem{})
	}
	return nil
}

func (vr *PointerRequiredValidator[T]) String() string {
	return fmt.Sprintf("PointerRequired(%v)", vr.Validator)
}

// PointerRequiredError is the error type returned by [PointerRequiredValidator].
type PointerRequiredError struct{}

func (e *PointerRequiredError) Error() string {
	return "pointer is nil"
}

// Localization implements [LocalizableError].
func (e *PointerRequiredError) Localization() (key string, args []any) {
	return "PointerRequiredError", nil
}
