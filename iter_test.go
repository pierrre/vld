package vld_test

import (
	"slices"
	"testing"

	. "github.com/pierrre/vld"
)

func TestSeqEach(t *testing.T) {
	testValidator(t,
		SeqEach(Equal(KeyValue[int, int]{Key: 0, Value: 1})),
		slices.Values([]int{1}),
		slices.Values([]int{2}),
	)
}

func TestSeqEachValue(t *testing.T) {
	testValidator(t,
		SeqEachValue(Equal(1)),
		slices.Values([]int{1}),
		slices.Values([]int{2}),
	)
}

func TestSeq2Each(t *testing.T) {
	testValidator(t,
		Seq2Each(Equal(KeyValue[int, int]{Key: 0, Value: 1})),
		slices.All([]int{1}),
		slices.All([]int{2}),
	)
}

func TestSeq2EachKey(t *testing.T) {
	testValidator(t,
		Seq2EachKey[int, int](Equal(0)),
		slices.All([]int{1}),
		slices.All([]int{1, 1}),
	)
}

func TestSeq2EachValue(t *testing.T) {
	testValidator(t,
		Seq2EachValue[int](Equal(1)),
		slices.All([]int{1}),
		slices.All([]int{2}),
	)
}
