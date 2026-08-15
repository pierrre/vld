package vld

import (
	"fmt"
	"net/url"
	"strings"
)

// URL creates a [URLValidator].
// It uses [url.Parse] if parser is nil.
func URL(parser func(string) (*url.URL, error), vr Validator[url.URL]) *URLValidator {
	return &URLValidator{
		Parser:    parser,
		Validator: vr,
	}
}

// URLValidator is a [Validator] that parses the string to a [url.URL] with the parser, and validates the parsed URL with the optional validator.
type URLValidator struct {
	Parser    func(string) (*url.URL, error)
	Validator Validator[url.URL]
}

// getParser returns the parser, or [url.Parse] if it is nil.
func (vr *URLValidator) getParser() func(string) (*url.URL, error) {
	if vr.Parser == nil {
		return url.Parse
	}
	return vr.Parser
}

// Validate implements [Validator].
func (vr *URLValidator) Validate(s string) error {
	parser := vr.getParser()
	u, err := parser(s)
	if err != nil {
		return &URLError{
			Value:  s,
			Parser: parser,
			Err:    err,
		}
	}
	if vr.Validator == nil {
		return nil
	}
	return vr.Validator.Validate(*u) //nolint:wrapcheck // Not needed.
}

func (vr *URLValidator) String() string {
	return fmt.Sprintf("URL(%s, %v)", getFuncName(vr.getParser()), vr.Validator)
}

// Localize implements [Localizer].
func (vr *URLValidator) Localize(locales ...string) string {
	sb := new(strings.Builder)
	sb.WriteString(Localize("URLValidator", []any{getFuncName(vr.getParser())}, locales...))
	if vr.Validator != nil {
		sb.WriteString("\n")
		writeStringIndent(sb, LocalizeValidator(vr.Validator, locales...))
	}
	return sb.String()
}

// URLError is the error type returned by [URLValidator] when the string cannot be parsed to a URL.
type URLError struct {
	Value  string
	Parser func(string) (*url.URL, error)
	Err    error
}

func (e *URLError) Error() string {
	return fmt.Sprintf("%q is not a valid URL with parser %s: %v", e.Value, getFuncName(e.Parser), e.Err)
}

func (e *URLError) Unwrap() error {
	return e.Err
}

// Localization implements [LocalizableError].
func (e *URLError) Localization() (key string, args []any) {
	return "URLError", []any{e.Value, getFuncName(e.Parser), e.Err}
}
