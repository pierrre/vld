package vld

import (
	"fmt"
)

// Equal creates a [EqualValidator].
func Equal[T comparable](v T) *EqualValidator[T] {
	return &EqualValidator[T]{
		Expected: v,
	}
}

// EqualValidator is a [Validator] that checks if the value is equal to a specific value.
type EqualValidator[T comparable] struct {
	Expected T
}

// Validate implements [Validator].
func (vr *EqualValidator[T]) Validate(v T) error {
	return validateEqualFunc(v, vr.Expected, func(a, b T) bool { return a == b })
}

func (vr *EqualValidator[T]) String() string {
	return fmt.Sprintf("Equal(%#v)", vr.Expected)
}

// EqualFunc creates a [EqualFuncValidator].
func EqualFunc[T any](v T, eqFunc func(a, b T) bool) *EqualFuncValidator[T] {
	return &EqualFuncValidator[T]{
		Expected: v,
		Func:     eqFunc,
	}
}

// EqualFuncValidator is a [Validator] that checks if the value is equal to a specific value using a custom equality function.
type EqualFuncValidator[T any] struct {
	Expected T
	Func     func(a, b T) bool
}

// Validate implements [Validator].
func (vr *EqualFuncValidator[T]) Validate(v T) error {
	return validateEqualFunc(v, vr.Expected, vr.Func)
}

func validateEqualFunc[T any](v T, expected T, eqFunc func(a, b T) bool) error {
	if !eqFunc(v, expected) {
		return &EqualError[T]{
			Value:    v,
			Expected: expected,
		}
	}
	return nil
}

func (vr *EqualFuncValidator[T]) String() string {
	return fmt.Sprintf("EqualFunc(%#v)", vr.Expected)
}

// EqualCmpFunc creates a [EqualCmpFuncValidator].
func EqualCmpFunc[T any](v T, cmpFunc func(a, b T) int) *EqualCmpFuncValidator[T] {
	return &EqualCmpFuncValidator[T]{
		Expected: v,
		Func:     cmpFunc,
	}
}

// EqualCmpFuncValidator is a [Validator] that checks if the value is equal to a specific value using a custom comparison function.
type EqualCmpFuncValidator[T any] struct {
	Expected T
	Func     func(a, b T) int
}

// Validate implements [Validator].
func (vr *EqualCmpFuncValidator[T]) Validate(v T) error {
	return validateEqualCmpFunc(v, vr.Expected, vr.Func)
}

func validateEqualCmpFunc[T any](v T, expected T, cmpFunc func(a, b T) int) error {
	return validateEqualFunc(v, expected, func(a, b T) bool { return cmpFunc(a, b) == 0 })
}

func (vr *EqualCmpFuncValidator[T]) String() string {
	return fmt.Sprintf("EqualCmpFunc(%#v)", vr.Expected)
}

// EqualError is the error type returned by validators that check for equality.
type EqualError[T any] struct {
	Value    T
	Expected T
}

func (e *EqualError[T]) Error() string {
	return fmt.Sprintf("%#v is not equal to %#v", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *EqualError[T]) Localization() (key string, args []any) {
	return "EqualError", []any{e.Value, e.Expected}
}

// NotEqual creates a [NotEqualValidator].
func NotEqual[T comparable](v T) *NotEqualValidator[T] {
	return &NotEqualValidator[T]{
		Expected: v,
	}
}

// NotEqualValidator is a [Validator] that checks if the value is not equal to a specific value.
type NotEqualValidator[T comparable] struct {
	Expected T
}

// Validate implements [Validator].
func (vr *NotEqualValidator[T]) Validate(v T) error {
	return validateNotEqualFunc(v, vr.Expected, func(a, b T) bool { return a == b })
}

func (vr *NotEqualValidator[T]) String() string {
	return fmt.Sprintf("NotEqual(%#v)", vr.Expected)
}

// NotEqualFunc creates a [NotEqualFuncValidator].
func NotEqualFunc[T any](v T, eqFunc func(a, b T) bool) *NotEqualFuncValidator[T] {
	return &NotEqualFuncValidator[T]{
		Expected: v,
		Func:     eqFunc,
	}
}

// NotEqualFuncValidator is a [Validator] that checks if the value is not equal to a specific value using a custom equality function.
type NotEqualFuncValidator[T any] struct {
	Expected T
	Func     func(a, b T) bool
}

// Validate implements [Validator].
func (vr *NotEqualFuncValidator[T]) Validate(v T) error {
	return validateNotEqualFunc(v, vr.Expected, vr.Func)
}

func validateNotEqualFunc[T any](v T, expected T, eqFunc func(a, b T) bool) error {
	if eqFunc(v, expected) {
		return &NotEqualError[T]{
			Value:    v,
			Expected: expected,
		}
	}
	return nil
}

func (vr *NotEqualFuncValidator[T]) String() string {
	return fmt.Sprintf("NotEqualFunc(%#v)", vr.Expected)
}

// NotEqualCmpFunc creates a [NotEqualCmpFuncValidator].
func NotEqualCmpFunc[T any](v T, cmpFunc func(a, b T) int) *NotEqualCmpFuncValidator[T] {
	return &NotEqualCmpFuncValidator[T]{
		Expected: v,
		Func:     cmpFunc,
	}
}

// NotEqualCmpFuncValidator is a [Validator] that checks if the value is not equal to a specific value using a custom comparison function.
type NotEqualCmpFuncValidator[T any] struct {
	Expected T
	Func     func(a, b T) int
}

// Validate implements [Validator].
func (vr *NotEqualCmpFuncValidator[T]) Validate(v T) error {
	return validateNotEqualCmpFunc(v, vr.Expected, vr.Func)
}

func validateNotEqualCmpFunc[T any](v T, expected T, cmpFunc func(a, b T) int) error {
	return validateNotEqualFunc(v, expected, func(a, b T) bool { return cmpFunc(a, b) == 0 })
}

func (vr *NotEqualCmpFuncValidator[T]) String() string {
	return fmt.Sprintf("NotEqualCmpFunc(%#v)", vr.Expected)
}

// NotEqualError is the error type returned by validators that check for inequality.
type NotEqualError[T any] struct {
	Value    T
	Expected T
}

func (e *NotEqualError[T]) Error() string {
	return fmt.Sprintf("%#v is equal to %#v", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *NotEqualError[T]) Localization() (key string, args []any) {
	return "NotEqualError", []any{e.Value, e.Expected}
}
