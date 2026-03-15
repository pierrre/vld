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
	// Output:
	// SeqEach(Equal(vld.KeyValue[int,int]{Key:0, Value:1}))
	// <nil>
	// 0: vld.KeyValue[int,int]{Key:0, Value:2} is not equal to vld.KeyValue[int,int]{Key:0, Value:1}
}

func ExampleSeqEachValue() {
	vr := SeqEachValue(Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.Values([]int{1})))
	fmt.Println(vr.Validate(slices.Values([]int{2})))
	// Output:
	// SeqEachValue(Equal(1))
	// <nil>
	// 0: 2 is not equal to 1
}

func ExampleSeq2Each() {
	vr := Seq2Each(Equal(KeyValue[int, int]{Key: 0, Value: 1}))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.All([]int{1})))
	fmt.Println(vr.Validate(slices.All([]int{2})))
	// Output:
	// Seq2Each(Equal(vld.KeyValue[int,int]{Key:0, Value:1}))
	// <nil>
	// 0: vld.KeyValue[int,int]{Key:0, Value:2} is not equal to vld.KeyValue[int,int]{Key:0, Value:1}
}

func ExampleSeq2EachKey() {
	vr := Seq2EachKey[int, int](Equal(0))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.All([]int{1})))
	fmt.Println(vr.Validate(slices.All([]int{1, 1})))
	// Output:
	// Seq2EachKey(Equal(0))
	// <nil>
	// 1: 1 is not equal to 0
}

func ExampleSeq2EachValue() {
	vr := Seq2EachValue[int, int](Equal(1))
	fmt.Println(vr)
	fmt.Println(vr.Validate(slices.All([]int{1})))
	fmt.Println(vr.Validate(slices.All([]int{2})))
	// Output:
	// Seq2EachValue(Equal(1))
	// <nil>
	// 0: 2 is not equal to 1
}
