package vld_test

import (
	"cmp"
	"testing"

	. "github.com/pierrre/vld"
)

func TestEqual(t *testing.T) {
	testValidator(t, Equal(1), 1, 2)
}

func TestEqualFunc(t *testing.T) {
	testValidator(t, EqualFunc(1, func(a, b int) bool { return a == b }), 1, 2)
}

func TestEqualCmpFunc(t *testing.T) {
	testValidator(t, EqualCmpFunc(1, cmp.Compare), 1, 2)
}

func TestNotEqual(t *testing.T) {
	testValidator(t, NotEqual(1), 2, 1)
}

func TestNotEqualFunc(t *testing.T) {
	testValidator(t, NotEqualFunc(1, func(a, b int) bool { return a == b }), 2, 1)
}

func TestNotEqualCmpFunc(t *testing.T) {
	testValidator(t, NotEqualCmpFunc(1, cmp.Compare), 2, 1)
}
