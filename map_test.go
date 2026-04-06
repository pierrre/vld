package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestMapLenEqual(t *testing.T) {
	testValidator(t, MapLenEqual[map[string]int](1), map[string]int{"a": 1}, map[string]int{})
}

func TestMapLenMin(t *testing.T) {
	testValidator(t, MapLenMin[map[string]int](1), map[string]int{"a": 1}, map[string]int{})
}

func TestMapLenMax(t *testing.T) {
	testValidator(t, MapLenMax[map[string]int](1), map[string]int{}, map[string]int{"a": 1, "b": 2})
}

func TestMapLenRange(t *testing.T) {
	testValidator(t, MapLenRange[map[string]int](1, 2), map[string]int{"a": 1}, map[string]int{})
}

func TestMapEmpty(t *testing.T) {
	testValidator(t, MapEmpty[map[string]int](), map[string]int{}, map[string]int{"a": 1})
}

func TestMapNotEmpty(t *testing.T) {
	testValidator(t, MapNotEmpty[map[string]int](), map[string]int{"a": 1}, map[string]int{})
}

func TestMapEach(t *testing.T) {
	testValidator(t,
		MapEach[map[string]int](Equal(KeyValue[string, int]{Key: "a", Value: 1})),
		map[string]int{"a": 1},
		map[string]int{"b": 2},
	)
}

func TestMapEachKey(t *testing.T) {
	testValidator(t,
		MapEachKey[map[string]int](Equal("a")),
		map[string]int{"a": 1},
		map[string]int{"b": 2},
	)
}

func TestMapEachValue(t *testing.T) {
	testValidator(t,
		MapEachValue[map[string]int](Equal(1)),
		map[string]int{"a": 1},
		map[string]int{"b": 2},
	)
}

func TestMapSortedEach(t *testing.T) {
	testValidator(t,
		MapSortedEach[map[string]int](Equal(KeyValue[string, int]{Key: "a", Value: 1})),
		map[string]int{"a": 1},
		map[string]int{"b": 2},
	)
}

func TestMapSortedEachKey(t *testing.T) {
	testValidator(t,
		MapSortedEachKey[map[string]int](Equal("a")),
		map[string]int{"a": 1},
		map[string]int{"b": 2},
	)
}

func TestMapSortedEachValue(t *testing.T) {
	testValidator(t,
		MapSortedEachValue[map[string]int](Equal(1)),
		map[string]int{"a": 1},
		map[string]int{"b": 2},
	)
}
