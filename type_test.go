package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleTypeOptional() {
	vr := TypeOptional[any](Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate("1"))
	fmt.Println(vr.Validate(2))
	fmt.Println(GetErrorLocalization(vr.Validate(2), "en"))
	// Output:
	// TypeOptional[int](Equal(1))
	// <nil>
	// <nil>
	// 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleTypeRequired() {
	vr := TypeRequired[any](Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate("1"))
	fmt.Println(vr.Validate(2))
	fmt.Println(GetErrorLocalization(vr.Validate("1"), "en"))
	// Output:
	// TypeRequired[int](Equal(1))
	// <nil>
	// string cannot be converted to int
	// 2 is not equal to 1
	// Type string cannot be converted to int.
}
