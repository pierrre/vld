package vld

import (
	"fmt"
)

// Comparer is an interface that defines a method for comparing values of type T.
type Comparer[T any] interface {
	Compare(v T) int
}

// CmpEqual creates a [CmpEqualValidator].
func CmpEqual[T Comparer[T]](v T) *CmpEqualValidator[T] {
	return &CmpEqualValidator[T]{
		Expected: v,
	}
}

// CmpEqualValidator is a [Validator] that checks if the value is equal to a specific value using the Compare method.
type CmpEqualValidator[T Comparer[T]] struct {
	Expected T
}

// Validate implements [Validator].
func (vr *CmpEqualValidator[T]) Validate(v T) error {
	return validateEqualCmpFunc(v, vr.Expected, T.Compare)
}

func (vr *CmpEqualValidator[T]) String() string {
	return fmt.Sprintf("CmpEqual(%#v)", vr.Expected)
}

// Localization implements [Localizable].
func (vr *CmpEqualValidator[T]) Localization() (key string, args []any) {
	return "Equal", []any{vr.Expected}
}

// CmpNotEqual creates a [CmpNotEqualValidator].
func CmpNotEqual[T Comparer[T]](v T) *CmpNotEqualValidator[T] {
	return &CmpNotEqualValidator[T]{
		Expected: v,
	}
}

// CmpNotEqualValidator is a [Validator] that checks if the value is not equal to a specific value using the Compare method.
type CmpNotEqualValidator[T Comparer[T]] struct {
	Expected T
}

// Validate implements [Validator].
func (vr *CmpNotEqualValidator[T]) Validate(v T) error {
	return validateNotEqualCmpFunc(v, vr.Expected, T.Compare)
}

func (vr *CmpNotEqualValidator[T]) String() string {
	return fmt.Sprintf("CmpNotEqual(%#v)", vr.Expected)
}

// Localization implements [Localizable].
func (vr *CmpNotEqualValidator[T]) Localization() (key string, args []any) {
	return "NotEqual", []any{vr.Expected}
}

// CmpMin creates a [CmpMinValidator].
func CmpMin[T Comparer[T]](minValue T) *CmpMinValidator[T] {
	return &CmpMinValidator[T]{
		Min: minValue,
	}
}

// CmpMinValidator is a [Validator] that checks if the value is greater than or equal to the minimum value using the Compare method.
type CmpMinValidator[T Comparer[T]] struct {
	Min T
}

// Validate implements [Validator].
func (vr *CmpMinValidator[T]) Validate(v T) error {
	return validateMinCmpFunc(v, vr.Min, T.Compare)
}

func (vr *CmpMinValidator[T]) String() string {
	return fmt.Sprintf("CmpMin(%#v)", vr.Min)
}

// Localization implements [Localizable].
func (vr *CmpMinValidator[T]) Localization() (key string, args []any) {
	return "Min", []any{vr.Min}
}

// CmpMax creates a [CmpMaxValidator].
func CmpMax[T Comparer[T]](maxValue T) *CmpMaxValidator[T] {
	return &CmpMaxValidator[T]{
		Max: maxValue,
	}
}

// CmpMaxValidator is a [Validator] that checks if the value is less than or equal to the maximum value using the Compare method.
type CmpMaxValidator[T Comparer[T]] struct {
	Max T
}

// Validate implements [Validator].
func (vr *CmpMaxValidator[T]) Validate(v T) error {
	return validateMaxCmpFunc(v, vr.Max, T.Compare)
}

func (vr *CmpMaxValidator[T]) String() string {
	return fmt.Sprintf("CmpMax(%#v)", vr.Max)
}

// Localization implements [Localizable].
func (vr *CmpMaxValidator[T]) Localization() (key string, args []any) {
	return "Max", []any{vr.Max}
}

// CmpRange creates a [CmpRangeValidator].
func CmpRange[T Comparer[T]](minValue, maxValue T) *CmpRangeValidator[T] {
	return &CmpRangeValidator[T]{
		Min: minValue,
		Max: maxValue,
	}
}

// CmpRangeValidator is a [Validator] that checks if the value is within the range [Min, Max] (inclusive) using the Compare method.
type CmpRangeValidator[T Comparer[T]] struct {
	Min T
	Max T
}

// Validate implements [Validator].
func (vr *CmpRangeValidator[T]) Validate(v T) error {
	return validateRangeCmpFunc(v, vr.Min, vr.Max, T.Compare)
}

func (vr *CmpRangeValidator[T]) String() string {
	return fmt.Sprintf("CmpRange(%#v, %#v)", vr.Min, vr.Max)
}

// Localization implements [Localizable].
func (vr *CmpRangeValidator[T]) Localization() (key string, args []any) {
	return "Range", []any{vr.Min, vr.Max}
}

// CmpLess creates a [CmpLessValidator].
func CmpLess[T Comparer[T]](maxValue T) *CmpLessValidator[T] {
	return &CmpLessValidator[T]{
		Max: maxValue,
	}
}

// CmpLessValidator is a [Validator] that checks if the value is less than the maximum value using the Compare method.
type CmpLessValidator[T Comparer[T]] struct {
	Max T
}

// Validate implements [Validator].
func (vr *CmpLessValidator[T]) Validate(v T) error {
	return validateLessCmpFunc(v, vr.Max, T.Compare)
}

func (vr *CmpLessValidator[T]) String() string {
	return fmt.Sprintf("CmpLess(%#v)", vr.Max)
}

// Localization implements [Localizable].
func (vr *CmpLessValidator[T]) Localization() (key string, args []any) {
	return "Less", []any{vr.Max}
}

// CmpGreater creates a [CmpGreaterValidator].
func CmpGreater[T Comparer[T]](minValue T) *CmpGreaterValidator[T] {
	return &CmpGreaterValidator[T]{
		Min: minValue,
	}
}

// CmpGreaterValidator is a [Validator] that checks if the value is greater than the minimum value using the Compare method.
type CmpGreaterValidator[T Comparer[T]] struct {
	Min T
}

// Validate implements [Validator].
func (vr *CmpGreaterValidator[T]) Validate(v T) error {
	return validateGreaterCmpFunc(v, vr.Min, T.Compare)
}

func (vr *CmpGreaterValidator[T]) String() string {
	return fmt.Sprintf("CmpGreater(%#v)", vr.Min)
}

// Localization implements [Localizable].
func (vr *CmpGreaterValidator[T]) Localization() (key string, args []any) {
	return "Greater", []any{vr.Min}
}
