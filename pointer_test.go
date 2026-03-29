package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExamplePointerOptional() {
	vr := PointerOptional(Equal(1))
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(new(1)))
	fmt.Println(vr.Validate(nil))
	fmt.Println(vr.Validate(new(2)))
	fmt.Println(LocalizeError(vr.Validate(new(2)), "en"))
	// Output:
	// PointerOptional(Equal(1))
	// Pointer must be nil or the dereferenced value must satisfy the following validator: Value must be equal to 1.
	// <nil>
	// <nil>
	// path pointer: 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExamplePointerRequired() {
	vr := PointerRequired(Equal(1))
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(new(1)))
	fmt.Println(vr.Validate(nil))
	fmt.Println(vr.Validate(new(2)))
	fmt.Println(LocalizeError(vr.Validate(nil), "en"))
	// Output:
	// PointerRequired(Equal(1))
	// Pointer must not be nil and the dereferenced value must satisfy the following validator: Value must be equal to 1.
	// <nil>
	// pointer is nil
	// path pointer: 2 is not equal to 1
	// Pointer is nil.
}
