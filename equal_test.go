package vld_test

import (
	"cmp"
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleEqual() {
	vr := Equal(1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(GetErrorLocalization(vr.Validate(2), "en"))
	// Output:
	// Equal(1)
	// <nil>
	// 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleEqualFunc() {
	vr := EqualFunc(1, func(a, b int) bool { return a == b })
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(GetErrorLocalization(vr.Validate(2), "en"))
	// Output:
	// EqualFunc(1)
	// <nil>
	// 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleEqualCmpFunc() {
	vr := EqualCmpFunc(1, cmp.Compare)
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(GetErrorLocalization(vr.Validate(2), "en"))
	// Output:
	// EqualCmpFunc(1)
	// <nil>
	// 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleNotEqual() {
	vr := NotEqual(1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(GetErrorLocalization(vr.Validate(1), "en"))
	// Output:
	// NotEqual(1)
	// <nil>
	// 1 is equal to 1
	// Value 1 is equal to 1.
}

func ExampleNotEqualFunc() {
	vr := NotEqualFunc(1, func(a, b int) bool { return a == b })
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(GetErrorLocalization(vr.Validate(1), "en"))
	// Output:
	// NotEqualFunc(1)
	// <nil>
	// 1 is equal to 1
	// Value 1 is equal to 1.
}

func ExampleNotEqualCmpFunc() {
	vr := NotEqualCmpFunc(1, cmp.Compare)
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(GetErrorLocalization(vr.Validate(1), "en"))
	// Output:
	// NotEqualCmpFunc(1)
	// <nil>
	// 1 is equal to 1
	// Value 1 is equal to 1.
}
