package vld

import (
	"fmt"
)

func getMapLen[M ~map[K]V, K comparable, V any](m M) int {
	return len(m)
}

// MapLenEqual returns a [Validator] that checks if the length of the map is equal to the specified length.
func MapLenEqual[M ~map[K]V, K comparable, V any](length int) Validator[M] {
	return WithStringFunc(func() string { return fmt.Sprintf("MapLenEqual(%d)", length) }, lenEqual("length", length, getMapLen[M]))
}

// MapLenMin returns a [Validator] that checks if the length of the map is greater than or equal to the minimum length.
func MapLenMin[M ~map[K]V, K comparable, V any](minLen int) Validator[M] {
	return WithStringFunc(func() string { return fmt.Sprintf("MapLenMin(%d)", minLen) }, lenMin("length", minLen, getMapLen[M]))
}

// MapLenMax returns a [Validator] that checks if the length of the map is less than or equal to the maximum length.
func MapLenMax[M ~map[K]V, K comparable, V any](maxLen int) Validator[M] {
	return WithStringFunc(func() string { return fmt.Sprintf("MapLenMax(%d)", maxLen) }, lenMax("length", maxLen, getMapLen[M]))
}

// MapLenRange returns a [Validator] that checks if the length of the map is within the range.
func MapLenRange[M ~map[K]V, K comparable, V any](minLen, maxLen int) Validator[M] {
	return WithStringFunc(func() string { return fmt.Sprintf("MapLenRange(%d, %d)", minLen, maxLen) }, lenRange("length", minLen, maxLen, getMapLen[M]))
}

// MapEmpty returns a [Validator] that checks if the map is empty.
func MapEmpty[M ~map[K]V, K comparable, V any]() Validator[M] {
	return WithStringFunc(func() string { return "MapEmpty" }, empty(getMapLen[M]))
}

// MapNotEmpty returns a [Validator] that checks if the map is not empty.
func MapNotEmpty[M ~map[K]V, K comparable, V any]() Validator[M] {
	return WithStringFunc(func() string { return "MapNotEmpty" }, notEmpty(getMapLen[M]))
}

// MapEach returns a [Validator] that checks each key and value of the map.
func MapEach[M ~map[K]V, K comparable, V any](vr Validator[KeyValue[K, V]]) Validator[M] {
	return WithStringFunc(func() string { return fmt.Sprintf("MapEach(%v)", vr) }, mapEach[M](vr.Validate))
}

// MapEachKey returns a [Validator] that checks each key of the map.
func MapEachKey[M ~map[K]V, K comparable, V any](vr Validator[K]) Validator[M] {
	return WithStringFunc(func() string { return fmt.Sprintf("MapEachKey(%v)", vr) }, mapEach[M](get(KeyValue[K, V].GetKey, vr.Validate)))
}

// MapEachValue returns a [Validator] that checks each value of the map.
func MapEachValue[M ~map[K]V, K comparable, V any](vr Validator[V]) Validator[M] {
	return WithStringFunc(func() string { return fmt.Sprintf("MapEachValue(%v)", vr) }, mapEach[M](get(KeyValue[K, V].GetValue, vr.Validate)))
}

func mapEach[M ~map[K]V, K comparable, V any](f func(KeyValue[K, V]) error) func(M) error {
	return func(m M) error {
		var errs []error
		for k, v := range m {
			err := f(KeyValue[K, V]{Key: k, Value: v})
			if err != nil {
				errs = append(errs, ErrorWrapMessagef(err, "%#v", k))
			}
		}
		return ErrorJoin(errs...)
	}
}
