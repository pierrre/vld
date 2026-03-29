package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleMapLenEqual() {
	vr := MapLenEqual[map[string]int](1)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{}), "en"))
	// Output:
	// MapLenEqual(1)
	// Length must be equal to 1.
	// <nil>
	// length 0 is not equal to 1
	// Length 0 is not equal to 1.
}

func ExampleMapLenMin() {
	vr := MapLenMin[map[string]int](1)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{}), "en"))
	// Output:
	// MapLenMin(1)
	// Length must be greater than or equal to 1.
	// <nil>
	// length 0 is less than 1
	// Length 0 is less than 1.
}

func ExampleMapLenMax() {
	vr := MapLenMax[map[string]int](1)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(map[string]int{}))
	fmt.Println(vr.Validate(map[string]int{"a": 1, "b": 2}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{"a": 1, "b": 2}), "en"))
	// Output:
	// MapLenMax(1)
	// Length must be less than or equal to 1.
	// <nil>
	// length 2 is greater than 1
	// Length 2 is greater than 1.
}

func ExampleMapLenRange() {
	vr := MapLenRange[map[string]int](1, 2)
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{}), "en"))
	// Output:
	// MapLenRange(1, 2)
	// Length must be in the range [1, 2].
	// <nil>
	// length 0 is not in the range [1, 2]
	// Length 0 is not in the range [1, 2].
}

func ExampleMapEmpty() {
	vr := MapEmpty[map[string]int]()
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(map[string]int{}))
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{"a": 1}), "en"))
	// Output:
	// MapEmpty
	// Value must be empty.
	// <nil>
	// is not empty (1)
	// Value is not empty (1).
}

func ExampleMapNotEmpty() {
	vr := MapNotEmpty[map[string]int]()
	fmt.Println(vr)
	fmt.Println(LocalizeValidator(vr, "en"))
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{}), "en"))
	// Output:
	// MapNotEmpty
	// Value must not be empty.
	// <nil>
	// is empty
	// Value is empty.
}

func ExampleMapEach() {
	vr := MapEach[map[string]int](Equal(KeyValue[string, int]{Key: "a", Value: 1}))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{"b": 2}), "en"))
	// Output:
	// MapEach(Equal(vld.KeyValue[string,int]{Key:"a", Value:1}))
	// <nil>
	// path key "b": vld.KeyValue[string,int]{Key:"b", Value:2} is not equal to vld.KeyValue[string,int]{Key:"a", Value:1}
	// Value vld.KeyValue[string,int]{Key:"b", Value:2} is not equal to vld.KeyValue[string,int]{Key:"a", Value:1}.
}

func ExampleMapEachKey() {
	vr := MapEachKey[map[string]int](Equal("a"))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{"b": 2}), "en"))
	// Output:
	// MapEachKey(Equal("a"))
	// <nil>
	// path key "b": path field "key": "b" is not equal to "a"
	// Value "b" is not equal to "a".
}

func ExampleMapEachValue() {
	vr := MapEachValue[map[string]int](Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{"b": 2}), "en"))
	// Output:
	// MapEachValue(Equal(1))
	// <nil>
	// path key "b": path field "value": 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleMapSortedEach() {
	vr := MapSortedEach[map[string]int](Equal(KeyValue[string, int]{Key: "a", Value: 1}))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{"b": 2}), "en"))
	// Output:
	// MapSortedEach(Equal(vld.KeyValue[string,int]{Key:"a", Value:1}))
	// <nil>
	// path key "b": vld.KeyValue[string,int]{Key:"b", Value:2} is not equal to vld.KeyValue[string,int]{Key:"a", Value:1}
	// Value vld.KeyValue[string,int]{Key:"b", Value:2} is not equal to vld.KeyValue[string,int]{Key:"a", Value:1}.
}

func ExampleMapSortedEachKey() {
	vr := MapSortedEachKey[map[string]int](Equal("a"))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{"b": 2}), "en"))
	// Output:
	// MapSortedEachKey(Equal("a"))
	// <nil>
	// path key "b": path field "key": "b" is not equal to "a"
	// Value "b" is not equal to "a".
}

func ExampleMapSortedEachValue() {
	vr := MapSortedEachValue[map[string]int](Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(map[string]int{"a": 1}))
	fmt.Println(vr.Validate(map[string]int{"b": 2}))
	fmt.Println(LocalizeError(vr.Validate(map[string]int{"b": 2}), "en"))
	// Output:
	// MapSortedEachValue(Equal(1))
	// <nil>
	// path key "b": path field "value": 2 is not equal to 1
	// Value 2 is not equal to 1.
}
