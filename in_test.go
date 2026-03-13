package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleIn() {
	vr := In(1, 2, 3)
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(4))
	// Output:
	// In([]int{1, 2, 3})
	// <nil>
	// 4 is not in []int{1, 2, 3}
}

func ExampleNotIn() {
	vr := NotIn(1, 2, 3)
	fmt.Println(vr)
	fmt.Println(vr.Validate(4))
	fmt.Println(vr.Validate(1))
	// Output:
	// NotIn([]int{1, 2, 3})
	// <nil>
	// 1 is in []int{1, 2, 3}
}
