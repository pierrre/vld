package vld_test

import (
	"cmp"
	"fmt"

	. "github.com/pierrre/vld"
)

type intComparer int

func (i intComparer) Compare(other intComparer) int {
	return cmp.Compare(i, other)
}

func ExampleCmpEqual() {
	vr := CmpEqual(intComparer(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(intComparer(1)))
	fmt.Println(vr.Validate(intComparer(2)))
	// Output:
	// CmpEqual(1)
	// <nil>
	// 2 is not equal to 1
}

func ExampleCmpNotEqual() {
	vr := CmpNotEqual(intComparer(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(intComparer(2)))
	fmt.Println(vr.Validate(intComparer(1)))
	// Output:
	// CmpNotEqual(1)
	// <nil>
	// 1 is equal to 1
}

func ExampleCmpMin() {
	vr := CmpMin(intComparer(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(intComparer(2)))
	fmt.Println(vr.Validate(intComparer(1)))
	fmt.Println(vr.Validate(intComparer(0)))
	// Output:
	// CmpMin(1)
	// <nil>
	// <nil>
	// 0 is less than 1
}

func ExampleCmpMax() {
	vr := CmpMax(intComparer(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(intComparer(0)))
	fmt.Println(vr.Validate(intComparer(1)))
	fmt.Println(vr.Validate(intComparer(2)))
	// Output:
	// CmpMax(1)
	// <nil>
	// <nil>
	// 2 is greater than 1
}

func ExampleCmpRange() {
	vr := CmpRange(intComparer(1), intComparer(3))
	fmt.Println(vr)
	fmt.Println(vr.Validate(intComparer(2)))
	fmt.Println(vr.Validate(intComparer(1)))
	fmt.Println(vr.Validate(intComparer(3)))
	fmt.Println(vr.Validate(intComparer(0)))
	fmt.Println(vr.Validate(intComparer(4)))
	// Output:
	// CmpRange(1, 3)
	// <nil>
	// <nil>
	// <nil>
	// 0 is not in the range [1, 3]
	// 4 is not in the range [1, 3]
}

func ExampleCmpLess() {
	vr := CmpLess(intComparer(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(intComparer(0)))
	fmt.Println(vr.Validate(intComparer(1)))
	fmt.Println(vr.Validate(intComparer(2)))
	// Output:
	// CmpLess(1)
	// <nil>
	// 1 is not less than 1
	// 2 is not less than 1
}

func ExampleCmpGreater() {
	vr := CmpGreater(intComparer(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(intComparer(2)))
	fmt.Println(vr.Validate(intComparer(1)))
	fmt.Println(vr.Validate(intComparer(0)))
	// Output:
	// CmpGreater(1)
	// <nil>
	// 1 is not greater than 1
	// 0 is not greater than 1
}
