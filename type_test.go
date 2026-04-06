package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestTypeOptional(t *testing.T) {
	testValidator[any](t, TypeOptional[any](Equal(1)), 1, "1", 2)
}

func TestTypeRequired(t *testing.T) {
	testValidator[any](t, TypeRequired[any](Equal(1)), 1, "1", 2)
}
