package vld

import (
	"fmt"
	"slices"
)

// SliceLenEqual creates a [SliceLenEqualValidator].
func SliceLenEqual[S ~[]E, E any](length int) *SliceLenEqualValidator[S, E] {
	return &SliceLenEqualValidator[S, E]{
		Length: length,
	}
}

// SliceLenEqualValidator is a [Validator] that checks if the length of the slice is equal to the specified length.
type SliceLenEqualValidator[S ~[]E, E any] struct {
	Length int
}

// Validate implements [Validator].
func (vr *SliceLenEqualValidator[S, E]) Validate(s S) error {
	return validateLenEqual(len(s), vr.Length)
}

func (vr *SliceLenEqualValidator[S, E]) String() string {
	return fmt.Sprintf("SliceLenEqual(%d)", vr.Length)
}

// SliceLenMin creates a [SliceLenMinValidator].
func SliceLenMin[S ~[]E, E any](minLen int) *SliceLenMinValidator[S, E] {
	return &SliceLenMinValidator[S, E]{
		Min: minLen,
	}
}

// SliceLenMinValidator is a [Validator] that checks if the length of the slice is greater than or equal to the minimum length.
type SliceLenMinValidator[S ~[]E, E any] struct {
	Min int
}

// Validate implements [Validator].
func (vr *SliceLenMinValidator[S, E]) Validate(s S) error {
	return validateLenMin(len(s), vr.Min)
}

func (vr *SliceLenMinValidator[S, E]) String() string {
	return fmt.Sprintf("SliceLenMin(%d)", vr.Min)
}

// SliceLenMax creates a [SliceLenMaxValidator].
func SliceLenMax[S ~[]E, E any](maxLen int) *SliceLenMaxValidator[S, E] {
	return &SliceLenMaxValidator[S, E]{
		Max: maxLen,
	}
}

// SliceLenMaxValidator is a [Validator] that checks if the length of the slice is less than or equal to the maximum length.
type SliceLenMaxValidator[S ~[]E, E any] struct {
	Max int
}

// Validate implements [Validator].
func (vr *SliceLenMaxValidator[S, E]) Validate(s S) error {
	return validateLenMax(len(s), vr.Max)
}

func (vr *SliceLenMaxValidator[S, E]) String() string {
	return fmt.Sprintf("SliceLenMax(%d)", vr.Max)
}

// SliceLenRange creates a [SliceLenRangeValidator].
func SliceLenRange[S ~[]E, E any](minLen, maxLen int) *SliceLenRangeValidator[S, E] {
	return &SliceLenRangeValidator[S, E]{
		Min: minLen,
		Max: maxLen,
	}
}

// SliceLenRangeValidator is a [Validator] that checks if the length of the slice is within the range.
type SliceLenRangeValidator[S ~[]E, E any] struct {
	Min int
	Max int
}

// Validate implements [Validator].
func (vr *SliceLenRangeValidator[S, E]) Validate(s S) error {
	return validateLenRange(len(s), vr.Min, vr.Max)
}

func (vr *SliceLenRangeValidator[S, E]) String() string {
	return fmt.Sprintf("SliceLenRange(%d, %d)", vr.Min, vr.Max)
}

// SliceEmpty creates a [SliceEmptyValidator].
func SliceEmpty[S ~[]E, E any]() *SliceEmptyValidator[S, E] {
	return &SliceEmptyValidator[S, E]{}
}

// SliceEmptyValidator is a [Validator] that checks if the slice is empty.
type SliceEmptyValidator[S ~[]E, E any] struct{}

// Validate implements [Validator].
func (vr *SliceEmptyValidator[S, E]) Validate(s S) error {
	return validateEmpty(len(s))
}

func (vr *SliceEmptyValidator[S, E]) String() string {
	return "SliceEmpty"
}

// SliceNotEmpty creates a [SliceNotEmptyValidator].
func SliceNotEmpty[S ~[]E, E any]() *SliceNotEmptyValidator[S, E] {
	return &SliceNotEmptyValidator[S, E]{}
}

// SliceNotEmptyValidator is a [Validator] that checks if the slice is not empty.
type SliceNotEmptyValidator[S ~[]E, E any] struct{}

// Validate implements [Validator].
func (vr *SliceNotEmptyValidator[S, E]) Validate(s S) error {
	return validateNotEmpty(len(s))
}

func (vr *SliceNotEmptyValidator[S, E]) String() string {
	return "SliceNotEmpty"
}

// SliceContains creates a [SliceContainsValidator].
func SliceContains[S ~[]E, E comparable](elem E) *SliceContainsValidator[S, E] {
	return &SliceContainsValidator[S, E]{
		Element: elem,
	}
}

// SliceContainsValidator is a [Validator] that checks if the slice contains the element.
type SliceContainsValidator[S ~[]E, E comparable] struct {
	Element E
}

// Validate implements [Validator].
func (vr *SliceContainsValidator[S, E]) Validate(s S) error {
	if !slices.Contains(s, vr.Element) {
		return &SliceContainsError[E]{
			Element: vr.Element,
		}
	}
	return nil
}

func (vr *SliceContainsValidator[S, E]) String() string {
	return fmt.Sprintf("SliceContains(%#v)", vr.Element)
}

// SliceContainsError is the error type returned by [SliceContainsValidator].
type SliceContainsError[E comparable] struct {
	Element E
}

func (e *SliceContainsError[E]) Error() string {
	return fmt.Sprintf("does not contain %#v", e.Element)
}

// Localization implements [LocalizableError].
func (e *SliceContainsError[E]) Localization() (key string, args []any) {
	return "SliceContainsError", []any{e.Element}
}

// SliceNotContains creates a [SliceNotContainsValidator].
func SliceNotContains[S ~[]E, E comparable](elem E) *SliceNotContainsValidator[S, E] {
	return &SliceNotContainsValidator[S, E]{
		Element: elem,
	}
}

// SliceNotContainsValidator is a [Validator] that checks if the slice does not contain the element.
type SliceNotContainsValidator[S ~[]E, E comparable] struct {
	Element E
}

// Validate implements [Validator].
func (vr *SliceNotContainsValidator[S, E]) Validate(s S) error {
	if slices.Contains(s, vr.Element) {
		return &SliceNotContainsError[E]{
			Element: vr.Element,
		}
	}
	return nil
}

func (vr *SliceNotContainsValidator[S, E]) String() string {
	return fmt.Sprintf("SliceNotContains(%#v)", vr.Element)
}

// SliceNotContainsError is the error type returned by [SliceNotContainsValidator].
type SliceNotContainsError[E comparable] struct {
	Element E
}

func (e *SliceNotContainsError[E]) Error() string {
	return fmt.Sprintf("contains %#v", e.Element)
}

// Localization implements [LocalizableError].
func (e *SliceNotContainsError[E]) Localization() (key string, args []any) {
	return "SliceNotContainsError", []any{e.Element}
}

// SliceEach creates a [SliceEachValidator].
func SliceEach[S ~[]E, E any](vr Validator[KeyValue[int, E]]) *SliceEachValidator[S, E] {
	return &SliceEachValidator[S, E]{
		Validator: vr,
	}
}

// SliceEachValidator is a [Validator] that checks each index and element of the slice.
type SliceEachValidator[S ~[]E, E any] struct {
	Validator Validator[KeyValue[int, E]]
}

// Validate implements [Validator].
func (vr *SliceEachValidator[S, E]) Validate(s S) error {
	return validateSliceEach(s, vr.Validator.Validate)
}

func (vr *SliceEachValidator[S, E]) String() string {
	return fmt.Sprintf("SliceEach(%v)", vr.Validator)
}

// SliceEachValue creates a [SliceEachValueValidator].
func SliceEachValue[S ~[]E, E any](vr Validator[E]) *SliceEachValueValidator[S, E] {
	return &SliceEachValueValidator[S, E]{
		Validator: vr,
	}
}

// SliceEachValueValidator is a [Validator] that checks each element of the slice.
type SliceEachValueValidator[S ~[]E, E any] struct {
	Validator Validator[E]
}

// Validate implements [Validator].
func (vr *SliceEachValueValidator[S, E]) Validate(s S) error {
	return validateSliceEach(s, get(KeyValue[int, E].GetValue, vr.Validator.Validate))
}

func (vr *SliceEachValueValidator[S, E]) String() string {
	return fmt.Sprintf("SliceEachValue(%v)", vr.Validator)
}

func validateSliceEach[S ~[]E, E any](s S, f func(KeyValue[int, E]) error) error {
	var errs []error
	for i, v := range s {
		err := f(KeyValue[int, E]{Key: i, Value: v})
		if err != nil {
			errs = append(errs, ErrorWrapPathElem(err, &IndexPathElem{Index: i}))
		}
	}
	return ErrorJoin(errs...)
}

// SliceUnique creates a [SliceUniqueValidator].
func SliceUnique[S ~[]E, E comparable]() *SliceUniqueValidator[S, E] {
	return &SliceUniqueValidator[S, E]{}
}

// SliceUniqueValidator is a [Validator] that checks if all elements of the slice are unique.
type SliceUniqueValidator[S ~[]E, E comparable] struct{}

// Validate implements [Validator].
func (vr *SliceUniqueValidator[S, E]) Validate(s S) error {
	seen := make(map[E]int, len(s))
	var errs []error
	for i, v := range s {
		_, ok := seen[v]
		if !ok {
			seen[v] = i
			continue
		}
		var err error = &SliceUniqueError[E]{
			Value: v,
			Index: seen[v],
		}
		err = &PathElemError{
			Err:      err,
			PathElem: &IndexPathElem{Index: i},
		}
		errs = append(errs, err)
	}
	return ErrorJoin(errs...)
}

func (vr *SliceUniqueValidator[S, E]) String() string {
	return "SliceUnique"
}

// SliceUniqueBy creates a [SliceUniqueByValidator].
func SliceUniqueBy[S ~[]E, E any, K comparable](getKey func(E) K) *SliceUniqueByValidator[S, E, K] {
	return &SliceUniqueByValidator[S, E, K]{
		GetKey: getKey,
	}
}

// SliceUniqueByValidator is a [Validator] that checks if all elements of the slice are unique by the key returned by the key function.
type SliceUniqueByValidator[S ~[]E, E any, K comparable] struct {
	GetKey func(E) K
}

// Validate implements [Validator].
func (vr *SliceUniqueByValidator[S, E, K]) Validate(s S) error {
	seen := make(map[K]int, len(s))
	var errs []error
	for i, v := range s {
		key := vr.GetKey(v)
		_, ok := seen[key]
		if !ok {
			seen[key] = i
			continue
		}
		var err error = &SliceUniqueError[E]{
			Value: v,
			Index: seen[key],
		}
		err = &PathElemError{
			Err:      err,
			PathElem: &IndexPathElem{Index: i},
		}
		errs = append(errs, err)
	}
	return ErrorJoin(errs...)
}

func (vr *SliceUniqueByValidator[S, E, K]) String() string {
	return "SliceUniqueBy"
}

// SliceUniqueError is the error type returned by [SliceUniqueValidator] and [SliceUniqueByValidator].
type SliceUniqueError[E any] struct {
	Value E
	Index int
}

func (e *SliceUniqueError[E]) Error() string {
	return fmt.Sprintf("duplicate %#v (index %d)", e.Value, e.Index)
}

// Localization implements [LocalizableError].
func (e *SliceUniqueError[E]) Localization() (key string, args []any) {
	return "SliceUniqueError", []any{e.Value, e.Index}
}
