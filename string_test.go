package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleStringLenEqual() {
	vr := StringLenEqual(3)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abcd"))
	// Output:
	// StringLenEqual(3)
	// <nil>
	// length 4 is not equal to 3
}

func ExampleStringLenMin() {
	vr := StringLenMin(3)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("ab"))
	// Output:
	// StringLenMin(3)
	// <nil>
	// length 2 is less than 3
}

func ExampleStringLenMax() {
	vr := StringLenMax(3)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abcd"))
	// Output:
	// StringLenMax(3)
	// <nil>
	// length 4 is greater than 3
}

func ExampleStringLenRange() {
	vr := StringLenRange(2, 4)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("a"))
	fmt.Println(vr.Validate("abcde"))
	// Output:
	// StringLenRange(2, 4)
	// <nil>
	// length 1 is not in the range [2, 4]
	// length 5 is not in the range [2, 4]
}

func ExampleStringEmpty() {
	vr := StringEmpty()
	fmt.Println(vr)
	fmt.Println(vr.Validate(""))
	fmt.Println(vr.Validate("a"))
	// Output:
	// StringEmpty
	// <nil>
	// is not empty (1)
}

func ExampleStringNotEmpty() {
	vr := StringNotEmpty()
	fmt.Println(vr)
	fmt.Println(vr.Validate("a"))
	fmt.Println(vr.Validate(""))
	// Output:
	// StringNotEmpty
	// <nil>
	// is empty
}

func ExampleStringRunesEqual() {
	vr := StringRunesEqual(3)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abcd"))
	// Output:
	// StringRunesEqual(3)
	// <nil>
	// length 4 is not equal to 3
}

func ExampleStringRunesMin() {
	vr := StringRunesMin(3)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("ab"))
	// Output:
	// StringRunesMin(3)
	// <nil>
	// length 2 is less than 3
}

func ExampleStringRunesMax() {
	vr := StringRunesMax(3)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abcd"))
	// Output:
	// StringRunesMax(3)
	// <nil>
	// length 4 is greater than 3
}

func ExampleStringRunesRange() {
	vr := StringRunesRange(2, 4)
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("a"))
	fmt.Println(vr.Validate("abcde"))
	// Output:
	// StringRunesRange(2, 4)
	// <nil>
	// length 1 is not in the range [2, 4]
	// length 5 is not in the range [2, 4]
}

func ExampleStringContains() {
	vr := StringContains("b")
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("acd"))
	// Output:
	// StringContains("b")
	// <nil>
	// "acd" does not contain "b"
}

func ExampleStringNotContains() {
	vr := StringNotContains("b")
	fmt.Println(vr)
	fmt.Println(vr.Validate("acd"))
	fmt.Println(vr.Validate("abc"))
	// Output:
	// StringNotContains("b")
	// <nil>
	// "abc" contains "b"
}

func ExampleStringHasPrefix() {
	vr := StringHasPrefix("a")
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("bac"))
	// Output:
	// StringHasPrefix("a")
	// <nil>
	// "bac" does not begin with "a"
}

func ExampleStringNotHasPrefix() {
	vr := StringNotHasPrefix("a")
	fmt.Println(vr)
	fmt.Println(vr.Validate("bac"))
	fmt.Println(vr.Validate("abc"))
	// Output:
	// StringNotHasPrefix("a")
	// <nil>
	// "abc" begins with "a"
}

func ExampleStringHasSuffix() {
	vr := StringHasSuffix("c")
	fmt.Println(vr)
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("acb"))
	// Output:
	// StringHasSuffix("c")
	// <nil>
	// "acb" does not end with "c"
}

func ExampleStringNotHasSuffix() {
	vr := StringNotHasSuffix("c")
	fmt.Println(vr)
	fmt.Println(vr.Validate("acb"))
	fmt.Println(vr.Validate("abc"))
	// Output:
	// StringNotHasSuffix("c")
	// <nil>
	// "abc" ends with "c"
}
