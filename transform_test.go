package vld_test

import (
	"errors"
	"fmt"
	"strconv"

	. "github.com/pierrre/vld"
)

func ExampleGet() {
	type User struct {
		Name string
	}
	vr := Get(func(u User) string { return u.Name }, StringLenMax(5))
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(User{Name: "Alice"}))
	fmt.Println(vr.Validate(User{Name: "Charlie"}))
	fmt.Println(LocalizeError(vr.Validate(User{Name: "Charlie"}), "en"))
	// Output:
	// Get(github.com/pierrre/vld_test.ExampleGet.func1, StringLenMax(5))
	// Value returned by function github.com/pierrre/vld_test.ExampleGet.func1 must satisfy the following validator: Length must be less than or equal to 5.
	// <nil>
	// length 7 is greater than 5
	// Length 7 is greater than 5.
}

func ExampleParse() {
	vr := Parse(strconv.Atoi, Equal(1))
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("1"))
	fmt.Println(vr.Validate("2"))
	fmt.Println(vr.Validate("a"))
	fmt.Println(errors.Unwrap(vr.Validate("a")))
	fmt.Println(LocalizeError(vr.Validate("a"), "en"))
	// Output:
	// Parse(strconv.Atoi, Equal(1))
	// Parsing value with function strconv.Atoi must succeed and the parsed value must satisfy the following validator: Value must be equal to 1.
	// <nil>
	// 2 is not equal to 1
	// parse "a" with strconv.Atoi: strconv.Atoi: parsing "a": invalid syntax
	// strconv.Atoi: parsing "a": invalid syntax
	// Parsing value "a" with function strconv.Atoi failed: strconv.Atoi: parsing "a": invalid syntax
}

func ExampleWrap() {
	vr := Wrap("message", Equal(1))
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(LocalizeError(vr.Validate(2), "en"))
	// Output:
	// Wrap("message", Equal(1))
	// Value must satisfy the following validator, with error wrapped by message "message": Value must be equal to 1.
	// <nil>
	// message: 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleField() {
	type User struct {
		Name string
	}
	vr := Field("Name", func(u User) string { return u.Name }, StringLenMax(5))
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(User{Name: "Alice"}))
	fmt.Println(vr.Validate(User{Name: "Charlie"}))
	fmt.Println(LocalizeError(vr.Validate(User{Name: "Charlie"}), "en"))
	// Output:
	// Field("Name", github.com/pierrre/vld_test.ExampleField.func1, StringLenMax(5))
	// Value of field "Name" returned by function github.com/pierrre/vld_test.ExampleField.func1 must satisfy the following validator: Length must be less than or equal to 5.
	// <nil>
	// path field "Name": length 7 is greater than 5
	// Length 7 is greater than 5.
}

func ExampleMessage() {
	vr := Message("message", Equal(1))
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	// Output:
	// Message("message", Equal(1))
	// Value must satisfy the following validator, with error message overridden by "message": Value must be equal to 1.
	// <nil>
	// message
}
