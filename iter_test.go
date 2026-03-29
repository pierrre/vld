package vld_test

import (
	"fmt"
	"slices"

	. "github.com/pierrre/vld"
)

func ExampleSeqEach() {
	vr := SeqEach(Equal(KeyValue[int, int]{Key: 0, Value: 1}))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.Values([]int{1})))
	fmt.Println(vr.Validate(slices.Values([]int{2})))
	fmt.Println(LocalizeError(vr.Validate(slices.Values([]int{2})), "en"))
	// Output:
	// SeqEach(Equal(vld.KeyValue[int,int]{Key:0, Value:1}))
	// <nil>
	// path index 0: vld.KeyValue[int,int]{Key:0, Value:2} is not equal to vld.KeyValue[int,int]{Key:0, Value:1}
	// Value vld.KeyValue[int,int]{Key:0, Value:2} is not equal to vld.KeyValue[int,int]{Key:0, Value:1}.
}

func ExampleSeqEachValue() {
	vr := SeqEachValue(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.Values([]int{1})))
	fmt.Println(vr.Validate(slices.Values([]int{2})))
	fmt.Println(LocalizeError(vr.Validate(slices.Values([]int{2})), "en"))
	// Output:
	// SeqEachValue(Equal(1))
	// <nil>
	// path index 0: 2 is not equal to 1
	// Value 2 is not equal to 1.
}

func ExampleSeq2Each() {
	vr := Seq2Each(Equal(KeyValue[int, int]{Key: 0, Value: 1}))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.All([]int{1})))
	fmt.Println(vr.Validate(slices.All([]int{2})))
	fmt.Println(LocalizeError(vr.Validate(slices.All([]int{2})), "en"))
	// Output:
	// Seq2Each(Equal(vld.KeyValue[int,int]{Key:0, Value:1}))
	// <nil>
	// path index 0: vld.KeyValue[int,int]{Key:0, Value:2} is not equal to vld.KeyValue[int,int]{Key:0, Value:1}
	// Value vld.KeyValue[int,int]{Key:0, Value:2} is not equal to vld.KeyValue[int,int]{Key:0, Value:1}.
}

func ExampleSeq2EachKey() {
	vr := Seq2EachKey[int, int](Equal(0))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.All([]int{1})))
	fmt.Println(vr.Validate(slices.All([]int{1, 1})))
	fmt.Println(LocalizeError(vr.Validate(slices.All([]int{1, 1})), "en"))
	// Output:
	// Seq2EachKey(Equal(0))
	// <nil>
	// path index 1: path field "key": 1 is not equal to 0
	// Value 1 is not equal to 0.
}

func ExampleSeq2EachValue() {
	vr := Seq2EachValue[int](Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.All([]int{1})))
	fmt.Println(vr.Validate(slices.All([]int{2})))
	fmt.Println(LocalizeError(vr.Validate(slices.All([]int{2})), "en"))
	// Output:
	// Seq2EachValue(Equal(1))
	// <nil>
	// path index 0: path field "value": 2 is not equal to 1
	// Value 2 is not equal to 1.
}
