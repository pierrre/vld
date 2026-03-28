package vld

import (
	"cmp"
	"fmt"
	"slices"
)

// MapLenEqual creates a [MapLenEqualValidator].
func MapLenEqual[M ~map[K]V, K comparable, V any](length int) *MapLenEqualValidator[M, K, V] {
	return &MapLenEqualValidator[M, K, V]{
		Length: length,
	}
}

// MapLenEqualValidator is a [Validator] that checks if the length of the map is equal to the specified length.
type MapLenEqualValidator[M ~map[K]V, K comparable, V any] struct {
	Length int
}

// Validate implements [Validator].
func (vr *MapLenEqualValidator[M, K, V]) Validate(m M) error {
	return validateLenEqual(len(m), vr.Length)
}

func (vr *MapLenEqualValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapLenEqual(%d)", vr.Length)
}

// MapLenMin creates a [MapLenMinValidator].
func MapLenMin[M ~map[K]V, K comparable, V any](minLen int) *MapLenMinValidator[M, K, V] {
	return &MapLenMinValidator[M, K, V]{
		Min: minLen,
	}
}

// MapLenMinValidator is a [Validator] that checks if the length of the map is greater than or equal to the minimum length.
type MapLenMinValidator[M ~map[K]V, K comparable, V any] struct {
	Min int
}

// Validate implements [Validator].
func (vr *MapLenMinValidator[M, K, V]) Validate(m M) error {
	return validateLenMin(len(m), vr.Min)
}

func (vr *MapLenMinValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapLenMin(%d)", vr.Min)
}

// MapLenMax creates a [MapLenMaxValidator].
func MapLenMax[M ~map[K]V, K comparable, V any](maxLen int) *MapLenMaxValidator[M, K, V] {
	return &MapLenMaxValidator[M, K, V]{
		Max: maxLen,
	}
}

// MapLenMaxValidator is a [Validator] that checks if the length of the map is less than or equal to the maximum length.
type MapLenMaxValidator[M ~map[K]V, K comparable, V any] struct {
	Max int
}

// Validate implements [Validator].
func (vr *MapLenMaxValidator[M, K, V]) Validate(m M) error {
	return validateLenMax(len(m), vr.Max)
}

func (vr *MapLenMaxValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapLenMax(%d)", vr.Max)
}

// MapLenRange creates a [MapLenRangeValidator].
func MapLenRange[M ~map[K]V, K comparable, V any](minLen, maxLen int) *MapLenRangeValidator[M, K, V] {
	return &MapLenRangeValidator[M, K, V]{
		Min: minLen,
		Max: maxLen,
	}
}

// MapLenRangeValidator is a [Validator] that checks if the length of the map is within the range.
type MapLenRangeValidator[M ~map[K]V, K comparable, V any] struct {
	Min int
	Max int
}

// Validate implements [Validator].
func (vr *MapLenRangeValidator[M, K, V]) Validate(m M) error {
	return validateLenRange(len(m), vr.Min, vr.Max)
}

func (vr *MapLenRangeValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapLenRange(%d, %d)", vr.Min, vr.Max)
}

// MapEmpty creates a [MapEmptyValidator].
func MapEmpty[M ~map[K]V, K comparable, V any]() *MapEmptyValidator[M, K, V] {
	return &MapEmptyValidator[M, K, V]{}
}

// MapEmptyValidator is a [Validator] that checks if the map is empty.
type MapEmptyValidator[M ~map[K]V, K comparable, V any] struct{}

// Validate implements [Validator].
func (vr *MapEmptyValidator[M, K, V]) Validate(m M) error {
	return validateEmpty(len(m))
}

func (vr *MapEmptyValidator[M, K, V]) String() string {
	return "MapEmpty"
}

// MapNotEmpty creates a [MapNotEmptyValidator].
func MapNotEmpty[M ~map[K]V, K comparable, V any]() *MapNotEmptyValidator[M, K, V] {
	return &MapNotEmptyValidator[M, K, V]{}
}

// MapNotEmptyValidator is a [Validator] that checks if the map is not empty.
type MapNotEmptyValidator[M ~map[K]V, K comparable, V any] struct{}

// Validate implements [Validator].
func (vr *MapNotEmptyValidator[M, K, V]) Validate(m M) error {
	return validateNotEmpty(len(m))
}

func (vr *MapNotEmptyValidator[M, K, V]) String() string {
	return "MapNotEmpty"
}

// MapEach creates a [MapEachValidator].
func MapEach[M ~map[K]V, K comparable, V any](vr Validator[KeyValue[K, V]]) *MapEachValidator[M, K, V] {
	return &MapEachValidator[M, K, V]{
		Validator: vr,
	}
}

// MapEachValidator is a [Validator] that checks each key and value of the map.
type MapEachValidator[M ~map[K]V, K comparable, V any] struct {
	Validator Validator[KeyValue[K, V]]
}

// Validate implements [Validator].
func (vr *MapEachValidator[M, K, V]) Validate(m M) error {
	return validateMapEach(m, vr.Validator.Validate)
}

func (vr *MapEachValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapEach(%v)", vr.Validator)
}

// MapEachKey creates a [MapEachKeyValidator].
func MapEachKey[M ~map[K]V, K comparable, V any](vr Validator[K]) *MapEachKeyValidator[M, K, V] {
	return &MapEachKeyValidator[M, K, V]{
		Validator: vr,
	}
}

// MapEachKeyValidator is a [Validator] that checks each key of the map.
type MapEachKeyValidator[M ~map[K]V, K comparable, V any] struct {
	Validator Validator[K]
}

// Validate implements [Validator].
func (vr *MapEachKeyValidator[M, K, V]) Validate(m M) error {
	return validateMapEach(m, field("key", KeyValue[K, V].GetKey, vr.Validator.Validate))
}

func (vr *MapEachKeyValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapEachKey(%v)", vr.Validator)
}

// MapEachValue creates a [MapEachValueValidator].
func MapEachValue[M ~map[K]V, K comparable, V any](vr Validator[V]) *MapEachValueValidator[M, K, V] {
	return &MapEachValueValidator[M, K, V]{
		Validator: vr,
	}
}

// MapEachValueValidator is a [Validator] that checks each value of the map.
type MapEachValueValidator[M ~map[K]V, K comparable, V any] struct {
	Validator Validator[V]
}

// Validate implements [Validator].
func (vr *MapEachValueValidator[M, K, V]) Validate(m M) error {
	return validateMapEach(m, field("value", KeyValue[K, V].GetValue, vr.Validator.Validate))
}

func (vr *MapEachValueValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapEachValue(%v)", vr.Validator)
}

func validateMapEach[M ~map[K]V, K comparable, V any](m M, f func(KeyValue[K, V]) error) error {
	var errs []error
	for k, v := range m {
		err := f(KeyValue[K, V]{Key: k, Value: v})
		if err != nil {
			errs = append(errs, ErrorWrapPathElem(err, &KeyPathElem{Key: k}))
		}
	}
	return ErrorJoin(errs...)
}

// MapSortedEach creates a [MapSortedEachValidator].
func MapSortedEach[M ~map[K]V, K cmp.Ordered, V any](vr Validator[KeyValue[K, V]]) *MapSortedEachValidator[M, K, V] {
	return &MapSortedEachValidator[M, K, V]{
		Validator: vr,
	}
}

// MapSortedEachValidator is a [Validator] that checks each key and value of the map in sorted order of keys.
type MapSortedEachValidator[M ~map[K]V, K cmp.Ordered, V any] struct {
	Validator Validator[KeyValue[K, V]]
}

// Validate implements [Validator].
func (vr *MapSortedEachValidator[M, K, V]) Validate(m M) error {
	return validateMapSortedEach(m, vr.Validator.Validate)
}

func (vr *MapSortedEachValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapSortedEach(%v)", vr.Validator)
}

// MapSortedEachKey creates a [MapSortedEachKeyValidator].
func MapSortedEachKey[M ~map[K]V, K cmp.Ordered, V any](vr Validator[K]) *MapSortedEachKeyValidator[M, K, V] {
	return &MapSortedEachKeyValidator[M, K, V]{
		Validator: vr,
	}
}

// MapSortedEachKeyValidator is a [Validator] that checks each key of the map in sorted order of keys.
type MapSortedEachKeyValidator[M ~map[K]V, K cmp.Ordered, V any] struct {
	Validator Validator[K]
}

// Validate implements [Validator].
func (vr *MapSortedEachKeyValidator[M, K, V]) Validate(m M) error {
	return validateMapSortedEach(m, field("key", KeyValue[K, V].GetKey, vr.Validator.Validate))
}

func (vr *MapSortedEachKeyValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapSortedEachKey(%v)", vr.Validator)
}

// MapSortedEachValue creates a [MapSortedEachValueValidator].
func MapSortedEachValue[M ~map[K]V, K cmp.Ordered, V any](vr Validator[V]) *MapSortedEachValueValidator[M, K, V] {
	return &MapSortedEachValueValidator[M, K, V]{
		Validator: vr,
	}
}

// MapSortedEachValueValidator is a [Validator] that checks each value of the map in sorted order of keys.
type MapSortedEachValueValidator[M ~map[K]V, K cmp.Ordered, V any] struct {
	Validator Validator[V]
}

// Validate implements [Validator].
func (vr *MapSortedEachValueValidator[M, K, V]) Validate(m M) error {
	return validateMapSortedEach(m, field("value", KeyValue[K, V].GetValue, vr.Validator.Validate))
}

func (vr *MapSortedEachValueValidator[M, K, V]) String() string {
	return fmt.Sprintf("MapSortedEachValue(%v)", vr.Validator)
}

func validateMapSortedEach[M ~map[K]V, K cmp.Ordered, V any](m M, f func(KeyValue[K, V]) error) error {
	var errs []error
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		v := m[k]
		err := f(KeyValue[K, V]{Key: k, Value: v})
		if err != nil {
			errs = append(errs, ErrorWrapPathElem(err, &KeyPathElem{Key: k}))
		}
	}
	return ErrorJoin(errs...)
}
