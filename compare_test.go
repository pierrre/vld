package vld_test

import (
	"cmp"
	"testing"

	. "github.com/pierrre/vld"
)

type intComparer int

func (i intComparer) Compare(other intComparer) int {
	return cmp.Compare(i, other)
}

func TestCmpEqual(t *testing.T) {
	testValidator(t, CmpEqual(intComparer(1)), intComparer(1), intComparer(2))
}

func TestCmpNotEqual(t *testing.T) {
	testValidator(t, CmpNotEqual(intComparer(1)), intComparer(2), intComparer(1))
}

func TestCmpMin(t *testing.T) {
	testValidator(t, CmpMin(intComparer(1)), intComparer(2), intComparer(1), intComparer(0))
}

func TestCmpMax(t *testing.T) {
	testValidator(t, CmpMax(intComparer(1)), intComparer(0), intComparer(1), intComparer(2))
}

func TestCmpRange(t *testing.T) {
	testValidator(t,
		CmpRange(intComparer(1), intComparer(3)),
		intComparer(2), intComparer(1), intComparer(3), intComparer(0), intComparer(4),
	)
}

func TestCmpLess(t *testing.T) {
	testValidator(t, CmpLess(intComparer(1)), intComparer(0), intComparer(1), intComparer(2))
}

func TestCmpGreater(t *testing.T) {
	testValidator(t, CmpGreater(intComparer(1)), intComparer(2), intComparer(1), intComparer(0))
}
