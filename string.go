package vld

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// StringLenEqual creates a [StringLenEqualValidator].
func StringLenEqual(length int) *StringLenEqualValidator {
	return &StringLenEqualValidator{
		Length: length,
	}
}

// StringLenEqualValidator is a [Validator] that checks if the length of the string is equal to the specified length.
type StringLenEqualValidator struct {
	Length int
}

// Validate implements [Validator].
func (vr *StringLenEqualValidator) Validate(s string) error {
	return validateLenEqual(len(s), vr.Length)
}

func (vr *StringLenEqualValidator) String() string {
	return fmt.Sprintf("StringLenEqual(%d)", vr.Length)
}

// StringLenMin creates a [StringLenMinValidator].
func StringLenMin(minLen int) *StringLenMinValidator {
	return &StringLenMinValidator{
		Min: minLen,
	}
}

// StringLenMinValidator is a [Validator] that checks if the length of the string is greater than or equal to the minimum length.
type StringLenMinValidator struct {
	Min int
}

// Validate implements [Validator].
func (vr *StringLenMinValidator) Validate(s string) error {
	return validateLenMin(len(s), vr.Min)
}

func (vr *StringLenMinValidator) String() string {
	return fmt.Sprintf("StringLenMin(%d)", vr.Min)
}

// StringLenMax creates a [StringLenMaxValidator].
func StringLenMax(maxLen int) *StringLenMaxValidator {
	return &StringLenMaxValidator{
		Max: maxLen,
	}
}

// StringLenMaxValidator is a [Validator] that checks if the length of the string is less than or equal to the maximum length.
type StringLenMaxValidator struct {
	Max int
}

// Validate implements [Validator].
func (vr *StringLenMaxValidator) Validate(s string) error {
	return validateLenMax(len(s), vr.Max)
}

func (vr *StringLenMaxValidator) String() string {
	return fmt.Sprintf("StringLenMax(%d)", vr.Max)
}

// StringLenRange creates a [StringLenRangeValidator].
func StringLenRange(minLen, maxLen int) *StringLenRangeValidator {
	return &StringLenRangeValidator{
		Min: minLen,
		Max: maxLen,
	}
}

// StringLenRangeValidator is a [Validator] that checks if the length of the string is within the range.
type StringLenRangeValidator struct {
	Min int
	Max int
}

// Validate implements [Validator].
func (vr *StringLenRangeValidator) Validate(s string) error {
	return validateLenRange(len(s), vr.Min, vr.Max)
}

func (vr *StringLenRangeValidator) String() string {
	return fmt.Sprintf("StringLenRange(%d, %d)", vr.Min, vr.Max)
}

// StringEmpty creates a [StringEmptyValidator].
func StringEmpty() *StringEmptyValidator {
	return &StringEmptyValidator{}
}

// StringEmptyValidator is a [Validator] that checks if the string is empty.
type StringEmptyValidator struct{}

// Validate implements [Validator].
func (vr *StringEmptyValidator) Validate(s string) error {
	return validateEmpty(len(s))
}

func (vr *StringEmptyValidator) String() string {
	return "StringEmpty"
}

// StringNotEmpty creates a [StringNotEmptyValidator].
func StringNotEmpty() *StringNotEmptyValidator {
	return &StringNotEmptyValidator{}
}

// StringNotEmptyValidator is a [Validator] that checks if the string is not empty.
type StringNotEmptyValidator struct{}

// Validate implements [Validator].
func (vr *StringNotEmptyValidator) Validate(s string) error {
	return validateNotEmpty(len(s))
}

func (vr *StringNotEmptyValidator) String() string {
	return "StringNotEmpty"
}

// StringRunesEqual creates a [StringRunesEqualValidator].
func StringRunesEqual(length int) *StringRunesEqualValidator {
	return &StringRunesEqualValidator{
		Length: length,
	}
}

// StringRunesEqualValidator is a [Validator] that checks if the number of runes in the string is equal to the specified length.
type StringRunesEqualValidator struct {
	Length int
}

// Validate implements [Validator].
func (vr *StringRunesEqualValidator) Validate(s string) error {
	return validateLenEqual(utf8.RuneCountInString(s), vr.Length)
}

func (vr *StringRunesEqualValidator) String() string {
	return fmt.Sprintf("StringRunesEqual(%d)", vr.Length)
}

// StringRunesMin creates a [StringRunesMinValidator].
func StringRunesMin(minLength int) *StringRunesMinValidator {
	return &StringRunesMinValidator{
		Min: minLength,
	}
}

// StringRunesMinValidator is a [Validator] that checks if the number of runes in the string is greater than or equal to the minimum number of runes.
type StringRunesMinValidator struct {
	Min int
}

// Validate implements [Validator].
func (vr *StringRunesMinValidator) Validate(s string) error {
	return validateLenMin(utf8.RuneCountInString(s), vr.Min)
}

func (vr *StringRunesMinValidator) String() string {
	return fmt.Sprintf("StringRunesMin(%d)", vr.Min)
}

// StringRunesMax creates a [StringRunesMaxValidator].
func StringRunesMax(maxLength int) *StringRunesMaxValidator {
	return &StringRunesMaxValidator{
		Max: maxLength,
	}
}

// StringRunesMaxValidator is a [Validator] that checks if the number of runes in the string is less than or equal to the maximum number of runes.
type StringRunesMaxValidator struct {
	Max int
}

// Validate implements [Validator].
func (vr *StringRunesMaxValidator) Validate(s string) error {
	return validateLenMax(utf8.RuneCountInString(s), vr.Max)
}

func (vr *StringRunesMaxValidator) String() string {
	return fmt.Sprintf("StringRunesMax(%d)", vr.Max)
}

// StringRunesRange creates a [StringRunesRangeValidator].
func StringRunesRange(minLength, maxLength int) *StringRunesRangeValidator {
	return &StringRunesRangeValidator{
		Min: minLength,
		Max: maxLength,
	}
}

// StringRunesRangeValidator is a [Validator] that checks if the number of runes in the string is within the range.
type StringRunesRangeValidator struct {
	Min int
	Max int
}

// Validate implements [Validator].
func (vr *StringRunesRangeValidator) Validate(s string) error {
	return validateLenRange(utf8.RuneCountInString(s), vr.Min, vr.Max)
}

func (vr *StringRunesRangeValidator) String() string {
	return fmt.Sprintf("StringRunesRange(%d, %d)", vr.Min, vr.Max)
}

// StringContains creates a [StringContainsValidator].
func StringContains(substr string) *StringContainsValidator {
	return &StringContainsValidator{
		Sub: substr,
	}
}

// StringContainsValidator is a [Validator] that checks if the string contains the substring.
type StringContainsValidator struct {
	Sub string
}

// Validate implements [Validator].
func (vr *StringContainsValidator) Validate(s string) error {
	if !strings.Contains(s, vr.Sub) {
		return &StringContainsError{
			Value: s,
			Sub:   vr.Sub,
		}
	}
	return nil
}

func (vr *StringContainsValidator) String() string {
	return fmt.Sprintf("StringContains(%q)", vr.Sub)
}

// StringContainsError is the error type returned by [StringContainsValidator].
type StringContainsError struct {
	Value string
	Sub   string
}

func (e *StringContainsError) Error() string {
	return fmt.Sprintf("%q does not contain %q", e.Value, e.Sub)
}

// Localization implements [LocalizableError].
func (e *StringContainsError) Localization() (key string, args []any) {
	return "StringContainsError", []any{e.Value, e.Sub}
}

// StringNotContains creates a [StringNotContainsValidator].
func StringNotContains(substr string) *StringNotContainsValidator {
	return &StringNotContainsValidator{
		Sub: substr,
	}
}

// StringNotContainsValidator is a [Validator] that checks if the string does not contain the substring.
type StringNotContainsValidator struct {
	Sub string
}

// Validate implements [Validator].
func (vr *StringNotContainsValidator) Validate(s string) error {
	if strings.Contains(s, vr.Sub) {
		return &StringNotContainsError{
			Value:    s,
			Expected: vr.Sub,
		}
	}
	return nil
}

func (vr *StringNotContainsValidator) String() string {
	return fmt.Sprintf("StringNotContains(%q)", vr.Sub)
}

// StringNotContainsError is the error type returned by [StringNotContainsValidator].
type StringNotContainsError struct {
	Value    string
	Expected string
}

func (e *StringNotContainsError) Error() string {
	return fmt.Sprintf("%q contains %q", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *StringNotContainsError) Localization() (key string, args []any) {
	return "StringNotContainsError", []any{e.Value, e.Expected}
}

// StringHasPrefix creates a [StringHasPrefixValidator].
func StringHasPrefix(prefix string) *StringHasPrefixValidator {
	return &StringHasPrefixValidator{
		Prefix: prefix,
	}
}

// StringHasPrefixValidator is a [Validator] that checks if the string begins with the prefix.
type StringHasPrefixValidator struct {
	Prefix string
}

// Validate implements [Validator].
func (vr *StringHasPrefixValidator) Validate(s string) error {
	if !strings.HasPrefix(s, vr.Prefix) {
		return &StringHasPrefixError{
			Value:  s,
			Prefix: vr.Prefix,
		}
	}
	return nil
}

func (vr *StringHasPrefixValidator) String() string {
	return fmt.Sprintf("StringHasPrefix(%q)", vr.Prefix)
}

// StringHasPrefixError is the error type returned by [StringHasPrefixValidator].
type StringHasPrefixError struct {
	Value  string
	Prefix string
}

func (e *StringHasPrefixError) Error() string {
	return fmt.Sprintf("%q does not begin with %q", e.Value, e.Prefix)
}

// Localization implements [LocalizableError].
func (e *StringHasPrefixError) Localization() (key string, args []any) {
	return "StringHasPrefixError", []any{e.Value, e.Prefix}
}

// StringNotHasPrefix creates a [StringNotHasPrefixValidator].
func StringNotHasPrefix(prefix string) *StringNotHasPrefixValidator {
	return &StringNotHasPrefixValidator{
		Prefix: prefix,
	}
}

// StringNotHasPrefixValidator is a [Validator] that checks if the string does not begin with the prefix.
type StringNotHasPrefixValidator struct {
	Prefix string
}

// Validate implements [Validator].
func (vr *StringNotHasPrefixValidator) Validate(s string) error {
	if strings.HasPrefix(s, vr.Prefix) {
		return &StringNotHasPrefixError{
			Value:  s,
			Prefix: vr.Prefix,
		}
	}
	return nil
}

func (vr *StringNotHasPrefixValidator) String() string {
	return fmt.Sprintf("StringNotHasPrefix(%q)", vr.Prefix)
}

// StringNotHasPrefixError is the error type returned by [StringNotHasPrefixValidator].
type StringNotHasPrefixError struct {
	Value  string
	Prefix string
}

func (e *StringNotHasPrefixError) Error() string {
	return fmt.Sprintf("%q begins with %q", e.Value, e.Prefix)
}

// Localization implements [LocalizableError].
func (e *StringNotHasPrefixError) Localization() (key string, args []any) {
	return "StringNotHasPrefixError", []any{e.Value, e.Prefix}
}

// StringHasSuffix creates a [StringHasSuffixValidator].
func StringHasSuffix(suffix string) *StringHasSuffixValidator {
	return &StringHasSuffixValidator{
		Suffix: suffix,
	}
}

// StringHasSuffixValidator is a [Validator] that checks if the string ends with the suffix.
type StringHasSuffixValidator struct {
	Suffix string
}

// Validate implements [Validator].
func (vr *StringHasSuffixValidator) Validate(s string) error {
	if !strings.HasSuffix(s, vr.Suffix) {
		return &StringHasSuffixError{
			Value:  s,
			Suffix: vr.Suffix,
		}
	}
	return nil
}

func (vr *StringHasSuffixValidator) String() string {
	return fmt.Sprintf("StringHasSuffix(%q)", vr.Suffix)
}

// StringHasSuffixError is the error type returned by [StringHasSuffixValidator].
type StringHasSuffixError struct {
	Value  string
	Suffix string
}

func (e *StringHasSuffixError) Error() string {
	return fmt.Sprintf("%q does not end with %q", e.Value, e.Suffix)
}

// Localization implements [LocalizableError].
func (e *StringHasSuffixError) Localization() (key string, args []any) {
	return "StringHasSuffixError", []any{e.Value, e.Suffix}
}

// StringNotHasSuffix creates a [StringNotHasSuffixValidator].
func StringNotHasSuffix(suffix string) *StringNotHasSuffixValidator {
	return &StringNotHasSuffixValidator{
		Suffix: suffix,
	}
}

// StringNotHasSuffixValidator is a [Validator] that checks if the string does not end with the suffix.
type StringNotHasSuffixValidator struct {
	Suffix string
}

// Validate implements [Validator].
func (vr *StringNotHasSuffixValidator) Validate(s string) error {
	if strings.HasSuffix(s, vr.Suffix) {
		return &StringNotHasSuffixError{
			Value:  s,
			Suffix: vr.Suffix,
		}
	}
	return nil
}

func (vr *StringNotHasSuffixValidator) String() string {
	return fmt.Sprintf("StringNotHasSuffix(%q)", vr.Suffix)
}

// StringNotHasSuffixError is the error type returned by [StringNotHasSuffixValidator].
type StringNotHasSuffixError struct {
	Value  string
	Suffix string
}

func (e *StringNotHasSuffixError) Error() string {
	return fmt.Sprintf("%q ends with %q", e.Value, e.Suffix)
}

// Localization implements [LocalizableError].
func (e *StringNotHasSuffixError) Localization() (key string, args []any) {
	return "StringNotHasSuffixError", []any{e.Value, e.Suffix}
}
