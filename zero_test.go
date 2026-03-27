package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleZero() {
	vr := Zero[int]()
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(1))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate(1), "en"))
	// Output:
	// Zero
	// <nil>
	// 1 is not zero
	// Value 1 is not zero.
}

func ExampleNotZero() {
	vr := NotZero[int]()
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate(0), "en"))
	// Output:
	// NotZero
	// <nil>
	// is zero
	// Value is zero.
}

func ExampleOptional() {
	vr := Optional(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate(2), "en"))
	// Output:
	// Optional(Equal(1))
	// <nil>
	// <nil>
	// 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleRequired() {
	vr := Required(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(2))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate(0), "en"))
	// Output:
	// Required(Equal(1))
	// <nil>
	// required
	// 2 is not equal to 1
	// Value is required.
}
