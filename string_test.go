package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleStringLenEqual() {
	vr := StringLenEqual(3)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abcd"))
	fmt.Println(LocalizeError(vr.Validate("abcd"), "en"))
	// Output:
	// StringLenEqual(3)
	// Length must be equal to 3.
	// <nil>
	// length 4 is not equal to 3
	// Length 4 is not equal to 3.
}

func ExampleStringLenMin() {
	vr := StringLenMin(3)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("ab"))
	fmt.Println(LocalizeError(vr.Validate("ab"), "en"))
	// Output:
	// StringLenMin(3)
	// Length must be greater than or equal to 3.
	// <nil>
	// length 2 is less than 3
	// Length 2 is less than 3.
}

func ExampleStringLenMax() {
	vr := StringLenMax(3)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abcd"))
	fmt.Println(LocalizeError(vr.Validate("abcd"), "en"))
	// Output:
	// StringLenMax(3)
	// Length must be less than or equal to 3.
	// <nil>
	// length 4 is greater than 3
	// Length 4 is greater than 3.
}

func ExampleStringLenRange() {
	vr := StringLenRange(2, 4)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("a"))
	fmt.Println(vr.Validate("abcde"))
	fmt.Println(LocalizeError(vr.Validate("a"), "en"))
	// Output:
	// StringLenRange(2, 4)
	// Length must be in the range [2, 4].
	// <nil>
	// length 1 is not in the range [2, 4]
	// length 5 is not in the range [2, 4]
	// Length 1 is not in the range [2, 4].
}

func ExampleStringEmpty() {
	vr := StringEmpty()
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(""))
	fmt.Println(vr.Validate("a"))
	fmt.Println(LocalizeError(vr.Validate("a"), "en"))
	// Output:
	// StringEmpty
	// Value must be empty.
	// <nil>
	// is not empty (1)
	// Value is not empty (1).
}

func ExampleStringNotEmpty() {
	vr := StringNotEmpty()
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("a"))
	fmt.Println(vr.Validate(""))
	fmt.Println(LocalizeError(vr.Validate(""), "en"))
	// Output:
	// StringNotEmpty
	// Value must not be empty.
	// <nil>
	// is empty
	// Value is empty.
}

func ExampleStringRunesEqual() {
	vr := StringRunesEqual(3)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abcd"))
	fmt.Println(LocalizeError(vr.Validate("abcd"), "en"))
	// Output:
	// StringRunesEqual(3)
	// Length must be equal to 3.
	// <nil>
	// length 4 is not equal to 3
	// Length 4 is not equal to 3.
}

func ExampleStringRunesMin() {
	vr := StringRunesMin(3)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("ab"))
	fmt.Println(LocalizeError(vr.Validate("ab"), "en"))
	// Output:
	// StringRunesMin(3)
	// Length must be greater than or equal to 3.
	// <nil>
	// length 2 is less than 3
	// Length 2 is less than 3.
}

func ExampleStringRunesMax() {
	vr := StringRunesMax(3)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("abcd"))
	fmt.Println(LocalizeError(vr.Validate("abcd"), "en"))
	// Output:
	// StringRunesMax(3)
	// Length must be less than or equal to 3.
	// <nil>
	// length 4 is greater than 3
	// Length 4 is greater than 3.
}

func ExampleStringRunesRange() {
	vr := StringRunesRange(2, 4)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("a"))
	fmt.Println(vr.Validate("abcde"))
	fmt.Println(LocalizeError(vr.Validate("a"), "en"))
	// Output:
	// StringRunesRange(2, 4)
	// Length must be in the range [2, 4].
	// <nil>
	// length 1 is not in the range [2, 4]
	// length 5 is not in the range [2, 4]
	// Length 1 is not in the range [2, 4].
}

func ExampleStringContains() {
	vr := StringContains("b")
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("acd"))
	fmt.Println(LocalizeError(vr.Validate("acd"), "en"))
	// Output:
	// StringContains("b")
	// String must contain "b".
	// <nil>
	// "acd" does not contain "b"
	// String "acd" does not contain "b".
}

func ExampleStringNotContains() {
	vr := StringNotContains("b")
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("acd"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(LocalizeError(vr.Validate("abc"), "en"))
	// Output:
	// StringNotContains("b")
	// String must not contain "b".
	// <nil>
	// "abc" contains "b"
	// String "abc" contains "b".
}

func ExampleStringHasPrefix() {
	vr := StringHasPrefix("a")
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("bac"))
	fmt.Println(LocalizeError(vr.Validate("bac"), "en"))
	// Output:
	// StringHasPrefix("a")
	// String must begin with "a".
	// <nil>
	// "bac" does not begin with "a"
	// String "bac" does not begin with "a".
}

func ExampleStringNotHasPrefix() {
	vr := StringNotHasPrefix("a")
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("bac"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(LocalizeError(vr.Validate("abc"), "en"))
	// Output:
	// StringNotHasPrefix("a")
	// String must not begin with "a".
	// <nil>
	// "abc" begins with "a"
	// String "abc" begins with "a".
}

func ExampleStringHasSuffix() {
	vr := StringHasSuffix("c")
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(vr.Validate("acb"))
	fmt.Println(LocalizeError(vr.Validate("acb"), "en"))
	// Output:
	// StringHasSuffix("c")
	// String must end with "c".
	// <nil>
	// "acb" does not end with "c"
	// String "acb" does not end with "c".
}

func ExampleStringNotHasSuffix() {
	vr := StringNotHasSuffix("c")
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate("acb"))
	fmt.Println(vr.Validate("abc"))
	fmt.Println(LocalizeError(vr.Validate("abc"), "en"))
	// Output:
	// StringNotHasSuffix("c")
	// String must not end with "c".
	// <nil>
	// "abc" ends with "c"
	// String "abc" ends with "c".
}
