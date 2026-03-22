package vld

import (
	"errors"
	"fmt"
)

// And returns a [Validator] that returns the first error returned by any of the validators.
func And[T any](vrs ...Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return getMultiValidatorString("And", vrs...) }, func(v T) error {
		for _, vr := range vrs {
			err := vr.Validate(v)
			if err != nil {
				return err //nolint:wrapcheck // Not needed.
			}
		}
		return nil
	})
}

// Or returns a [Validator] that returns nil if any of the validators returns nil, or an error that joins all errors returned by the validators.
func Or[T any](vrs ...Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return getMultiValidatorString("Or", vrs...) }, func(v T) error {
		var errs []error
		for _, vr := range vrs {
			err := vr.Validate(v)
			if err == nil {
				return nil
			}
			errs = append(errs, err)
		}
		return ErrorJoin(errs...)
	})
}

// All returns a [Validator] that returns an error that joins all errors returned by the validators.
func All[T any](vrs ...Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return getMultiValidatorString("All", vrs...) }, func(v T) error {
		var errs []error
		for _, vr := range vrs {
			err := vr.Validate(v)
			if err != nil {
				errs = append(errs, err)
			}
		}
		return ErrorJoin(errs...)
	})
}

// Not returns a [Validator] that returns nil if the validator returns an error, or an error with the message if the validator returns nil.
func Not[T any](msg string, vr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Not(%q, %v)", msg, vr) }, func(v T) error {
		err := vr.Validate(v)
		if err == nil {
			return errors.New(msg)
		}
		return nil
	})
}

// If returns a [Validator] that validates the value if the condition is true.
func If[T any](cond func(v T) bool, vr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("If(%s, %v)", getFuncName(cond), vr) }, func(v T) error {
		if !cond(v) {
			return nil
		}
		return vr.Validate(v)
	})
}

// IfElse returns a [Validator] that validates the value with the thenVr validator if the condition is true, or with the elseVr validator if the condition is false.
func IfElse[T any](cond func(v T) bool, thenVr Validator[T], elseVr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("IfElse(%s, %v, %v)", getFuncName(cond), thenVr, elseVr) }, func(v T) error {
		if cond(v) {
			return thenVr.Validate(v)
		}
		return elseVr.Validate(v)
	})
}
