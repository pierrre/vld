package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestPointerOptional(t *testing.T) {
	testValidator(t, PointerOptional(Equal(1)), new(1), nil, new(2))
}

func TestPointerRequired(t *testing.T) {
	testValidator(t, PointerRequired(Equal(1)), new(1), nil, new(2))
}
