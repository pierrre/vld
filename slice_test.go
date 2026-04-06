package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestSliceLenEqual(t *testing.T) {
	testValidator(t, SliceLenEqual[[]int](1), []int{1}, []int{})
}

func TestSliceLenMin(t *testing.T) {
	testValidator(t, SliceLenMin[[]int](1), []int{1}, []int{})
}

func TestSliceLenMax(t *testing.T) {
	testValidator(t, SliceLenMax[[]int](1), []int{}, []int{1, 2})
}

func TestSliceLenRange(t *testing.T) {
	testValidator(t, SliceLenRange[[]int](1, 2), []int{1}, []int{})
}

func TestSliceEmpty(t *testing.T) {
	testValidator(t, SliceEmpty[[]int](), []int{}, []int{1})
}

func TestSliceNotEmpty(t *testing.T) {
	testValidator(t, SliceNotEmpty[[]int](), []int{1}, []int{})
}

func TestSliceContains(t *testing.T) {
	testValidator(t, SliceContains[[]int](1), []int{1, 2}, []int{2, 3})
}

func TestSliceNotContains(t *testing.T) {
	testValidator(t, SliceNotContains[[]int](1), []int{2, 3}, []int{1, 2})
}

func TestSliceEach(t *testing.T) {
	testValidator(t,
		SliceEach[[]int](Equal(KeyValue[int, int]{Key: 0, Value: 1})),
		[]int{1},
		[]int{2},
	)
}

func TestSliceEachValue(t *testing.T) {
	testValidator(t,
		SliceEachValue[[]int](Equal(1)),
		[]int{1, 1},
		[]int{1, 2},
	)
}

func TestSliceUnique(t *testing.T) {
	testValidator(t, SliceUnique[[]int](), []int{1, 2}, []int{1, 2, 1})
}

func TestSliceUniqueBy(t *testing.T) {
	testValidator(t,
		SliceUniqueBy[[]int](func(v int) int { return v % 2 }),
		[]int{1, 2},
		[]int{1, 2, 3},
	)
}
