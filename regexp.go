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

// RegexpMatch creates a [RegexpMatchValidator].
func RegexpMatch[RS RegexpString](rs RS) *RegexpMatchValidator {
	return &RegexpMatchValidator{
		Regexp: getRegexp(rs),
	}
}

// RegexpMatchValidator is a [Validator] that checks if the string matches the regular expression.
type RegexpMatchValidator struct {
	Regexp *regexp.Regexp
}

// Validate implements [Validator].
func (vr *RegexpMatchValidator) Validate(s string) error {
	if !vr.Regexp.MatchString(s) {
		return &RegexpMatchError{
			Value:  s,
			Regexp: vr.Regexp,
		}
	}
	return nil
}

func (vr *RegexpMatchValidator) String() string {
	return fmt.Sprintf("RegexpMatch(%q)", vr.Regexp)
}

// Localization implements [Localizable].
func (vr *RegexpMatchValidator) Localization() (key string, args []any) {
	return "RegexpMatch", []any{vr.Regexp}
}

// RegexpMatchError is the error type returned by [RegexpMatchValidator].
type RegexpMatchError struct {
	Value  string
	Regexp *regexp.Regexp
}

func (e *RegexpMatchError) Error() string {
	return fmt.Sprintf("%q does not match regexp %q", e.Value, e.Regexp)
}

// Localization implements [LocalizableError].
func (e *RegexpMatchError) Localization() (key string, args []any) {
	return "RegexpMatchError", []any{e.Value, e.Regexp}
}

// RegexpNotMatch creates a [RegexpNotMatchValidator].
func RegexpNotMatch[RS RegexpString](rs RS) *RegexpNotMatchValidator {
	return &RegexpNotMatchValidator{
		Regexp: getRegexp(rs),
	}
}

// RegexpNotMatchValidator is a [Validator] that checks if the string does not match the regular expression.
type RegexpNotMatchValidator struct {
	Regexp *regexp.Regexp
}

// Validate implements [Validator].
func (vr *RegexpNotMatchValidator) Validate(s string) error {
	if vr.Regexp.MatchString(s) {
		return &RegexpNotMatchError{
			Value:  s,
			Regexp: vr.Regexp,
		}
	}
	return nil
}

func (vr *RegexpNotMatchValidator) String() string {
	return fmt.Sprintf("RegexpNotMatch(%q)", vr.Regexp)
}

// Localization implements [Localizable].
func (vr *RegexpNotMatchValidator) Localization() (key string, args []any) {
	return "RegexpNotMatch", []any{vr.Regexp}
}

// RegexpNotMatchError is the error type returned by [RegexpNotMatchValidator].
type RegexpNotMatchError struct {
	Value  string
	Regexp *regexp.Regexp
}

func (e *RegexpNotMatchError) Error() string {
	return fmt.Sprintf("%q matches regexp %q", e.Value, e.Regexp)
}

// Localization implements [LocalizableError].
func (e *RegexpNotMatchError) Localization() (key string, args []any) {
	return "RegexpNotMatchError", []any{e.Value, e.Regexp}
}
