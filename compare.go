package vld

import (
	"fmt"
)

// Comparer is an interface that defines a method for comparing values of type T.
type Comparer[T any] interface {
	Compare(v T) int
}

// CmpEqual returns a [Validator] that checks if the value is equal to a specific value using the Compare method.
func CmpEqual[T Comparer[T]](v T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("CmpEqual(%#v)", v) }, equalCmpFunc(v, T.Compare))
}

// CmpNotEqual returns a [Validator] that checks if the value is not equal to a specific value using the Compare method.
func CmpNotEqual[T Comparer[T]](v T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("CmpNotEqual(%#v)", v) }, notEqualCmpFunc(v, T.Compare))
}

// CmpMin returns a [Validator] that checks if the value is greater than or equal to the minimum value using the Compare method.
func CmpMin[T Comparer[T]](minValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("CmpMin(%#v)", minValue) }, minCmpFunc(minValue, T.Compare))
}

// CmpMax returns a [Validator] that checks if the value is less than or equal to the maximum value using the Compare method.
func CmpMax[T Comparer[T]](maxValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("CmpMax(%#v)", maxValue) }, maxCmpFunc(maxValue, T.Compare))
}

// CmpRange returns a [Validator] that checks if the value is within the range [minValue, maxValue] (inclusive) using the Compare method.
func CmpRange[T Comparer[T]](minValue, maxValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("CmpRange(%#v, %#v)", minValue, maxValue) }, rangeCmpFunc(minValue, maxValue, T.Compare))
}

// CmpLess returns a [Validator] that checks if the value is less than the maximum value using the Compare method.
func CmpLess[T Comparer[T]](maxValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("CmpLess(%#v)", maxValue) }, lessCmpFunc(maxValue, T.Compare))
}

// CmpGreater returns a [Validator] that checks if the value is greater than the minimum value using the Compare method.
func CmpGreater[T Comparer[T]](minValue T) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("CmpGreater(%#v)", minValue) }, greaterCmpFunc(minValue, T.Compare))
}
