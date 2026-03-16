package vld_test

import (
	"errors"
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleErrorWrap() {
	wrapFunc := func(err error) error {
		return fmt.Errorf("example: %w", err)
	}
	fmt.Println(ErrorWrap(nil, wrapFunc))
	fmt.Println(ErrorWrap(errors.New("error"), wrapFunc))
	fmt.Println(ErrorWrap(errors.Join(errors.New("error 1"), errors.New("error 2")), wrapFunc))
	// Output:
	// <nil>
	// example: error
	// example: error 1
	// example: error 2
}

func ExampleErrorWrapMessage() {
	fmt.Println(ErrorWrapMessage(nil, "example"))
	fmt.Println(ErrorWrapMessage(errors.New("error"), "example"))
	// Output:
	// <nil>
	// example: error
}

func ExampleErrorWrapMessagef() {
	fmt.Println(ErrorWrapMessagef(nil, "example %d", 1))
	fmt.Println(ErrorWrapMessagef(errors.New("error"), "example %d", 1))
	// Output:
	// <nil>
	// example 1: error
}

func ExampleErrorJoin() {
	fmt.Println(ErrorJoin())
	fmt.Println(ErrorJoin(nil))
	fmt.Println(ErrorJoin(errors.Join(errors.New("error"))))
	// Output:
	// <nil>
	// <nil>
	// error
}

func ExampleGetErrors() {
	fmt.Println(GetErrors(nil))
	fmt.Println(GetErrors(errors.New("error")))
	fmt.Println(GetErrors(errors.Join(errors.New("error1"), errors.New("error2"))))
	// Output:
	// []
	// [error]
	// [error1 error2]
}
