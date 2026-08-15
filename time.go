package vld

import (
	"fmt"
	"strings"
	"time"
)

// Time creates a [TimeValidator].
func Time(layout string, loc *time.Location, vr Validator[time.Time]) *TimeValidator {
	return &TimeValidator{
		Layout:    layout,
		Location:  loc,
		Validator: vr,
	}
}

// TimeValidator is a [Validator] that parses the string to a [time.Time] with the layout, and validates the parsed time with the optional validator.
type TimeValidator struct {
	Layout    string
	Location  *time.Location
	Validator Validator[time.Time]
}

// Validate implements [Validator].
func (vr *TimeValidator) Validate(s string) error {
	var t time.Time
	var err error
	if vr.Location != nil {
		t, err = time.ParseInLocation(vr.Layout, s, vr.Location)
	} else {
		t, err = time.Parse(vr.Layout, s)
	}
	if err != nil {
		return &TimeError{
			Value:    s,
			Layout:   vr.Layout,
			Location: vr.Location,
			Err:      err,
		}
	}
	if vr.Validator == nil {
		return nil
	}
	return vr.Validator.Validate(t) //nolint:wrapcheck // Not needed.
}

func (vr *TimeValidator) String() string {
	var loc any
	if vr.Location != nil {
		loc = vr.Location
	}
	return fmt.Sprintf("Time(%q, %v, %v)", vr.Layout, loc, vr.Validator)
}

// Localize implements [Localizer].
func (vr *TimeValidator) Localize(locales ...string) string {
	args := []any{vr.Layout}
	key := "TimeValidator"
	if vr.Location != nil {
		key = "TimeInLocationValidator"
		args = append(args, vr.Location.String())
	}
	sb := new(strings.Builder)
	sb.WriteString(Localize(key, args, locales...))
	if vr.Validator != nil {
		sb.WriteString("\n")
		writeStringIndent(sb, LocalizeValidator(vr.Validator, locales...))
	}
	return sb.String()
}

// TimeError is the error type returned by [TimeValidator] when the string cannot be parsed to a time.
type TimeError struct {
	Value    string
	Layout   string
	Location *time.Location
	Err      error
}

func (e *TimeError) Error() string {
	if e.Location != nil {
		return fmt.Sprintf("%q is not a valid time with layout %q in location %s: %v", e.Value, e.Layout, e.Location, e.Err)
	}
	return fmt.Sprintf("%q is not a valid time with layout %q: %v", e.Value, e.Layout, e.Err)
}

func (e *TimeError) Unwrap() error {
	return e.Err
}

// Localization implements [LocalizableError].
func (e *TimeError) Localization() (key string, args []any) {
	if e.Location != nil {
		return "TimeInLocationError", []any{e.Value, e.Layout, e.Location.String(), e.Err}
	}
	return "TimeError", []any{e.Value, e.Layout, e.Err}
}
