package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleGet() {
	type User struct {
		Name string
	}
	vr := Get(func(u User) string { return u.Name }, StringLenMax(5))
	fmt.Println(vr)
	fmt.Println(vr.Validate(User{Name: "Alice"}))
	fmt.Println(vr.Validate(User{Name: "Charlie"}))
	// Output:
	// Get(vld_test.User => string, StringLenMax(5))
	// <nil>
	// length 7 is greater than 5
}

func ExampleWrap() {
	vr := Wrap("message", Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	// Output:
	// Wrap("message", Equal(1))
	// <nil>
	// message: 2 is not equal to 1
}

func ExampleField() {
	type User struct {
		Name string
	}
	vr := Field("Name", func(u User) string { return u.Name }, StringLenMax(5))
	fmt.Println(vr)
	fmt.Println(vr.Validate(User{Name: "Alice"}))
	fmt.Println(vr.Validate(User{Name: "Charlie"}))
	// Output:
	// Field("Name", StringLenMax(5))
	// <nil>
	// path field "Name": length 7 is greater than 5
}

func ExampleMessage() {
	vr := Message("message", Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	// Output:
	// Message("message", Equal(1))
	// <nil>
	// message
}
