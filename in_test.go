package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestIn(t *testing.T) {
	testValidator(t, In(1, 2, 3), 1, 4)
}

func TestNotIn(t *testing.T) {
	testValidator(t, NotIn(1, 2, 3), 4, 1)
}
