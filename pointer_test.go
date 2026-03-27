package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExamplePointerOptional() {
	vr := PointerOptional(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(new(1)))
	fmt.Println(vr.Validate(nil))
	fmt.Println(vr.Validate(new(2)))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate(new(2)), "en"))
	// Output:
	// PointerOptional(Equal(1))
	// <nil>
	// <nil>
	// path pointer: 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExamplePointerRequired() {
	vr := PointerRequired(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(new(1)))
	fmt.Println(vr.Validate(nil))
	fmt.Println(vr.Validate(new(2)))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate(nil), "en"))
	// Output:
	// PointerRequired(Equal(1))
	// <nil>
	// pointer is nil
	// path pointer: 2 is not equal to 1
	// Pointer is nil.
}
