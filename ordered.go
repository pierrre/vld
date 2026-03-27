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
			return &MinError[T]{
				Value: v,
				Min:   minValue,
			}
		}
		return nil
	}
}

// MinError is the error type returned by validators using [minCmpFunc].
type MinError[T any] struct {
	Value T
	Min   T
}

// Error implements [error].
func (e *MinError[T]) Error() string {
	return fmt.Sprintf("%#v is less than %#v", e.Value, e.Min)
}

// Localization implements [LocalizableError].
func (e *MinError[T]) Localization() (key string, args []any) {
	return "MinError", []any{e.Value, e.Min}
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
			return &MaxError[T]{
				Value: v,
				Max:   maxValue,
			}
		}
		return nil
	}
}

// MaxError is the error type returned by validators using [maxCmpFunc].
type MaxError[T any] struct {
	Value T
	Max   T
}

// Error implements [error].
func (e *MaxError[T]) Error() string {
	return fmt.Sprintf("%#v is greater than %#v", e.Value, e.Max)
}

// Localization implements [LocalizableError].
func (e *MaxError[T]) Localization() (key string, args []any) {
	return "MaxError", []any{e.Value, e.Max}
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
			return &RangeError[T]{
				Value: v,
				Min:   minValue,
				Max:   maxValue,
			}
		}
		return nil
	}
}

// RangeError is the error type returned by validators using [rangeCmpFunc].
type RangeError[T any] struct {
	Value T
	Min   T
	Max   T
}

// Error implements [error].
func (e *RangeError[T]) Error() string {
	return fmt.Sprintf("%#v is not in the range [%#v, %#v]", e.Value, e.Min, e.Max)
}

// Localization implements [LocalizableError].
func (e *RangeError[T]) Localization() (key string, args []any) {
	return "RangeError", []any{e.Value, e.Min, e.Max}
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
			return &LessError[T]{
				Value: v,
				Max:   maxValue,
			}
		}
		return nil
	}
}

// LessError is the error type returned by validators using [lessCmpFunc].
type LessError[T any] struct {
	Value T
	Max   T
}

// Error implements [error].
func (e *LessError[T]) Error() string {
	return fmt.Sprintf("%#v is not less than %#v", e.Value, e.Max)
}

// Localization implements [LocalizableError].
func (e *LessError[T]) Localization() (key string, args []any) {
	return "LessError", []any{e.Value, e.Max}
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
			return &GreaterError[T]{
				Value: v,
				Min:   minValue,
			}
		}
		return nil
	}
}

// GreaterError is the error type returned by validators using [greaterCmpFunc].
type GreaterError[T any] struct {
	Value T
	Min   T
}

// Error implements [error].
func (e *GreaterError[T]) Error() string {
	return fmt.Sprintf("%#v is not greater than %#v", e.Value, e.Min)
}

// Localization implements [LocalizableError].
func (e *GreaterError[T]) Localization() (key string, args []any) {
	return "GreaterError", []any{e.Value, e.Min}
}
