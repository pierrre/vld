package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestZero(t *testing.T) {
	testValidator(t, Zero[int](), 0, 1)
}

func TestNotZero(t *testing.T) {
	testValidator(t, NotZero[int](), 1, 0)
}

func TestOptional(t *testing.T) {
	testValidator(t, Optional(Equal(1)), 0, 1, 2)
}

func TestRequired(t *testing.T) {
	testValidator(t, Required(Equal(1)), 1, 0, 2)
}
