package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExamplePositive() {
	vr := Positive[int]()
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(-1))
	fmt.Println(LocalizeError(vr.Validate(0), "en"))
	// Output:
	// Positive
	// Value must be positive.
	// <nil>
	// 0 is not positive
	// -1 is not positive
	// Value 0 is not positive.
}

func ExampleNegative() {
	vr := Negative[int]()
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(-1))
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(1))
	fmt.Println(LocalizeError(vr.Validate(0), "en"))
	// Output:
	// Negative
	// Value must be negative.
	// <nil>
	// 0 is not negative
	// 1 is not negative
	// Value 0 is not negative.
}
