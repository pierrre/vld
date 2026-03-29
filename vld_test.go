package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleValidatorFunc() {
	vr := ValidatorFunc[int](Equal(1).Validate)
	fmt.Println(vr)
	fmt.Println(vr.Validate(1))
	fmt.Println(vr.Validate(2))
	fmt.Println(LocalizeError(vr.Validate(2), "en"))
	// Output:
	// ValidatorFunc
	// <nil>
	// 2 is not equal to 1
	// Value 2 is not equal to 1.
}
