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
			err := fmt.Errorf("%q does not contain %q", s, substr)
			err = ErrorWrapLocalization(err, "StringContains", s, substr)
			return err
		}
		return nil
	})
}

// StringNotContains returns a [Validator] that checks if the string does not contain the substring.
func StringNotContains(substr string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringNotContains(%q)", substr) }, func(s string) error {
		if strings.Contains(s, substr) {
			err := fmt.Errorf("%q contains %q", s, substr)
			err = ErrorWrapLocalization(err, "StringNotContains", s, substr)
			return err
		}
		return nil
	})
}

// StringHasPrefix returns a [Validator] that checks if the string begins with the prefix.
func StringHasPrefix(prefix string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringHasPrefix(%q)", prefix) }, func(s string) error {
		if !strings.HasPrefix(s, prefix) {
			err := fmt.Errorf("%q does not begin with %q", s, prefix)
			err = ErrorWrapLocalization(err, "StringHasPrefix", s, prefix)
			return err
		}
		return nil
	})
}

// StringNotHasPrefix returns a [Validator] that checks if the string does not begin with the prefix.
func StringNotHasPrefix(prefix string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringNotHasPrefix(%q)", prefix) }, func(s string) error {
		if strings.HasPrefix(s, prefix) {
			err := fmt.Errorf("%q begins with %q", s, prefix)
			err = ErrorWrapLocalization(err, "StringNotHasPrefix", s, prefix)
			return err
		}
		return nil
	})
}

// StringHasSuffix returns a [Validator] that checks if the string ends with the suffix.
func StringHasSuffix(suffix string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringHasSuffix(%q)", suffix) }, func(s string) error {
		if !strings.HasSuffix(s, suffix) {
			err := fmt.Errorf("%q does not end with %q", s, suffix)
			err = ErrorWrapLocalization(err, "StringHasSuffix", s, suffix)
			return err
		}
		return nil
	})
}

// StringNotHasSuffix returns a [Validator] that checks if the string does not end with the suffix.
func StringNotHasSuffix(suffix string) Validator[string] {
	return WithStringFunc(func() string { return fmt.Sprintf("StringNotHasSuffix(%q)", suffix) }, func(s string) error {
		if strings.HasSuffix(s, suffix) {
			err := fmt.Errorf("%q ends with %q", s, suffix)
			err = ErrorWrapLocalization(err, "StringNotHasSuffix", s, suffix)
			return err
		}
		return nil
	})
}
