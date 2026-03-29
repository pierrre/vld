package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleBytesEqual() {
	vr := BytesEqual([]byte("abc"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(vr.Validate([]byte("abd")))
	fmt.Println(LocalizeError(vr.Validate([]byte("abd")), "en"))
	// Output:
	// BytesEqual("abc")
	// <nil>
	// "abd" is not equal to "abc"
	// Bytes "abd" is not equal to "abc".
}

func ExampleBytesNotEqual() {
	vr := BytesNotEqual([]byte("abc"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abd")))
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(LocalizeError(vr.Validate([]byte("abc")), "en"))
	// Output:
	// BytesNotEqual("abc")
	// <nil>
	// "abc" is equal to "abc"
	// Bytes "abc" is equal to "abc".
}

func ExampleBytesContains() {
	vr := BytesContains([]byte("b"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(vr.Validate([]byte("acd")))
	fmt.Println(LocalizeError(vr.Validate([]byte("acd")), "en"))
	// Output:
	// BytesContains("b")
	// <nil>
	// "acd" does not contain "b"
	// Bytes "acd" does not contain "b".
}

func ExampleBytesNotContains() {
	vr := BytesNotContains([]byte("b"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("acd")))
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(LocalizeError(vr.Validate([]byte("abc")), "en"))
	// Output:
	// BytesNotContains("b")
	// <nil>
	// "abc" contains "b"
	// Bytes "abc" contains "b".
}

func ExampleBytesHasPrefix() {
	vr := BytesHasPrefix([]byte("a"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(vr.Validate([]byte("bac")))
	fmt.Println(LocalizeError(vr.Validate([]byte("bac")), "en"))
	// Output:
	// BytesHasPrefix("a")
	// <nil>
	// "bac" does not have prefix "a"
	// Bytes "bac" does not have prefix "a".
}

func ExampleBytesNotHasPrefix() {
	vr := BytesNotHasPrefix([]byte("a"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("bac")))
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(LocalizeError(vr.Validate([]byte("abc")), "en"))
	// Output:
	// BytesNotHasPrefix("a")
	// <nil>
	// "abc" has prefix "a"
	// Bytes "abc" has prefix "a".
}

func ExampleBytesHasSuffix() {
	vr := BytesHasSuffix([]byte("c"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(vr.Validate([]byte("acb")))
	fmt.Println(LocalizeError(vr.Validate([]byte("acb")), "en"))
	// Output:
	// BytesHasSuffix("c")
	// <nil>
	// "acb" does not have suffix "c"
	// Bytes "acb" does not have suffix "c".
}

func ExampleBytesNotHasSuffix() {
	vr := BytesNotHasSuffix([]byte("c"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("acb")))
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(LocalizeError(vr.Validate([]byte("abc")), "en"))
	// Output:
	// BytesNotHasSuffix("c")
	// <nil>
	// "abc" has suffix "c"
	// Bytes "abc" has suffix "c".
}
