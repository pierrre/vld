package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleMapLenEqual() {
	vr := MapLenEqual[map[string]int](1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{}))
	// Output:
	// MapLenEqual(1)
	// <nil>
	// length 0 is not equal to 1
}

func ExampleMapLenMin() {
	vr := MapLenMin[map[string]int](1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{}))
	// Output:
	// MapLenMin(1)
	// <nil>
	// length 0 is less than 1
}

func ExampleMapLenMax() {
	vr := MapLenMax[map[string]int](1)
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{}))
	fmt.Println(vr.Validate(map[string]int{"a": 1, "b": 2}))
	// Output:
	// MapLenMax(1)
	// <nil>
	// length 2 is greater than 1
}

func ExampleMapLenRange() {
	vr := MapLenRange[map[string]int](1, 2)
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{}))
	// Output:
	// MapLenRange(1, 2)
	// <nil>
	// length 0 is not in the range [1, 2]
}

func ExampleMapEmpty() {
	vr := MapEmpty[map[string]int]()
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{}))
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	// Output:
	// MapEmpty
	// <nil>
	// is not empty (1)
}

func ExampleMapNotEmpty() {
	vr := MapNotEmpty[map[string]int]()
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{}))
	// Output:
	// MapNotEmpty
	// <nil>
	// is empty
}

func ExampleMapEach() {
	vr := MapEach[map[string]int](Equal(KeyValue[string, int]{Key: "a", Value: 1}))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	// Output:
	// MapEach(Equal(vld.KeyValue[string,int]{Key:"a", Value:1}))
	// <nil>
	// "b": vld.KeyValue[string,int]{Key:"b", Value:2} is not equal to vld.KeyValue[string,int]{Key:"a", Value:1}
}

func ExampleMapEachKey() {
	vr := MapEachKey[map[string]int](Equal("a"))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	// Output:
	// MapEachKey(Equal("a"))
	// <nil>
	// "b": "b" is not equal to "a"
}

func ExampleMapEachValue() {
	vr := MapEachValue[map[string]int](Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	// Output:
	// MapEachValue(Equal(1))
	// <nil>
	// "b": 2 is not equal to 1
}
