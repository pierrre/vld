package vld_test

import (
	"cmp"
	"testing"

	. "github.com/pierrre/vld"
)

func TestMin(t *testing.T) {
	testValidator(t, Min(1), 2, 1, 0)
}

func TestMinCmpFunc(t *testing.T) {
	testValidator(t, MinCmpFunc(1, cmp.Compare), 2, 1, 0)
}

func TestMax(t *testing.T) {
	testValidator(t, Max(1), 0, 1, 2)
}

func TestMaxCmpFunc(t *testing.T) {
	testValidator(t, MaxCmpFunc(1, cmp.Compare), 0, 1, 2)
}

func TestRange(t *testing.T) {
	testValidator(t, Range(1, 3), 2, 1, 3, 0, 4)
}

func TestRangeCmpFunc(t *testing.T) {
	testValidator(t, RangeCmpFunc(1, 3, cmp.Compare), 2, 1, 3, 0, 4)
}

func TestLess(t *testing.T) {
	testValidator(t, Less(1), 0, 1, 2)
}

func TestLessCmpFunc(t *testing.T) {
	testValidator(t, LessCmpFunc(1, cmp.Compare), 0, 1, 2)
}

func TestGreater(t *testing.T) {
	testValidator(t, Greater(1), 2, 1, 0)
}

func TestGreaterCmpFunc(t *testing.T) {
	testValidator(t, GreaterCmpFunc(1, cmp.Compare), 2, 1, 0)
}
