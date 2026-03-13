package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleRegexpMatch() {
	vr := RegexpMatch(`^[a-z]+$`)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abc123"))
	// Output:
	// RegexpMatch("^[a-z]+$")
	// <nil>
	// "abc123" does not match regexp "^[a-z]+$"
}

func ExampleRegexpNotMatch() {
	vr := RegexpNotMatch(`^[a-z]+$`)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc123"))
	fmt.Println(vr.Validate("abc"))
	// Output:
	// RegexpNotMatch("^[a-z]+$")
	// <nil>
	// "abc" matches regexp "^[a-z]+$"
}
