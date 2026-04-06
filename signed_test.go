package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestPositive(t *testing.T) {
	testValidator(t, Positive[int](), 1, 0, -1)
}

func TestNegative(t *testing.T) {
	testValidator(t, Negative[int](), -1, 0, 1)
}
