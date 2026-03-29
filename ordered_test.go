package vld_test

import (
	"cmp"
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleMin() {
	vr := Min(1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	fmt.Println(LocalizeError(vr.Validate(0), "en"))
	// Output:
	// Min(1)
	// <nil>
	// <nil>
	// 0 is less than 1
	// Value 0 is less than 1.
}

func ExampleMinCmpFunc() {
	vr := MinCmpFunc(1, cmp.Compare)
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	fmt.Println(LocalizeError(vr.Validate(0), "en"))
	// Output:
	// MinCmpFunc(1, cmp.Compare[...])
	// <nil>
	// <nil>
	// 0 is less than 1
	// Value 0 is less than 1.
}

func ExampleMax() {
	vr := Max(1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(LocalizeError(vr.Validate(2), "en"))
	// Output:
	// Max(1)
	// <nil>
	// <nil>
	// 2 is greater than 1
	// Value 2 is greater than 1.
}

func ExampleMaxCmpFunc() {
	vr := MaxCmpFunc(1, cmp.Compare)
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(LocalizeError(vr.Validate(2), "en"))
	// Output:
	// MaxCmpFunc(1, cmp.Compare[...])
	// <nil>
	// <nil>
	// 2 is greater than 1
	// Value 2 is greater than 1.
}

func ExampleRange() {
	vr := Range(1, 3)
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(3))
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(4))
	fmt.Println(LocalizeError(vr.Validate(0), "en"))
	// Output:
	// Range(1, 3)
	// <nil>
	// <nil>
	// <nil>
	// 0 is not in the range [1, 3]
	// 4 is not in the range [1, 3]
	// Value 0 is not in the range [1, 3].
}

func ExampleRangeCmpFunc() {
	vr := RangeCmpFunc(1, 3, cmp.Compare)
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(3))
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(4))
	fmt.Println(LocalizeError(vr.Validate(0), "en"))
	// Output:
	// RangeCmpFunc(1, 3, cmp.Compare[...])
	// <nil>
	// <nil>
	// <nil>
	// 0 is not in the range [1, 3]
	// 4 is not in the range [1, 3]
	// Value 0 is not in the range [1, 3].
}

func ExampleLess() {
	vr := Less(1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(LocalizeError(vr.Validate(1), "en"))
	// Output:
	// Less(1)
	// <nil>
	// 1 is not less than 1
	// 2 is not less than 1
	// Value 1 is not less than 1.
}

func ExampleLessCmpFunc() {
	vr := LessCmpFunc(1, cmp.Compare)
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(LocalizeError(vr.Validate(1), "en"))
	// Output:
	// LessCmpFunc(1, cmp.Compare[...])
	// <nil>
	// 1 is not less than 1
	// 2 is not less than 1
	// Value 1 is not less than 1.
}

func ExampleGreater() {
	vr := Greater(1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	fmt.Println(LocalizeError(vr.Validate(1), "en"))
	// Output:
	// Greater(1)
	// <nil>
	// 1 is not greater than 1
	// 0 is not greater than 1
	// Value 1 is not greater than 1.
}

func ExampleGreaterCmpFunc() {
	vr := GreaterCmpFunc(1, cmp.Compare)
	fmt.Println(vr)
	fmt.Println(vr.Validate(2))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(0))
	fmt.Println(LocalizeError(vr.Validate(1), "en"))
	// Output:
	// GreaterCmpFunc(1, cmp.Compare[...])
	// <nil>
	// 1 is not greater than 1
	// 0 is not greater than 1
	// Value 1 is not greater than 1.
}
