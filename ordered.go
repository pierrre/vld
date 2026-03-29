package vld

import (
	"cmp"
	"fmt"
)

// Min creates a [MinValidator].
func Min[T cmp.Ordered](minValue T) *MinValidator[T] {
	return &MinValidator[T]{
		Min: minValue,
	}
}

// MinValidator is a [Validator] that checks if the value is greater than or equal to the minimum value.
type MinValidator[T cmp.Ordered] struct {
	Min T
}

// Validate implements [Validator].
func (vr *MinValidator[T]) Validate(v T) error {
	return validateMinCmpFunc(v, vr.Min, cmp.Compare[T])
}

func (vr *MinValidator[T]) String() string {
	return fmt.Sprintf("Min(%#v)", vr.Min)
}

// MinCmpFunc creates a [MinCmpFuncValidator].
func MinCmpFunc[T any](minValue T, cmpFunc func(a, b T) int) *MinCmpFuncValidator[T] {
	return &MinCmpFuncValidator[T]{
		Min:  minValue,
		Func: cmpFunc,
	}
}

// MinCmpFuncValidator is a [Validator] that checks if the value is greater than or equal to the minimum value using a custom comparison function.
type MinCmpFuncValidator[T any] struct {
	Min  T
	Func func(a, b T) int
}

// Validate implements [Validator].
func (vr *MinCmpFuncValidator[T]) Validate(v T) error {
	return validateMinCmpFunc(v, vr.Min, vr.Func)
}

func (vr *MinCmpFuncValidator[T]) String() string {
	return fmt.Sprintf("MinCmpFunc(%#v, %s)", vr.Min, getFuncName(vr.Func))
}

func validateMinCmpFunc[T any](v T, minValue T, cmpFunc func(a, b T) int) error {
	if cmpFunc(v, minValue) < 0 {
		return &MinError[T]{
			Value: v,
			Min:   minValue,
		}
	}
	return nil
}

// MinError is the error type returned by validators that check for minimum value.
type MinError[T any] struct {
	Value T
	Min   T
}

func (e *MinError[T]) Error() string {
	return fmt.Sprintf("%#v is less than %#v", e.Value, e.Min)
}

// Localization implements [LocalizableError].
func (e *MinError[T]) Localization() (key string, args []any) {
	return "MinError", []any{e.Value, e.Min}
}

// Max creates a [MaxValidator].
func Max[T cmp.Ordered](maxValue T) *MaxValidator[T] {
	return &MaxValidator[T]{
		Max: maxValue,
	}
}

// MaxValidator is a [Validator] that checks if the value is less than or equal to the maximum value.
type MaxValidator[T cmp.Ordered] struct {
	Max T
}

// Validate implements [Validator].
func (vr *MaxValidator[T]) Validate(v T) error {
	return validateMaxCmpFunc(v, vr.Max, cmp.Compare[T])
}

func (vr *MaxValidator[T]) String() string {
	return fmt.Sprintf("Max(%#v)", vr.Max)
}

// MaxCmpFunc creates a [MaxCmpFuncValidator].
func MaxCmpFunc[T any](maxValue T, cmpFunc func(a, b T) int) *MaxCmpFuncValidator[T] {
	return &MaxCmpFuncValidator[T]{
		Max:  maxValue,
		Func: cmpFunc,
	}
}

// MaxCmpFuncValidator is a [Validator] that checks if the value is less than or equal to the maximum value using a custom comparison function.
type MaxCmpFuncValidator[T any] struct {
	Max  T
	Func func(a, b T) int
}

// Validate implements [Validator].
func (vr *MaxCmpFuncValidator[T]) Validate(v T) error {
	return validateMaxCmpFunc(v, vr.Max, vr.Func)
}

func (vr *MaxCmpFuncValidator[T]) String() string {
	return fmt.Sprintf("MaxCmpFunc(%#v, %s)", vr.Max, getFuncName(vr.Func))
}

func validateMaxCmpFunc[T any](v T, maxValue T, cmpFunc func(a, b T) int) error {
	if cmpFunc(v, maxValue) > 0 {
		return &MaxError[T]{
			Value: v,
			Max:   maxValue,
		}
	}
	return nil
}

// MaxError is the error type returned by validators that check for maximum value.
type MaxError[T any] struct {
	Value T
	Max   T
}

func (e *MaxError[T]) Error() string {
	return fmt.Sprintf("%#v is greater than %#v", e.Value, e.Max)
}

// Localization implements [LocalizableError].
func (e *MaxError[T]) Localization() (key string, args []any) {
	return "MaxError", []any{e.Value, e.Max}
}

// Range creates a [RangeValidator].
func Range[T cmp.Ordered](minValue, maxValue T) *RangeValidator[T] {
	return &RangeValidator[T]{
		Min: minValue,
		Max: maxValue,
	}
}

// RangeValidator is a [Validator] that checks if the value is within the range [minValue, maxValue] (inclusive).
type RangeValidator[T cmp.Ordered] struct {
	Min T
	Max T
}

// Validate implements [Validator].
func (vr *RangeValidator[T]) Validate(v T) error {
	return validateRangeCmpFunc(v, vr.Min, vr.Max, cmp.Compare[T])
}

func (vr *RangeValidator[T]) String() string {
	return fmt.Sprintf("Range(%#v, %#v)", vr.Min, vr.Max)
}

// RangeCmpFunc creates a [RangeCmpFuncValidator].
func RangeCmpFunc[T any](minValue, maxValue T, cmpFunc func(a, b T) int) *RangeCmpFuncValidator[T] {
	return &RangeCmpFuncValidator[T]{
		Min:  minValue,
		Max:  maxValue,
		Func: cmpFunc,
	}
}

// RangeCmpFuncValidator is a [Validator] that checks if the value is within the range [minValue, maxValue] (inclusive) using a custom comparison function.
type RangeCmpFuncValidator[T any] struct {
	Min  T
	Max  T
	Func func(a, b T) int
}

// Validate implements [Validator].
func (vr *RangeCmpFuncValidator[T]) Validate(v T) error {
	return validateRangeCmpFunc(v, vr.Min, vr.Max, vr.Func)
}

func (vr *RangeCmpFuncValidator[T]) String() string {
	return fmt.Sprintf("RangeCmpFunc(%#v, %#v, %s)", vr.Min, vr.Max, getFuncName(vr.Func))
}

func validateRangeCmpFunc[T any](v T, minValue, maxValue T, cmpFunc func(a, b T) int) error {
	if cmpFunc(v, minValue) < 0 || cmpFunc(v, maxValue) > 0 {
		return &RangeError[T]{
			Value: v,
			Min:   minValue,
			Max:   maxValue,
		}
	}
	return nil
}

// RangeError is the error type returned by validators that check for range.
type RangeError[T any] struct {
	Value T
	Min   T
	Max   T
}

func (e *RangeError[T]) Error() string {
	return fmt.Sprintf("%#v is not in the range [%#v, %#v]", e.Value, e.Min, e.Max)
}

// Localization implements [LocalizableError].
func (e *RangeError[T]) Localization() (key string, args []any) {
	return "RangeError", []any{e.Value, e.Min, e.Max}
}

// Less creates a [LessValidator].
func Less[T cmp.Ordered](maxValue T) *LessValidator[T] {
	return &LessValidator[T]{
		Max: maxValue,
	}
}

// LessValidator is a [Validator] that checks if the value is less than the maximum value.
type LessValidator[T cmp.Ordered] struct {
	Max T
}

// Validate implements [Validator].
func (vr *LessValidator[T]) Validate(v T) error {
	return validateLessCmpFunc(v, vr.Max, cmp.Compare[T])
}

func (vr *LessValidator[T]) String() string {
	return fmt.Sprintf("Less(%#v)", vr.Max)
}

// LessCmpFunc creates a [LessCmpFuncValidator].
func LessCmpFunc[T any](maxValue T, cmpFunc func(a, b T) int) *LessCmpFuncValidator[T] {
	return &LessCmpFuncValidator[T]{
		Max:  maxValue,
		Func: cmpFunc,
	}
}

// LessCmpFuncValidator is a [Validator] that checks if the value is less than the maximum value using a custom comparison function.
type LessCmpFuncValidator[T any] struct {
	Max  T
	Func func(a, b T) int
}

// Validate implements [Validator].
func (vr *LessCmpFuncValidator[T]) Validate(v T) error {
	return validateLessCmpFunc(v, vr.Max, vr.Func)
}

func (vr *LessCmpFuncValidator[T]) String() string {
	return fmt.Sprintf("LessCmpFunc(%#v, %s)", vr.Max, getFuncName(vr.Func))
}

func validateLessCmpFunc[T any](v T, maxValue T, cmpFunc func(a, b T) int) error {
	if cmpFunc(v, maxValue) >= 0 {
		return &LessError[T]{
			Value: v,
			Max:   maxValue,
		}
	}
	return nil
}

// LessError is the error type returned by validators that check for maximum value.
type LessError[T any] struct {
	Value T
	Max   T
}

func (e *LessError[T]) Error() string {
	return fmt.Sprintf("%#v is not less than %#v", e.Value, e.Max)
}

// Localization implements [LocalizableError].
func (e *LessError[T]) Localization() (key string, args []any) {
	return "LessError", []any{e.Value, e.Max}
}

// Greater creates a [GreaterValidator].
func Greater[T cmp.Ordered](minValue T) *GreaterValidator[T] {
	return &GreaterValidator[T]{
		Min: minValue,
	}
}

// GreaterValidator is a [Validator] that checks if the value is greater than the minimum value.
type GreaterValidator[T cmp.Ordered] struct {
	Min T
}

// Validate implements [Validator].
func (vr *GreaterValidator[T]) Validate(v T) error {
	return validateGreaterCmpFunc(v, vr.Min, cmp.Compare[T])
}

func (vr *GreaterValidator[T]) String() string {
	return fmt.Sprintf("Greater(%#v)", vr.Min)
}

// GreaterCmpFunc creates a [GreaterCmpFuncValidator].
func GreaterCmpFunc[T any](minValue T, cmpFunc func(a, b T) int) *GreaterCmpFuncValidator[T] {
	return &GreaterCmpFuncValidator[T]{
		Min:  minValue,
		Func: cmpFunc,
	}
}

// GreaterCmpFuncValidator is a [Validator] that checks if the value is greater than the minimum value using a custom comparison function.
type GreaterCmpFuncValidator[T any] struct {
	Min  T
	Func func(a, b T) int
}

// Validate implements [Validator].
func (vr *GreaterCmpFuncValidator[T]) Validate(v T) error {
	return validateGreaterCmpFunc(v, vr.Min, vr.Func)
}

func (vr *GreaterCmpFuncValidator[T]) String() string {
	return fmt.Sprintf("GreaterCmpFunc(%#v, %s)", vr.Min, getFuncName(vr.Func))
}

func validateGreaterCmpFunc[T any](v T, minValue T, cmpFunc func(a, b T) int) error {
	if cmpFunc(v, minValue) <= 0 {
		return &GreaterError[T]{
			Value: v,
			Min:   minValue,
		}
	}
	return nil
}

// GreaterError is the error type returned by validators that check for minimum value.
type GreaterError[T any] struct {
	Value T
	Min   T
}

func (e *GreaterError[T]) Error() string {
	return fmt.Sprintf("%#v is not greater than %#v", e.Value, e.Min)
}

// Localization implements [LocalizableError].
func (e *GreaterError[T]) Localization() (key string, args []any) {
	return "GreaterError", []any{e.Value, e.Min}
}
