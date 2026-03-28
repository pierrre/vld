package vld

import (
	"bytes"
	"fmt"
)

// BytesEqual creates a [BytesEqualValidator].
func BytesEqual(s []byte) *BytesEqualValidator {
	return &BytesEqualValidator{
		Expected: s,
	}
}

// BytesEqualValidator is a [Validator] that checks if the byte slice is equal to the expected byte slice.
type BytesEqualValidator struct {
	Expected []byte
}

// Validate implements [Validator].
func (vr *BytesEqualValidator) Validate(v []byte) error {
	if !bytes.Equal(v, vr.Expected) {
		return &BytesEqualError{
			Value:    v,
			Expected: vr.Expected,
		}
	}
	return nil
}

func (vr *BytesEqualValidator) String() string {
	return fmt.Sprintf("BytesEqual(%q)", vr.Expected)
}

// BytesEqualError is the error type returned by [BytesEqualValidator].
type BytesEqualError struct {
	Value    []byte
	Expected []byte
}

func (e *BytesEqualError) Error() string {
	return fmt.Sprintf("%q is not equal to %q", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *BytesEqualError) Localization() (key string, args []any) {
	return "BytesEqualError", []any{e.Value, e.Expected}
}

// BytesNotEqual creates a [BytesNotEqualValidator].
func BytesNotEqual(s []byte) *BytesNotEqualValidator {
	return &BytesNotEqualValidator{
		Expected: s,
	}
}

// BytesNotEqualValidator is a [Validator] that checks if the byte slice is not equal to the expected byte slice.
type BytesNotEqualValidator struct {
	Expected []byte
}

// Validate implements [Validator].
func (vr *BytesNotEqualValidator) Validate(v []byte) error {
	if bytes.Equal(v, vr.Expected) {
		return &BytesNotEqualError{
			Value:    v,
			Expected: vr.Expected,
		}
	}
	return nil
}

func (vr *BytesNotEqualValidator) String() string {
	return fmt.Sprintf("BytesNotEqual(%q)", vr.Expected)
}

// BytesNotEqualError is the error type returned by [BytesNotEqualValidator].
type BytesNotEqualError struct {
	Value    []byte
	Expected []byte
}

func (e *BytesNotEqualError) Error() string {
	return fmt.Sprintf("%q is equal to %q", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *BytesNotEqualError) Localization() (key string, args []any) {
	return "BytesNotEqualError", []any{e.Value, e.Expected}
}

// BytesContains creates a [BytesContainsValidator].
func BytesContains(sub []byte) *BytesContainsValidator {
	return &BytesContainsValidator{
		Sub: sub,
	}
}

// BytesContainsValidator is a [Validator] that checks if the byte slice contains the specified byte slice.
type BytesContainsValidator struct {
	Sub []byte
}

// Validate implements [Validator].
func (vr *BytesContainsValidator) Validate(v []byte) error {
	if !bytes.Contains(v, vr.Sub) {
		return &BytesContainsError{
			Value: v,
			Sub:   vr.Sub,
		}
	}
	return nil
}

func (vr *BytesContainsValidator) String() string {
	return fmt.Sprintf("BytesContains(%q)", vr.Sub)
}

// BytesContainsError is the error type returned by [BytesContainsValidator].
type BytesContainsError struct {
	Value []byte
	Sub   []byte
}

func (e *BytesContainsError) Error() string {
	return fmt.Sprintf("%q does not contain %q", e.Value, e.Sub)
}

// Localization implements [LocalizableError].
func (e *BytesContainsError) Localization() (key string, args []any) {
	return "BytesContainsError", []any{e.Value, e.Sub}
}

// BytesNotContains creates a [BytesNotContainsValidator].
func BytesNotContains(sub []byte) *BytesNotContainsValidator {
	return &BytesNotContainsValidator{
		Sub: sub,
	}
}

// BytesNotContainsValidator is a [Validator] that checks if the byte slice does not contain the specified byte slice.
type BytesNotContainsValidator struct {
	Sub []byte
}

// Validate implements [Validator].
func (vr *BytesNotContainsValidator) Validate(v []byte) error {
	if bytes.Contains(v, vr.Sub) {
		return &BytesNotContainsError{
			Value: v,
			Sub:   vr.Sub,
		}
	}
	return nil
}

func (vr *BytesNotContainsValidator) String() string {
	return fmt.Sprintf("BytesNotContains(%q)", vr.Sub)
}

// BytesNotContainsError is the error type returned by [BytesNotContainsValidator].
type BytesNotContainsError struct {
	Value []byte
	Sub   []byte
}

func (e *BytesNotContainsError) Error() string {
	return fmt.Sprintf("%q contains %q", e.Value, e.Sub)
}

// Localization implements [LocalizableError].
func (e *BytesNotContainsError) Localization() (key string, args []any) {
	return "BytesNotContainsError", []any{e.Value, e.Sub}
}

// BytesHasPrefix creates a [BytesHasPrefixValidator].
func BytesHasPrefix(prefix []byte) *BytesHasPrefixValidator {
	return &BytesHasPrefixValidator{
		Prefix: prefix,
	}
}

// BytesHasPrefixValidator is a [Validator] that checks if the byte slice has the specified prefix.
type BytesHasPrefixValidator struct {
	Prefix []byte
}

// Validate implements [Validator].
func (vr *BytesHasPrefixValidator) Validate(v []byte) error {
	if !bytes.HasPrefix(v, vr.Prefix) {
		return &BytesHasPrefixError{
			Value:  v,
			Prefix: vr.Prefix,
		}
	}
	return nil
}

func (vr *BytesHasPrefixValidator) String() string {
	return fmt.Sprintf("BytesHasPrefix(%q)", vr.Prefix)
}

// BytesHasPrefixError is the error type returned by [BytesHasPrefixValidator].
type BytesHasPrefixError struct {
	Value  []byte
	Prefix []byte
}

func (e *BytesHasPrefixError) Error() string {
	return fmt.Sprintf("%q does not have prefix %q", e.Value, e.Prefix)
}

// Localization implements [LocalizableError].
func (e *BytesHasPrefixError) Localization() (key string, args []any) {
	return "BytesHasPrefixError", []any{e.Value, e.Prefix}
}

// BytesNotHasPrefix creates a [BytesNotHasPrefixValidator].
func BytesNotHasPrefix(prefix []byte) *BytesNotHasPrefixValidator {
	return &BytesNotHasPrefixValidator{
		Prefix: prefix,
	}
}

// BytesNotHasPrefixValidator is a [Validator] that checks if the byte slice does not have the specified prefix.
type BytesNotHasPrefixValidator struct {
	Prefix []byte
}

// Validate implements [Validator].
func (vr *BytesNotHasPrefixValidator) Validate(v []byte) error {
	if bytes.HasPrefix(v, vr.Prefix) {
		return &BytesNotHasPrefixError{
			Value:  v,
			Prefix: vr.Prefix,
		}
	}
	return nil
}

func (vr *BytesNotHasPrefixValidator) String() string {
	return fmt.Sprintf("BytesNotHasPrefix(%q)", vr.Prefix)
}

// BytesNotHasPrefixError is the error type returned by [BytesNotHasPrefixValidator].
type BytesNotHasPrefixError struct {
	Value  []byte
	Prefix []byte
}

func (e *BytesNotHasPrefixError) Error() string {
	return fmt.Sprintf("%q has prefix %q", e.Value, e.Prefix)
}

// Localization implements [LocalizableError].
func (e *BytesNotHasPrefixError) Localization() (key string, args []any) {
	return "BytesNotHasPrefixError", []any{e.Value, e.Prefix}
}

// BytesHasSuffix creates a [BytesHasSuffixValidator].
func BytesHasSuffix(suffix []byte) *BytesHasSuffixValidator {
	return &BytesHasSuffixValidator{
		Suffix: suffix,
	}
}

// BytesHasSuffixValidator is a [Validator] that checks if the byte slice has the specified suffix.
type BytesHasSuffixValidator struct {
	Suffix []byte
}

// Validate implements [Validator].
func (vr *BytesHasSuffixValidator) Validate(v []byte) error {
	if !bytes.HasSuffix(v, vr.Suffix) {
		return &BytesHasSuffixError{
			Value:  v,
			Suffix: vr.Suffix,
		}
	}
	return nil
}

func (vr *BytesHasSuffixValidator) String() string {
	return fmt.Sprintf("BytesHasSuffix(%q)", vr.Suffix)
}

// BytesHasSuffixError is the error type returned by [BytesHasSuffixValidator].
type BytesHasSuffixError struct {
	Value  []byte
	Suffix []byte
}

func (e *BytesHasSuffixError) Error() string {
	return fmt.Sprintf("%q does not have suffix %q", e.Value, e.Suffix)
}

// Localization implements [LocalizableError].
func (e *BytesHasSuffixError) Localization() (key string, args []any) {
	return "BytesHasSuffixError", []any{e.Value, e.Suffix}
}

// BytesNotHasSuffix creates a [BytesNotHasSuffixValidator].
func BytesNotHasSuffix(suffix []byte) *BytesNotHasSuffixValidator {
	return &BytesNotHasSuffixValidator{
		Suffix: suffix,
	}
}

// BytesNotHasSuffixValidator is a [Validator] that checks if the byte slice does not have the specified suffix.
type BytesNotHasSuffixValidator struct {
	Suffix []byte
}

// Validate implements [Validator].
func (vr *BytesNotHasSuffixValidator) Validate(v []byte) error {
	if bytes.HasSuffix(v, vr.Suffix) {
		return &BytesNotHasSuffixError{
			Value:  v,
			Suffix: vr.Suffix,
		}
	}
	return nil
}

func (vr *BytesNotHasSuffixValidator) String() string {
	return fmt.Sprintf("BytesNotHasSuffix(%q)", vr.Suffix)
}

// BytesNotHasSuffixError is the error type returned by [BytesNotHasSuffixValidator].
type BytesNotHasSuffixError struct {
	Value  []byte
	Suffix []byte
}

func (e *BytesNotHasSuffixError) Error() string {
	return fmt.Sprintf("%q has suffix %q", e.Value, e.Suffix)
}

// Localization implements [LocalizableError].
func (e *BytesNotHasSuffixError) Localization() (key string, args []any) {
	return "BytesNotHasSuffixError", []any{e.Value, e.Suffix}
}
