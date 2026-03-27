package vld_test

import (
	"errors"
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleGetLocalizedMessage() {
	fmt.Printf("%#v\n", GetLocalizedMessage("EqualError", []any{1, 2}, "fr"))
	fmt.Printf("%#v\n", GetLocalizedMessage("_Unknown", nil, "en"))
	fmt.Printf("%#v\n", GetLocalizedMessage("EqualError", []any{1, 2}, "unknown"))
	// Output:
	// "La valeur 1 n'est pas égale à 2."
	// ""
	// ""
}

func ExampleGetErrorLocalizedMessage() {
	fmt.Printf("%#v\n", GetErrorLocalizedMessage(&EqualError[int]{Value: 1, Expected: 2}, "fr"))
	fmt.Printf("%#v\n", GetErrorLocalizedMessage(nil, "en"))
	fmt.Printf("%#v\n", GetErrorLocalizedMessage(errors.New("error"), "en"))
	// Output:
	// "La valeur 1 n'est pas égale à 2."
	// ""
	// ""
}
