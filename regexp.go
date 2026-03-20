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
			err := fmt.Errorf("%q does not match regexp %q", s, r)
			err = ErrorWrapLocalization(err, "RegexpMatch", s, r)
			return err
		}
		return nil
	})
}

// RegexpNotMatch returns a [Validator] that checks if the string does not match the regular expression.
func RegexpNotMatch[RS RegexpString](rs RS) Validator[string] {
	r := getRegexp(rs)
	return WithStringFunc(func() string { return fmt.Sprintf("RegexpNotMatch(%q)", r) }, func(s string) error {
		if r.MatchString(s) {
			err := fmt.Errorf("%q matches regexp %q", s, r)
			err = ErrorWrapLocalization(err, "RegexpNotMatch", s, r)
			return err
		}
		return nil
	})
}
