package vld_test

import (
	"errors"
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleLocalize() {
	fmt.Printf("%#v\n", Localize("EqualError", []any{1, 2}, "fr"))
	fmt.Printf("%#v\n", Localize("_Unknown", nil, "en"))
	fmt.Printf("%#v\n", Localize("EqualError", []any{1, 2}, "unknown"))
	// Output:
	// "La valeur 1 n'est pas égale à 2."
	// ""
	// ""
}

func ExampleLocalizeError() {
	fmt.Printf("%#v\n", LocalizeError(&EqualError[int]{Value: 1, Expected: 2}, "fr"))
	fmt.Printf("%#v\n", LocalizeError(nil, "en"))
	fmt.Printf("%#v\n", LocalizeError(errors.New("error"), "en"))
	// Output:
	// "La valeur 1 n'est pas égale à 2."
	// ""
	// ""
}
