package vld_test

import (
	"fmt"

	. "github.com/pierrre/vld"
)

func ExampleSliceLenEqual() {
	vr := SliceLenEqual[[]int](1)
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1}))
	fmt.Println(vr.Validate([]int{}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{}), "en"))
	// Output:
	// SliceLenEqual(1)
	// <nil>
	// length 0 is not equal to 1
	// Length 0 is not equal to 1.
}

func ExampleSliceLenMin() {
	vr := SliceLenMin[[]int](1)
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1}))
	fmt.Println(vr.Validate([]int{}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{}), "en"))
	// Output:
	// SliceLenMin(1)
	// <nil>
	// length 0 is less than 1
	// Length 0 is less than 1.
}

func ExampleSliceLenMax() {
	vr := SliceLenMax[[]int](1)
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{}))
	fmt.Println(vr.Validate([]int{1, 2}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{1, 2}), "en"))
	// Output:
	// SliceLenMax(1)
	// <nil>
	// length 2 is greater than 1
	// Length 2 is greater than 1.
}

func ExampleSliceLenRange() {
	vr := SliceLenRange[[]int](1, 2)
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1}))
	fmt.Println(vr.Validate([]int{}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{}), "en"))
	// Output:
	// SliceLenRange(1, 2)
	// <nil>
	// length 0 is not in the range [1, 2]
	// Length 0 is not in the range [1, 2].
}

func ExampleSliceEmpty() {
	vr := SliceEmpty[[]int]()
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{}))
	fmt.Println(vr.Validate([]int{1}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{1}), "en"))
	// Output:
	// SliceEmpty
	// <nil>
	// is not empty (1)
	// Value is not empty (1).
}

func ExampleSliceNotEmpty() {
	vr := SliceNotEmpty[[]int]()
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1}))
	fmt.Println(vr.Validate([]int{}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{}), "en"))
	// Output:
	// SliceNotEmpty
	// <nil>
	// is empty
	// Value is empty.
}

func ExampleSliceContains() {
	vr := SliceContains[[]int](1)
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1, 2}))
	fmt.Println(vr.Validate([]int{2, 3}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{2, 3}), "en"))
	// Output:
	// SliceContains(1)
	// <nil>
	// does not contain 1
	// Slice does not contain 1.
}

func ExampleSliceNotContains() {
	vr := SliceNotContains[[]int](1)
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{2, 3}))
	fmt.Println(vr.Validate([]int{1, 2}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{1, 2}), "en"))
	// Output:
	// SliceNotContains(1)
	// <nil>
	// contains 1
	// Slice contains 1.
}

func ExampleSliceEach() {
	vr := SliceEach[[]int](Equal(KeyValue[int, int]{Key: 0, Value: 1}))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1}))
	fmt.Println(vr.Validate([]int{2}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{2}), "en"))
	// Output:
	// SliceEach(Equal(vld.KeyValue[int,int]{Key:0, Value:1}))
	// <nil>
	// path index 0: vld.KeyValue[int,int]{Key:0, Value:2} is not equal to vld.KeyValue[int,int]{Key:0, Value:1}
	// Value vld.KeyValue[int,int]{Key:0, Value:2} is not equal to vld.KeyValue[int,int]{Key:0, Value:1}.
}

func ExampleSliceEachValue() {
	vr := SliceEachValue[[]int](Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1, 1}))
	fmt.Println(vr.Validate([]int{1, 2}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{1, 2}), "en"))
	// Output:
	// SliceEachValue(Equal(1))
	// <nil>
	// path index 1: 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleSliceUnique() {
	vr := SliceUnique[[]int]()
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1, 2}))
	fmt.Println(vr.Validate([]int{1, 2, 1}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{1, 2, 1}), "en"))
	// Output:
	// SliceUnique
	// <nil>
	// path index 2: duplicate 1 (index 0)
	// Duplicate 1 (index 0).
}

func ExampleSliceUniqueBy() {
	vr := SliceUniqueBy[[]int](func(v int) int { return v % 2 })
	fmt.Println(vr)
	fmt.Println(vr.Validate([]int{1, 2}))
	fmt.Println(vr.Validate([]int{1, 2, 3}))
	fmt.Println(GetErrorLocalizedMessage(vr.Validate([]int{1, 2, 3}), "en"))
	// Output:
	// SliceUniqueBy
	// <nil>
	// path index 2: duplicate 3 (index 0)
	// Duplicate 3 (index 0).
}
