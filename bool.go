package vld

import (
	"fmt"
)

// And creates a [AndValidator].
func And[T any](vrs ...Validator[T]) *AndValidator[T] {
	return &AndValidator[T]{
		Validators: vrs,
	}
}

// AndValidator is a [Validator] that validates the value with all validators and returns the first error.
type AndValidator[T any] struct {
	Validators []Validator[T]
}

// Validate implements [Validator].
func (vr *AndValidator[T]) Validate(v T) error {
	for _, validator := range vr.Validators {
		err := validator.Validate(v)
		if err != nil {
			return err //nolint:wrapcheck // Not needed.
		}
	}
	return nil
}

func (vr *AndValidator[T]) String() string {
	return getMultiValidatorString("And", vr.Validators...)
}

// Or creates a [OrValidator].
func Or[T any](vrs ...Validator[T]) *OrValidator[T] {
	return &OrValidator[T]{
		Validators: vrs,
	}
}

// OrValidator is a [Validator] that validates the value with all validators and returns nil if any returns nil, or joins all errors otherwise.
type OrValidator[T any] struct {
	Validators []Validator[T]
}

// Validate implements [Validator].
func (vr *OrValidator[T]) Validate(v T) error {
	var errs []error
	for _, validator := range vr.Validators {
		err := validator.Validate(v)
		if err == nil {
			return nil
		}
		errs = append(errs, err)
	}
	return ErrorJoin(errs...)
}

func (vr *OrValidator[T]) String() string {
	return getMultiValidatorString("Or", vr.Validators...)
}

// All creates an [AllValidator].
func All[T any](vrs ...Validator[T]) *AllValidator[T] {
	return &AllValidator[T]{
		Validators: vrs,
	}
}

// AllValidator is a [Validator] that validates the value with all validators and joins all errors.
type AllValidator[T any] struct {
	Validators []Validator[T]
}

// Validate implements [Validator].
func (vr *AllValidator[T]) Validate(v T) error {
	var errs []error
	for _, validator := range vr.Validators {
		err := validator.Validate(v)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return ErrorJoin(errs...)
}

func (vr *AllValidator[T]) String() string {
	return getMultiValidatorString("All", vr.Validators...)
}

// Not creates a [NotValidator].
func Not[T any](msg string, vr Validator[T]) *NotValidator[T] {
	return &NotValidator[T]{
		Message:   msg,
		Validator: vr,
	}
}

// NotValidator is a [Validator] that validates the value and returns nil if the underlying validator returns an error.
type NotValidator[T any] struct {
	Message   string
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *NotValidator[T]) Validate(v T) error {
	err := vr.Validator.Validate(v)
	if err == nil {
		return &NotError{Message: vr.Message}
	}
	return nil
}

func (vr *NotValidator[T]) String() string {
	return fmt.Sprintf("Not(%q, %v)", vr.Message, vr.Validator)
}

// NotError is the error type returned by [NotValidator].
type NotError struct {
	Message string
}

func (e *NotError) Error() string {
	return e.Message
}
