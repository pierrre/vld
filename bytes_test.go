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
	fmt.Println(GetErrorLocalization(vr.Validate([]byte("abd")), "en"))
	// Output:
	// BytesEqual([97 98 99])
	// <nil>
	// [97 98 100] is not equal to [97 98 99]
	// Bytes [97 98 100] is not equal to [97 98 99].
}

func ExampleBytesNotEqual() {
	vr := BytesNotEqual([]byte("abc"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abd")))
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(GetErrorLocalization(vr.Validate([]byte("abc")), "en"))
	// Output:
	// BytesNotEqual([97 98 99])
	// <nil>
	// [97 98 99] is equal to [97 98 99]
	// Bytes [97 98 99] is equal to [97 98 99].
}

func ExampleBytesContains() {
	vr := BytesContains([]byte("b"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(vr.Validate([]byte("acd")))
	fmt.Println(GetErrorLocalization(vr.Validate([]byte("acd")), "en"))
	// Output:
	// BytesContains([98])
	// <nil>
	// [97 99 100] does not contain [98]
	// Bytes [97 99 100] does not contain [98].
}

func ExampleBytesNotContains() {
	vr := BytesNotContains([]byte("b"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("acd")))
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(GetErrorLocalization(vr.Validate([]byte("abc")), "en"))
	// Output:
	// BytesNotContains([98])
	// <nil>
	// [97 98 99] contains [98]
	// Bytes [97 98 99] contains [98].
}

func ExampleBytesHasPrefix() {
	vr := BytesHasPrefix([]byte("a"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(vr.Validate([]byte("bac")))
	fmt.Println(GetErrorLocalization(vr.Validate([]byte("bac")), "en"))
	// Output:
	// BytesHasPrefix([97])
	// <nil>
	// [98 97 99] does not have prefix [97]
	// Bytes [98 97 99] does not have prefix [97].
}

func ExampleBytesNotHasPrefix() {
	vr := BytesNotHasPrefix([]byte("a"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("bac")))
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(GetErrorLocalization(vr.Validate([]byte("abc")), "en"))
	// Output:
	// BytesNotHasPrefix([97])
	// <nil>
	// [97 98 99] has prefix [97]
	// Bytes [97 98 99] has prefix [97].
}

func ExampleBytesHasSuffix() {
	vr := BytesHasSuffix([]byte("c"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(vr.Validate([]byte("acb")))
	fmt.Println(GetErrorLocalization(vr.Validate([]byte("acb")), "en"))
	// Output:
	// BytesHasSuffix([99])
	// <nil>
	// [97 99 98] does not have suffix [99]
	// Bytes [97 99 98] does not have suffix [99].
}

func ExampleBytesNotHasSuffix() {
	vr := BytesNotHasSuffix([]byte("c"))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]byte("acb")))
	fmt.Println(vr.Validate([]byte("abc")))
	fmt.Println(GetErrorLocalization(vr.Validate([]byte("abc")), "en"))
	// Output:
	// BytesNotHasSuffix([99])
	// <nil>
	// [97 98 99] has suffix [99]
	// Bytes [97 98 99] has suffix [99].
}
