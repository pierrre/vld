package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleIn() {
	vr := In(1, 2, 3)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(4))
	fmt.Println(LocalizeError(vr.Validate(4), "en"))
	// Output:
	// In([]int{1, 2, 3})
	// Value must be in []int{1, 2, 3}.
	// <nil>
	// 4 is not in []int{1, 2, 3}
	// Value 4 is not in []int{1, 2, 3}.
}

func ExampleNotIn() {
	vr := NotIn(1, 2, 3)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(4))
	fmt.Println(vr.Validate(1))
	fmt.Println(LocalizeError(vr.Validate(1), "en"))
	// Output:
	// NotIn([]int{1, 2, 3})
	// Value must not be in []int{1, 2, 3}.
	// <nil>
	// 1 is in []int{1, 2, 3}
	// Value 1 is in []int{1, 2, 3}.
}
