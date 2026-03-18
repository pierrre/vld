package vld

import (
	"fmt"
	"slices"
)

func getSliceLen[S ~[]E, E any](s S) int {
	return len(s)
}

// SliceLenEqual returns a [Validator] that checks if the length of the slice is equal to the specified length.
func SliceLenEqual[S ~[]E, E any](length int) Validator[S] {
	return WithStringFunc(func() string { return fmt.Sprintf("SliceLenEqual(%d)", length) }, lenEqual(length, getSliceLen[S]))
}

// SliceLenMin returns a [Validator] that checks if the length of the slice is greater than or equal to the minimum length.
func SliceLenMin[S ~[]E, E any](minLen int) Validator[S] {
	return WithStringFunc(func() string { return fmt.Sprintf("SliceLenMin(%d)", minLen) }, lenMin(minLen, getSliceLen[S]))
}

// SliceLenMax returns a [Validator] that checks if the length of the slice is less than or equal to the maximum length.
func SliceLenMax[S ~[]E, E any](maxLen int) Validator[S] {
	return WithStringFunc(func() string { return fmt.Sprintf("SliceLenMax(%d)", maxLen) }, lenMax(maxLen, getSliceLen[S]))
}

// SliceLenRange returns a [Validator] that checks if the length of the slice is within the range.
func SliceLenRange[S ~[]E, E any](minLen, maxLen int) Validator[S] {
	return WithStringFunc(func() string { return fmt.Sprintf("SliceLenRange(%d, %d)", minLen, maxLen) }, lenRange(minLen, maxLen, getSliceLen[S]))
}

// SliceEmpty returns a [Validator] that checks if the slice is empty.
func SliceEmpty[S ~[]E, E any]() Validator[S] {
	return WithStringFunc(func() string { return "SliceEmpty" }, empty(getSliceLen[S]))
}

// SliceNotEmpty returns a [Validator] that checks if the slice is not empty.
func SliceNotEmpty[S ~[]E, E any]() Validator[S] {
	return WithStringFunc(func() string { return "SliceNotEmpty" }, notEmpty(getSliceLen[S]))
}

// SliceContains returns a [Validator] that checks if the slice contains the element.
func SliceContains[S ~[]E, E comparable](elem E) Validator[S] {
	return WithStringFunc(func() string { return fmt.Sprintf("SliceContains(%#v)", elem) }, func(s S) error {
		if !slices.Contains(s, elem) {
			err := fmt.Errorf("does not contain %#v", elem)
			err = ErrorWrapLocalization(err, "SliceContains", elem)
			return err
		}
		return nil
	})
}

// SliceNotContains returns a [Validator] that checks if the slice does not contain the element.
func SliceNotContains[S ~[]E, E comparable](elem E) Validator[S] {
	return WithStringFunc(func() string { return fmt.Sprintf("SliceNotContains(%#v)", elem) }, func(s S) error {
		if slices.Contains(s, elem) {
			err := fmt.Errorf("contains %#v", elem)
			err = ErrorWrapLocalization(err, "SliceNotContains", elem)
			return err
		}
		return nil
	})
}

// SliceEach returns a [Validator] that checks each index and element of the slice.
func SliceEach[S ~[]E, E any](vr Validator[KeyValue[int, E]]) Validator[S] {
	return WithStringFunc(func() string { return fmt.Sprintf("SliceEach(%v)", vr) }, sliceEach[S](vr.Validate))
}

// SliceEachValue returns a [Validator] that checks each element of the slice.
func SliceEachValue[S ~[]E, E any](vr Validator[E]) Validator[S] {
	return WithStringFunc(func() string { return fmt.Sprintf("SliceEachValue(%v)", vr) }, sliceEach[S](get(KeyValue[int, E].GetValue, vr.Validate)))
}

func sliceEach[S ~[]E, E any](f func(KeyValue[int, E]) error) func(S) error {
	return func(s S) error {
		var errs []error
		for i, v := range s {
			err := f(KeyValue[int, E]{Key: i, Value: v})
			if err != nil {
				errs = append(errs, ErrorWrapPathElement(err, &IndexPathElem{Index: i}))
			}
		}
		return ErrorJoin(errs...)
	}
}

// SliceUnique returns a [Validator] that checks if all elements of the slice are unique.
func SliceUnique[S ~[]E, E comparable]() Validator[S] {
	return WithStringFunc(func() string { return "SliceUnique" }, func(s S) error {
		seen := make(map[E]int, len(s))
		var errs []error
		for i, v := range s {
			_, ok := seen[v]
			if ok {
				err := fmt.Errorf("duplicate %#v (index %d)", v, seen[v])
				err = ErrorWrapLocalization(err, "SliceUnique", v, seen[v])
				errs = append(errs, &PathElemError{
					Err:      err,
					PathElem: &IndexPathElem{Index: i},
				})
			} else {
				seen[v] = i
			}
		}
		return ErrorJoin(errs...)
	})
}

// SliceUniqueBy returns a [Validator] that checks if all elements of the slice are unique by the key returned by the key function.
func SliceUniqueBy[S ~[]E, E any, K comparable](getKey func(E) K) Validator[S] {
	return WithStringFunc(func() string { return "SliceUniqueBy" }, func(s S) error {
		seen := make(map[K]int, len(s))
		var errs []error
		for i, v := range s {
			key := getKey(v)
			_, ok := seen[key]
			if ok {
				err := fmt.Errorf("duplicate %#v (index %d)", v, seen[key])
				err = ErrorWrapLocalization(err, "SliceUnique", v, seen[key])
				errs = append(errs, &PathElemError{
					Err:      err,
					PathElem: &IndexPathElem{Index: i},
				})
			} else {
				seen[key] = i
			}
		}
		return ErrorJoin(errs...)
	})
}
