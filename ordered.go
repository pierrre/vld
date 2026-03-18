package vld

import (
	"cmp"
	"fmt"
)

// Min returns a [Validator] that checks if the value is greater than or equal to the minimum value.
func Min[T cmp.Ordered](minValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Min(%#v)", minValue) }, minCmpFunc(minValue, cmp.Compare[T]))
}

// MinCmpFunc returns a [Validator] that checks if the value is greater than or equal to the minimum value using a custom comparison function.
func MinCmpFunc[T any](minValue T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("MinCmpFunc(%#v)", minValue) }, minCmpFunc(minValue, cmpFunc))
}

func minCmpFunc[T any](minValue T, cmpFunc func(a, b T) int) func(T) error {
	return func(v T) error {
		c := cmpFunc(v, minValue)
		if c < 0 {
			err := fmt.Errorf("%#v is less than %#v", v, minValue)
			err = ErrorWrapLocalization(err, "Min", v, minValue)
			return err
		}
		return nil
	}
}

// Max returns a [Validator] that checks if the value is less than or equal to the maximum value.
func Max[T cmp.Ordered](maxValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Max(%#v)", maxValue) }, maxCmpFunc(maxValue, cmp.Compare[T]))
}

// MaxCmpFunc returns a [Validator] that checks if the value is less than or equal to the maximum value using a custom comparison function.
func MaxCmpFunc[T any](maxValue T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("MaxCmpFunc(%#v)", maxValue) }, maxCmpFunc(maxValue, cmpFunc))
}

func maxCmpFunc[T any](maxValue T, cmpFunc func(a, b T) int) func(T) error {
	return func(v T) error {
		c := cmpFunc(v, maxValue)
		if c > 0 {
			err := fmt.Errorf("%#v is greater than %#v", v, maxValue)
			err = ErrorWrapLocalization(err, "Max", v, maxValue)
			return err
		}
		return nil
	}
}

// Range returns a [Validator] that checks if the value is within the range [minValue, maxValue] (inclusive).
func Range[T cmp.Ordered](minValue, maxValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Range(%#v, %#v)", minValue, maxValue) }, rangeCmpFunc(minValue, maxValue, cmp.Compare[T]))
}

// RangeCmpFunc returns a [Validator] that checks if the value is within the range [minValue, maxValue] (inclusive) using a custom comparison function.
func RangeCmpFunc[T any](minValue, maxValue T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("RangeCmpFunc(%#v, %#v)", minValue, maxValue) }, rangeCmpFunc(minValue, maxValue, cmpFunc))
}

func rangeCmpFunc[T any](minValue, maxValue T, cmpFunc func(a, b T) int) func(T) error {
	return func(v T) error {
		cMin := cmpFunc(v, minValue)
		cMax := cmpFunc(v, maxValue)
		if cMin < 0 || cMax > 0 {
			err := fmt.Errorf("%#v is not in the range [%#v, %#v]", v, minValue, maxValue)
			err = ErrorWrapLocalization(err, "Range", v, minValue, maxValue)
			return err
		}
		return nil
	}
}

// Less returns a [Validator] that checks if the value is less than the maximum value.
func Less[T cmp.Ordered](maxValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Less(%#v)", maxValue) }, lessCmpFunc(maxValue, cmp.Compare[T]))
}

// LessCmpFunc returns a [Validator] that checks if the value is less than the maximum value using a custom comparison function.
func LessCmpFunc[T any](maxValue T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("LessCmpFunc(%#v)", maxValue) }, lessCmpFunc(maxValue, cmpFunc))
}

func lessCmpFunc[T any](maxValue T, cmpFunc func(a, b T) int) func(T) error {
	return func(v T) error {
		c := cmpFunc(v, maxValue)
		if c >= 0 {
			err := fmt.Errorf("%#v is not less than %#v", v, maxValue)
			err = ErrorWrapLocalization(err, "Less", v, maxValue)
			return err
		}
		return nil
	}
}

// Greater returns a [Validator] that checks if the value is greater than the minimum value.
func Greater[T cmp.Ordered](minValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Greater(%#v)", minValue) }, greaterCmpFunc(minValue, cmp.Compare[T]))
}

// GreaterCmpFunc returns a [Validator] that checks if the value is greater than the minimum value using a custom comparison function.
func GreaterCmpFunc[T any](minValue T, cmpFunc func(a, b T) int) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("GreaterCmpFunc(%#v)", minValue) }, greaterCmpFunc(minValue, cmpFunc))
}

func greaterCmpFunc[T any](minValue T, cmpFunc func(a, b T) int) func(T) error {
	return func(v T) error {
		c := cmpFunc(v, minValue)
		if c <= 0 {
			err := fmt.Errorf("%#v is not greater than %#v", v, minValue)
			err = ErrorWrapLocalization(err, "Greater", v, minValue)
			return err
		}
		return nil
	}
}
