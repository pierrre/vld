package vld

import (
	"fmt"
)

func isZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

// Zero creates a [ZeroValidator].
func Zero[T comparable]() *ZeroValidator[T] {
	return &ZeroValidator[T]{}
}

// ZeroValidator is a [Validator] that checks if the value is the zero value.
type ZeroValidator[T comparable] struct{}

// Validate implements [Validator].
func (vr *ZeroValidator[T]) Validate(v T) error {
	if !isZero(v) {
		return &ZeroError[T]{
			Value: v,
		}
	}
	return nil
}

func (vr *ZeroValidator[T]) String() string {
	return "Zero"
}

// Localization implements [Localizable].
func (vr *ZeroValidator[T]) Localization() (key string, args []any) {
	return "Zero", nil
}

// ZeroError is the error type returned by [ZeroValidator].
type ZeroError[T comparable] struct {
	Value T
}

func (e *ZeroError[T]) Error() string {
	return fmt.Sprintf("%#v is not zero", e.Value)
}

// Localization implements [LocalizableError].
func (e *ZeroError[T]) Localization() (key string, args []any) {
	return "ZeroError", []any{e.Value}
}

// NotZero creates a [NotZeroValidator].
func NotZero[T comparable]() *NotZeroValidator[T] {
	return &NotZeroValidator[T]{}
}

// NotZeroValidator is a [Validator] that checks if the value is not the zero value.
type NotZeroValidator[T comparable] struct{}

// Validate implements [Validator].
func (vr *NotZeroValidator[T]) Validate(v T) error {
	if isZero(v) {
		return &NotZeroError{}
	}
	return nil
}

func (vr *NotZeroValidator[T]) String() string {
	return "NotZero"
}

// Localization implements [Localizable].
func (vr *NotZeroValidator[T]) Localization() (key string, args []any) {
	return "NotZero", nil
}

// NotZeroError is the error type returned by [NotZeroValidator].
type NotZeroError struct{}

func (e *NotZeroError) Error() string {
	return "is zero"
}

// Localization implements [LocalizableError].
func (e *NotZeroError) Localization() (key string, args []any) {
	return "NotZeroError", nil
}

// Optional creates a [OptionalValidator].
func Optional[T comparable](vr Validator[T]) *OptionalValidator[T] {
	return &OptionalValidator[T]{
		Validator: vr,
	}
}

// OptionalValidator is a [Validator] that validates the value if it's not the zero value.
type OptionalValidator[T comparable] struct {
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *OptionalValidator[T]) Validate(v T) error {
	if isZero(v) {
		return nil
	}
	return vr.Validator.Validate(v) //nolint:wrapcheck // Not needed.
}

func (vr *OptionalValidator[T]) String() string {
	return fmt.Sprintf("Optional(%v)", vr.Validator)
}

// Required creates a [RequiredValidator].
func Required[T comparable](vr Validator[T]) *RequiredValidator[T] {
	return &RequiredValidator[T]{
		Validator: vr,
	}
}

// RequiredValidator is a [Validator] that checks if the value is not the zero value, and validates the value.
type RequiredValidator[T comparable] struct {
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *RequiredValidator[T]) Validate(v T) error {
	if isZero(v) {
		return &RequiredError{}
	}
	return vr.Validator.Validate(v) //nolint:wrapcheck // Not needed.
}

func (vr *RequiredValidator[T]) String() string {
	return fmt.Sprintf("Required(%v)", vr.Validator)
}

// RequiredError is the error type returned by [RequiredValidator].
type RequiredError struct{}

func (e *RequiredError) Error() string {
	return "required"
}

// Localization implements [LocalizableError].
func (e *RequiredError) Localization() (key string, args []any) {
	return "RequiredError", nil
}
