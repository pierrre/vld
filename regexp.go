package vld

import (
	"fmt"
	"regexp"
)

// RegexpString is a type that can be either a [*regexp.Regexp] or a string.
// If the value is a string, it is compiled with [regexp.MustCompile].
// This panics if the string is not a valid regular expression.
type RegexpString interface {
	*regexp.Regexp | string
}

func getRegexp[RS RegexpString](rs RS) *regexp.Regexp {
	r, ok := any(rs).(*regexp.Regexp)
	if !ok {
		s, _ := any(rs).(string)
		r = regexp.MustCompile(s)
	}
	return r
}

// RegexpMatch returns a [Validator] that checks if the string matches the regular expression.
func RegexpMatch[RS RegexpString](rs RS) Validator[string] {
	r := getRegexp(rs)
	return WithStringFunc(func() string { return fmt.Sprintf("RegexpMatch(%q)", r) }, func(s string) error {
		if !r.MatchString(s) {
			return &RegexpMatchError{
				Value:  s,
				Regexp: r,
			}
		}
		return nil
	})
}

// RegexpMatchError is the error type returned by [RegexpMatch].
type RegexpMatchError struct {
	Value  string
	Regexp *regexp.Regexp
}

// Error implements [error].
func (e *RegexpMatchError) Error() string {
	return fmt.Sprintf("%q does not match regexp %q", e.Value, e.Regexp)
}

// Localization implements [LocalizableError].
func (e *RegexpMatchError) Localization() (key string, args []any) {
	return "RegexpMatchError", []any{e.Value, e.Regexp}
}

// RegexpNotMatch returns a [Validator] that checks if the string does not match the regular expression.
func RegexpNotMatch[RS RegexpString](rs RS) Validator[string] {
	r := getRegexp(rs)
	return WithStringFunc(func() string { return fmt.Sprintf("RegexpNotMatch(%q)", r) }, func(s string) error {
		if r.MatchString(s) {
			return &RegexpNotMatchError{
				Value:  s,
				Regexp: r,
			}
		}
		return nil
	})
}

// RegexpNotMatchError is the error type returned by [RegexpNotMatch].
type RegexpNotMatchError struct {
	Value  string
	Regexp *regexp.Regexp
}

// Error implements [error].
func (e *RegexpNotMatchError) Error() string {
	return fmt.Sprintf("%q matches regexp %q", e.Value, e.Regexp)
}

// Localization implements [LocalizableError].
func (e *RegexpNotMatchError) Localization() (key string, args []any) {
	return "RegexpNotMatchError", []any{e.Value, e.Regexp}
}
