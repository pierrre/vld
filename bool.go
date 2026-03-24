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
