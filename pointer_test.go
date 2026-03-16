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
	// Output:
	// PointerOptional(Equal(1))
	// <nil>
	// <nil>
	// path pointer: 2 is not equal to 1
}

func ExamplePointerRequired() {
	vr := PointerRequired(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(new(1)))
	fmt.Println(vr.Validate(nil))
	fmt.Println(vr.Validate(new(2)))
	// Output:
	// PointerRequired(Equal(1))
	// <nil>
	// pointer is nil
	// path pointer: 2 is not equal to 1
}
