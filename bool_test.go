package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestAnd(t *testing.T) {
	testValidator(t, And(Min(1), Max(10)), 5, 0)
}

func TestOr(t *testing.T) {
	testValidator(t, Or(Min(10), Max(1)), 11, 0, 5)
}

func TestOrEmpty(t *testing.T) {
	testValidator(t, Or[int](), 1)
}

func TestAll(t *testing.T) {
	testValidator(t, All(Min(1), Max(10)), 5, 0)
}
