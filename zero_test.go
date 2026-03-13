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
	// Output:
	// Zero
	// <nil>
	// 1 is not zero
}

func ExampleNotZero() {
	vr := NotZero[int]()
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	// Output:
	// NotZero
	// <nil>
	// is zero
}

func ExampleOptional() {
	vr := Optional(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	// Output:
	// Optional(Equal(1))
	// <nil>
	// <nil>
	// 2 is not equal to 1
}

func ExampleRequired() {
	vr := Required(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(2))
	// Output:
	// Required(Equal(1))
	// <nil>
	// required
	// 2 is not equal to 1
}
