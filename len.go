package vld

import (
	"fmt"
)

func validateLenEqual(length int, expected int) error {
	if length != expected {
		return &LenEqualError{
			Length:   length,
			Expected: expected,
		}
	}
	return nil
}

// LenEqualError is the error type returned by validators that check for length equality.
type LenEqualError struct {
	Length   int
	Expected int
}

func (e *LenEqualError) Error() string {
	return fmt.Sprintf("length %d is not equal to %d", e.Length, e.Expected)
}

// Localization implements [LocalizableError].
func (e *LenEqualError) Localization() (key string, args []any) {
	return "LenEqualError", []any{e.Length, e.Expected}
}

func validateLenMin(length int, minLen int) error {
	if length < minLen {
		return &LenMinError{
			Length: length,
			Min:    minLen,
		}
	}
	return nil
}

// LenMinError is the error type returned by validators that check for minimum length.
type LenMinError struct {
	Length int
	Min    int
}

func (e *LenMinError) Error() string {
	return fmt.Sprintf("length %d is less than %d", e.Length, e.Min)
}

// Localization implements [LocalizableError].
func (e *LenMinError) Localization() (key string, args []any) {
	return "LenMinError", []any{e.Length, e.Min}
}

func validateLenMax(length int, maxLen int) error {
	if length > maxLen {
		return &LenMaxError{
			Length: length,
			Max:    maxLen,
		}
	}
	return nil
}

// LenMaxError is the error type returned by validators that check for maximum length.
type LenMaxError struct {
	Length int
	Max    int
}

func (e *LenMaxError) Error() string {
	return fmt.Sprintf("length %d is greater than %d", e.Length, e.Max)
}

// Localization implements [LocalizableError].
func (e *LenMaxError) Localization() (key string, args []any) {
	return "LenMaxError", []any{e.Length, e.Max}
}

func validateLenRange(length int, minLen, maxLen int) error {
	if length < minLen || length > maxLen {
		return &LenRangeError{
			Length: length,
			Min:    minLen,
			Max:    maxLen,
		}
	}
	return nil
}

// LenRangeError is the error type returned by validators that check for length range.
type LenRangeError struct {
	Length int
	Min    int
	Max    int
}

func (e *LenRangeError) Error() string {
	return fmt.Sprintf("length %d is not in the range [%d, %d]", e.Length, e.Min, e.Max)
}

// Localization implements [LocalizableError].
func (e *LenRangeError) Localization() (key string, args []any) {
	return "LenRangeError", []any{e.Length, e.Min, e.Max}
}

func validateEmpty(length int) error {
	if length != 0 {
		return &EmptyError{
			Length: length,
		}
	}
	return nil
}

// EmptyError is the error type returned by validators that check for empty values.
type EmptyError struct {
	Length int
}

func (e *EmptyError) Error() string {
	return fmt.Sprintf("is not empty (%d)", e.Length)
}

// Localization implements [LocalizableError].
func (e *EmptyError) Localization() (key string, args []any) {
	return "EmptyError", []any{e.Length}
}

func validateNotEmpty(length int) error {
	if length == 0 {
		return &NotEmptyError{}
	}
	return nil
}

// NotEmptyError is the error type returned by validators that check for non-empty values.
type NotEmptyError struct{}

func (e *NotEmptyError) Error() string {
	return "is empty"
}

// Localization implements [LocalizableError].
func (e *NotEmptyError) Localization() (key string, args []any) {
	return "NotEmptyError", nil
}
