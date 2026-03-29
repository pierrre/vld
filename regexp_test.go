package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleRegexpMatch() {
	vr := RegexpMatch(`^[a-z]+$`)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abc123"))
	fmt.Println(LocalizeError(vr.Validate("abc123"), "en"))
	// Output:
	// RegexpMatch("^[a-z]+$")
	// String must match regexp "^[a-z]+$".
	// <nil>
	// "abc123" does not match regexp "^[a-z]+$"
	// String "abc123" does not match regexp "^[a-z]+$".
}

func ExampleRegexpNotMatch() {
	vr := RegexpNotMatch(`^[a-z]+$`)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc123"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(LocalizeError(vr.Validate("abc"), "en"))
	// Output:
	// RegexpNotMatch("^[a-z]+$")
	// String must not match regexp "^[a-z]+$".
	// <nil>
	// "abc" matches regexp "^[a-z]+$"
	// String "abc" matches regexp "^[a-z]+$".
}
