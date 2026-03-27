package vld

import (
	"fmt"
)

func lenEqual[T any](length int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l != length {
			return &LenEqualError{
				Length:   l,
				Expected: length,
			}
		}
		return nil
	}
}

// LenEqualError is the error type returned by [lenEqual].
type LenEqualError struct {
	Length   int
	Expected int
}

// Error implements [error].
func (e *LenEqualError) Error() string {
	return fmt.Sprintf("length %d is not equal to %d", e.Length, e.Expected)
}

// Localization implements [LocalizableError].
func (e *LenEqualError) Localization() (key string, args []any) {
	return "LenEqualError", []any{e.Length, e.Expected}
}

func lenMin[T any](minLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l < minLen {
			return &LenMinError{
				Length: l,
				Min:    minLen,
			}
		}
		return nil
	}
}

// LenMinError is the error type returned by [lenMin].
type LenMinError struct {
	Length int
	Min    int
}

// Error implements [error].
func (e *LenMinError) Error() string {
	return fmt.Sprintf("length %d is less than %d", e.Length, e.Min)
}

// Localization implements [LocalizableError].
func (e *LenMinError) Localization() (key string, args []any) {
	return "LenMinError", []any{e.Length, e.Min}
}

func lenMax[T any](maxLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l > maxLen {
			return &LenMaxError{
				Length: l,
				Max:    maxLen,
			}
		}
		return nil
	}
}

// LenMaxError is the error type returned by [lenMax].
type LenMaxError struct {
	Length int
	Max    int
}

// Error implements [error].
func (e *LenMaxError) Error() string {
	return fmt.Sprintf("length %d is greater than %d", e.Length, e.Max)
}

// Localization implements [LocalizableError].
func (e *LenMaxError) Localization() (key string, args []any) {
	return "LenMaxError", []any{e.Length, e.Max}
}

func lenRange[T any](minLen, maxLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l < minLen || l > maxLen {
			return &LenRangeError{
				Length: l,
				Min:    minLen,
				Max:    maxLen,
			}
		}
		return nil
	}
}

// LenRangeError is the error type returned by [lenRange].
type LenRangeError struct {
	Length int
	Min    int
	Max    int
}

// Error implements [error].
func (e *LenRangeError) Error() string {
	return fmt.Sprintf("length %d is not in the range [%d, %d]", e.Length, e.Min, e.Max)
}

// Localization implements [LocalizableError].
func (e *LenRangeError) Localization() (key string, args []any) {
	return "LenRangeError", []any{e.Length, e.Min, e.Max}
}

func empty[T any](getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l != 0 {
			return &NotEmptyError{
				Length: l,
			}
		}
		return nil
	}
}

// NotEmptyError is the error type returned by [empty].
type NotEmptyError struct {
	Length int
}

// Error implements [error].
func (e *NotEmptyError) Error() string {
	return fmt.Sprintf("is not empty (%d)", e.Length)
}

// Localization implements [LocalizableError].
func (e *NotEmptyError) Localization() (key string, args []any) {
	return "NotEmptyError", []any{e.Length}
}

func notEmpty[T any](getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l == 0 {
			return &EmptyError{}
		}
		return nil
	}
}

// EmptyError is the error type returned by [notEmpty].
type EmptyError struct{}

// Error implements [error].
func (e *EmptyError) Error() string {
	return "is empty"
}

// Localization implements [LocalizableError].
func (e *EmptyError) Localization() (key string, args []any) {
	return "EmptyError", nil
}
