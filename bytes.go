package vld

import (
	"bytes"
	"fmt"
)

// BytesEqual returns a [Validator] that checks if the byte slice is equal to the specified byte slice.
func BytesEqual(s []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesEqual(%q)", s) }, func(b []byte) error {
		if !bytes.Equal(b, s) {
			return &BytesEqualError{
				Value:    b,
				Expected: s,
			}
		}
		return nil
	})
}

// BytesEqualError is the error type returned by [BytesEqual].
type BytesEqualError struct {
	Value    []byte
	Expected []byte
}

// Error implements [error].
func (e *BytesEqualError) Error() string {
	return fmt.Sprintf("%q is not equal to %q", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *BytesEqualError) Localization() (key string, args []any) {
	return "BytesEqualError", []any{e.Value, e.Expected}
}

// BytesNotEqual returns a [Validator] that checks if the byte slice is not equal to the specified byte slice.
func BytesNotEqual(s []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesNotEqual(%q)", s) }, func(b []byte) error {
		if bytes.Equal(b, s) {
			return &BytesNotEqualError{
				Value:    b,
				Expected: s,
			}
		}
		return nil
	})
}

// BytesNotEqualError is the error type returned by [BytesNotEqual].
type BytesNotEqualError struct {
	Value    []byte
	Expected []byte
}

// Error implements [error].
func (e *BytesNotEqualError) Error() string {
	return fmt.Sprintf("%q is equal to %q", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *BytesNotEqualError) Localization() (key string, args []any) {
	return "BytesNotEqualError", []any{e.Value, e.Expected}
}

// BytesContains returns a [Validator] that checks if the byte slice contains the specified byte slice.
func BytesContains(sub []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesContains(%q)", sub) }, func(b []byte) error {
		if !bytes.Contains(b, sub) {
			return &BytesContainsError{
				Value: b,
				Sub:   sub,
			}
		}
		return nil
	})
}

// BytesContainsError is the error type returned by [BytesContains].
type BytesContainsError struct {
	Value []byte
	Sub   []byte
}

// Error implements [error].
func (e *BytesContainsError) Error() string {
	return fmt.Sprintf("%q does not contain %q", e.Value, e.Sub)
}

// Localization implements [LocalizableError].
func (e *BytesContainsError) Localization() (key string, args []any) {
	return "BytesContainsError", []any{e.Value, e.Sub}
}

// BytesNotContains returns a [Validator] that checks if the byte slice does not contain the specified byte slice.
func BytesNotContains(sub []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesNotContains(%q)", sub) }, func(b []byte) error {
		if bytes.Contains(b, sub) {
			return &BytesNotContainsError{
				Value: b,
				Sub:   sub,
			}
		}
		return nil
	})
}

// BytesNotContainsError is the error type returned by [BytesNotContains].
type BytesNotContainsError struct {
	Value []byte
	Sub   []byte
}

// Error implements [error].
func (e *BytesNotContainsError) Error() string {
	return fmt.Sprintf("%q contains %q", e.Value, e.Sub)
}

// Localization implements [LocalizableError].
func (e *BytesNotContainsError) Localization() (key string, args []any) {
	return "BytesNotContainsError", []any{e.Value, e.Sub}
}

// BytesHasPrefix returns a [Validator] that checks if the byte slice has the specified prefix.
func BytesHasPrefix(prefix []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesHasPrefix(%q)", prefix) }, func(b []byte) error {
		if !bytes.HasPrefix(b, prefix) {
			return &BytesHasPrefixError{
				Value:  b,
				Prefix: prefix,
			}
		}
		return nil
	})
}

// BytesHasPrefixError is the error type returned by [BytesHasPrefix].
type BytesHasPrefixError struct {
	Value  []byte
	Prefix []byte
}

// Error implements [error].
func (e *BytesHasPrefixError) Error() string {
	return fmt.Sprintf("%q does not have prefix %q", e.Value, e.Prefix)
}

// Localization implements [LocalizableError].
func (e *BytesHasPrefixError) Localization() (key string, args []any) {
	return "BytesHasPrefixError", []any{e.Value, e.Prefix}
}

// BytesNotHasPrefix returns a [Validator] that checks if the byte slice does not have the specified prefix.
func BytesNotHasPrefix(prefix []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesNotHasPrefix(%q)", prefix) }, func(b []byte) error {
		if bytes.HasPrefix(b, prefix) {
			return &BytesNotHasPrefixError{
				Value:  b,
				Prefix: prefix,
			}
		}
		return nil
	})
}

// BytesNotHasPrefixError is the error type returned by [BytesNotHasPrefix].
type BytesNotHasPrefixError struct {
	Value  []byte
	Prefix []byte
}

// Error implements [error].
func (e *BytesNotHasPrefixError) Error() string {
	return fmt.Sprintf("%q has prefix %q", e.Value, e.Prefix)
}

// Localization implements [LocalizableError].
func (e *BytesNotHasPrefixError) Localization() (key string, args []any) {
	return "BytesNotHasPrefixError", []any{e.Value, e.Prefix}
}

// BytesHasSuffix returns a [Validator] that checks if the byte slice has the specified suffix.
func BytesHasSuffix(suffix []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesHasSuffix(%q)", suffix) }, func(b []byte) error {
		if !bytes.HasSuffix(b, suffix) {
			return &BytesHasSuffixError{
				Value:  b,
				Suffix: suffix,
			}
		}
		return nil
	})
}

// BytesHasSuffixError is the error type returned by [BytesHasSuffix].
type BytesHasSuffixError struct {
	Value  []byte
	Suffix []byte
}

// Error implements [error].
func (e *BytesHasSuffixError) Error() string {
	return fmt.Sprintf("%q does not have suffix %q", e.Value, e.Suffix)
}

// Localization implements [LocalizableError].
func (e *BytesHasSuffixError) Localization() (key string, args []any) {
	return "BytesHasSuffixError", []any{e.Value, e.Suffix}
}

// BytesNotHasSuffix returns a [Validator] that checks if the byte slice does not have the specified suffix.
func BytesNotHasSuffix(suffix []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesNotHasSuffix(%q)", suffix) }, func(b []byte) error {
		if bytes.HasSuffix(b, suffix) {
			return &BytesNotHasSuffixError{
				Value:  b,
				Suffix: suffix,
			}
		}
		return nil
	})
}

// BytesNotHasSuffixError is the error type returned by [BytesNotHasSuffix].
type BytesNotHasSuffixError struct {
	Value  []byte
	Suffix []byte
}

// Error implements [error].
func (e *BytesNotHasSuffixError) Error() string {
	return fmt.Sprintf("%q has suffix %q", e.Value, e.Suffix)
}

// Localization implements [LocalizableError].
func (e *BytesNotHasSuffixError) Localization() (key string, args []any) {
	return "BytesNotHasSuffixError", []any{e.Value, e.Suffix}
}
