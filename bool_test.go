package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleAnd() {
	vr := And(Min(1), Max(10))
	fmt.Println(vr)
	fmt.Println(vr.Validate(5))
	fmt.Println(vr.Validate(0))
	// Output:
	// And(
	// 	Min(1),
	// 	Max(10),
	// )
	// <nil>
	// 0 is less than 1
}

func ExampleOr() {
	vr := Or(Min(10), Max(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(11))
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(5))
	// Output:
	// Or(
	// 	Min(10),
	// 	Max(1),
	// )
	// <nil>
	// <nil>
	// 5 is less than 10
	// 5 is greater than 1
}

func ExampleAll() {
	vr := All(Min(1), Max(10))
	fmt.Println(vr)
	fmt.Println(vr.Validate(5))
	fmt.Println(vr.Validate(0))
	// Output:
	// All(
	// 	Min(1),
	// 	Max(10),
	// )
	// <nil>
	// 0 is less than 1
}

func ExampleNot() {
	vr := Not("custom error", Min(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(5))
	// Output:
	// Not("custom error", Min(1))
	// <nil>
	// custom error
}

func ExampleIf() {
	cond := func(v int) bool { return v != 0 }
	vr := If(cond, Min(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(0))
	fmt.Println(vr.Validate(5))
	fmt.Println(vr.Validate(-1))
	// Output:
	// If(condition, Min(1))
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
	// IfElse(condition, Max(10), Min(-10))
	// <nil>
	// 11 is greater than 10
	// <nil>
	// -11 is less than -10
}
