package vld

import (
	"errors"
	"fmt"
)

func ExampleGetErrorLocalization() {
	fmt.Printf("%#v\n", GetErrorLocalization(ErrorWrapLocalization(errors.New("error"), "Equal", 1, 2), "fr"))
	fmt.Printf("%#v\n", GetErrorLocalization(ErrorWrapLocalization(errors.New("error"), "LenEqual", 1, 2), "en"))
	fmt.Printf("%#v\n", GetErrorLocalization(ErrorWrapLocalization(errors.New("error"), "LenEqual", 1, 2), "fr"))
	fmt.Printf("%#v\n", ErrorWrapLocalization(nil, ""))
	fmt.Printf("%#v\n", GetErrorLocalization(nil, "en"))
	fmt.Printf("%#v\n", GetErrorLocalization(errors.New("error"), "en"))
	fmt.Printf("%#v\n", GetErrorLocalization(ErrorWrapLocalization(errors.New("error"), "Unknown"), "en"))
	fmt.Printf("%#v\n", GetErrorLocalization(ErrorWrapLocalization(errors.New("error"), "Equal", 1, 2), "unknown"))
	// Output:
	// "La valeur 1 n'est pas égale à 2."
	// "Length 1 is not equal to 2."
	// "La longueur 1 n'est pas égale à 2."
	// <nil>
	// ""
	// ""
	// ""
	// ""
}
