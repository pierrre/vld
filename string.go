package vld

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func getStringLen(s string) int {
	return len(s)
}

// StringLenEqual returns a [Validator] that checks if the length of the string is equal to the specified length.
func StringLenEqual(length int) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringLenEqual(%d)", length) }, lenEqual(length, getStringLen))
}

// StringLenMin returns a [Validator] that checks if the length of the string is greater than or equal to the minimum length.
func StringLenMin(minLen int) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringLenMin(%d)", minLen) }, lenMin(minLen, getStringLen))
}

// StringLenMax returns a [Validator] that checks if the length of the string is less than or equal to the maximum length.
func StringLenMax(maxLen int) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringLenMax(%d)", maxLen) }, lenMax(maxLen, getStringLen))
}

// StringLenRange returns a [Validator] that checks if the length of the string is within the range.
func StringLenRange(minLen, maxLen int) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringLenRange(%d, %d)", minLen, maxLen) }, lenRange(minLen, maxLen, getStringLen))
}

// StringEmpty returns a [Validator] that checks if the string is empty.
func StringEmpty() Validator[string] {
	return WithStringFunc(func() string { return "StringEmpty" }, empty(getStringLen))
}

// StringNotEmpty returns a [Validator] that checks if the string is not empty.
func StringNotEmpty() Validator[string] {
	return WithStringFunc(func() string { return "StringNotEmpty" }, notEmpty(getStringLen))
}

// StringRunesEqual returns a [Validator] that checks if the number of runes in the string is equal to the specified number of runes.
func StringRunesEqual(numRunes int) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringRunesEqual(%d)", numRunes) }, lenEqual(numRunes, utf8.RuneCountInString))
}

// StringRunesMin returns a [Validator] that checks if the number of runes in the string is greater than or equal to the minimum number of runes.
func StringRunesMin(minRunes int) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringRunesMin(%d)", minRunes) }, lenMin(minRunes, utf8.RuneCountInString))
}

// StringRunesMax returns a [Validator] that checks if the number of runes in the string is less than or equal to the maximum number of runes.
func StringRunesMax(maxRunes int) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringRunesMax(%d)", maxRunes) }, lenMax(maxRunes, utf8.RuneCountInString))
}

// StringRunesRange returns a [Validator] that checks if the number of runes in the string is within the range.
func StringRunesRange(minRunes, maxRunes int) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringRunesRange(%d, %d)", minRunes, maxRunes) }, lenRange(minRunes, maxRunes, utf8.RuneCountInString))
}

// StringContains returns a [Validator] that checks if the string contains the substring.
func StringContains(substr string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringContains(%q)", substr) }, func(s string) error {
		if !strings.Contains(s, substr) {
			return &StringContainsError{
				Value: s,
				Sub:   substr,
			}
		}
		return nil
	})
}

// StringContainsError is the error type returned by [StringContains].
type StringContainsError struct {
	Value string
	Sub   string
}

// Error implements [error].
func (e *StringContainsError) Error() string {
	return fmt.Sprintf("%q does not contain %q", e.Value, e.Sub)
}

// Localization implements [LocalizableError].
func (e *StringContainsError) Localization() (key string, args []any) {
	return "StringContainsError", []any{e.Value, e.Sub}
}

// StringNotContains returns a [Validator] that checks if the string does not contain the substring.
func StringNotContains(substr string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringNotContains(%q)", substr) }, func(s string) error {
		if strings.Contains(s, substr) {
			return &StringNotContainsError{
				Value:    s,
				Expected: substr,
			}
		}
		return nil
	})
}

// StringNotContainsError is the error type returned by [StringNotContains].
type StringNotContainsError struct {
	Value    string
	Expected string
}

// Error implements [error].
func (e *StringNotContainsError) Error() string {
	return fmt.Sprintf("%q contains %q", e.Value, e.Expected)
}

// Localization implements [LocalizableError].
func (e *StringNotContainsError) Localization() (key string, args []any) {
	return "StringNotContainsError", []any{e.Value, e.Expected}
}

// StringHasPrefix returns a [Validator] that checks if the string begins with the prefix.
func StringHasPrefix(prefix string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringHasPrefix(%q)", prefix) }, func(s string) error {
		if !strings.HasPrefix(s, prefix) {
			return &StringHasPrefixError{
				Value:  s,
				Prefix: prefix,
			}
		}
		return nil
	})
}

// StringHasPrefixError is the error type returned by [StringHasPrefix].
type StringHasPrefixError struct {
	Value  string
	Prefix string
}

// Error implements [error].
func (e *StringHasPrefixError) Error() string {
	return fmt.Sprintf("%q does not begin with %q", e.Value, e.Prefix)
}

// Localization implements [LocalizableError].
func (e *StringHasPrefixError) Localization() (key string, args []any) {
	return "StringHasPrefixError", []any{e.Value, e.Prefix}
}

// StringNotHasPrefix returns a [Validator] that checks if the string does not begin with the prefix.
func StringNotHasPrefix(prefix string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringNotHasPrefix(%q)", prefix) }, func(s string) error {
		if strings.HasPrefix(s, prefix) {
			return &StringNotHasPrefixError{
				Value:  s,
				Prefix: prefix,
			}
		}
		return nil
	})
}

// StringNotHasPrefixError is the error type returned by [StringNotHasPrefix].
type StringNotHasPrefixError struct {
	Value  string
	Prefix string
}

// Error implements [error].
func (e *StringNotHasPrefixError) Error() string {
	return fmt.Sprintf("%q begins with %q", e.Value, e.Prefix)
}

// Localization implements [LocalizableError].
func (e *StringNotHasPrefixError) Localization() (key string, args []any) {
	return "StringNotHasPrefixError", []any{e.Value, e.Prefix}
}

// StringHasSuffix returns a [Validator] that checks if the string ends with the suffix.
func StringHasSuffix(suffix string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringHasSuffix(%q)", suffix) }, func(s string) error {
		if !strings.HasSuffix(s, suffix) {
			return &StringHasSuffixError{
				Value:  s,
				Suffix: suffix,
			}
		}
		return nil
	})
}

// StringHasSuffixError is the error type returned by [StringHasSuffix].
type StringHasSuffixError struct {
	Value  string
	Suffix string
}

// Error implements [error].
func (e *StringHasSuffixError) Error() string {
	return fmt.Sprintf("%q does not end with %q", e.Value, e.Suffix)
}

// Localization implements [LocalizableError].
func (e *StringHasSuffixError) Localization() (key string, args []any) {
	return "StringHasSuffixError", []any{e.Value, e.Suffix}
}

// StringNotHasSuffix returns a [Validator] that checks if the string does not end with the suffix.
func StringNotHasSuffix(suffix string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringNotHasSuffix(%q)", suffix) }, func(s string) error {
		if strings.HasSuffix(s, suffix) {
			return &StringNotHasSuffixError{
				Value:  s,
				Suffix: suffix,
			}
		}
		return nil
	})
}

// StringNotHasSuffixError is the error type returned by [StringNotHasSuffix].
type StringNotHasSuffixError struct {
	Value  string
	Suffix string
}

// Error implements [error].
func (e *StringNotHasSuffixError) Error() string {
	return fmt.Sprintf("%q ends with %q", e.Value, e.Suffix)
}

// Localization implements [LocalizableError].
func (e *StringNotHasSuffixError) Localization() (key string, args []any) {
	return "StringNotHasSuffixError", []any{e.Value, e.Suffix}
}
