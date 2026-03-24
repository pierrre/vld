package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleIf() {
	cond := func(v int) bool { return v != 0 }
	vr := If(cond, Min(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(5))
	fmt.Println(vr.Validate(-1))
	// Output:
	// If(github.com/pierrre/vld_test.ExampleIf.func1, Min(1))
	// <nil>
	// <nil>
	// -1 is less than 1
}

func ExampleIfElse() {
	cond := func(v int) bool { return v > 0 }
	vr := IfElse(cond, Max(10), Min(-10))
	fmt.Println(vr)
	fmt.Println(vr.Validate(5))
	fmt.Println(vr.Validate(11))
	fmt.Println(vr.Validate(-5))
	fmt.Println(vr.Validate(-11))
	// Output:
	// IfElse(github.com/pierrre/vld_test.ExampleIfElse.func1, Max(10), Min(-10))
	// <nil>
	// 11 is greater than 10
	// <nil>
	// -11 is less than -10
}

func ExampleSwitch() {
	vr := Switch(
		Case(func(v int) bool { return v > 0 }, Max(10)),
		Case(func(v int) bool { return v < 0 }, Min(-10)),
	)
	fmt.Println(vr)
	fmt.Println(vr.Validate(5))
	fmt.Println(vr.Validate(-5))
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(11))
	fmt.Println(vr.Validate(-11))

	// Output:
	// Switch(
	// 	Case(github.com/pierrre/vld_test.ExampleSwitch.func1, Max(10)),
	// 	Case(github.com/pierrre/vld_test.ExampleSwitch.func2, Min(-10)),
	// )
	// <nil>
	// <nil>
	// <nil>
	// 11 is greater than 10
	// -11 is less than -10
}
