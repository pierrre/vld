package vld

import (
	"errors"
	"fmt"
)

type errorsUnwrapper interface {
	Unwrap() []error
}

// ErrorWrap applies the wrapFunc to the error.
// If the error is nil, it returns nil.
// If the error implements Unwrap() []error, it applies wrapFunc to each error in the slice.
// It then returns a joined error.
// Otherwise, it applies wrapFunc to the error and returns the result.
func ErrorWrap(err error, wrapFunc func(error) error) error {
	if err == nil {
		return nil
	}
	unwrapper, ok := err.(errorsUnwrapper)
	if !ok {
		return wrapFunc(err)
	}
	errs := unwrapper.Unwrap()
	resErrs := make([]error, len(errs))
	for i, e := range errs {
		if e != nil {
			resErrs[i] = wrapFunc(e)
		}
	}
	return errors.Join(resErrs...)
}

// ErrorWrapMessage wraps the error with the message.
// See [ErrorWrap] for details.
func ErrorWrapMessage(err error, msg string) error {
	if err == nil {
		return nil
	}
	return ErrorWrap(err, func(err error) error {
		return fmt.Errorf("%s: %w", msg, err)
	})
}

// ErrorWrapMessagef wraps the error with the formatted message.
// See [ErrorWrap] for details.
func ErrorWrapMessagef(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, args...)
	return ErrorWrapMessage(err, msg)
}

// ErrorJoin takes a list of errors and returns a single error that joins them.
// If an error is nil, it is ignored.
// If an error implements Unwrap() []error, it is unwrapped.
// Each error in the slice is included in the result.
func ErrorJoin(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}
	n := 0
	for _, err := range errs {
		if err != nil {
			unwrapper, ok := err.(errorsUnwrapper)
			if ok {
				n += len(unwrapper.Unwrap())
			} else {
				n++
			}
		}
	}
	if n == 0 {
		return nil
	}
	resErrs := make([]error, 0, n)
	for _, err := range errs {
		if err != nil {
			unwrapper, ok := err.(errorsUnwrapper)
			if ok {
				resErrs = append(resErrs, unwrapper.Unwrap()...)
			} else {
				resErrs = append(resErrs, err)
			}
		}
	}
	return errors.Join(resErrs...)
}

// GetErrors returns a slice of errors from an error.
// If the error is nil, it returns nil.
// If the error implements Unwrap() []error, it returns the slice of errors.
// Otherwise it returns a slice containing the error.
func GetErrors(err error) []error {
	if err == nil {
		return nil
	}
	unwrapper, ok := err.(errorsUnwrapper)
	if !ok {
		return []error{err}
	}
	return unwrapper.Unwrap()
}
