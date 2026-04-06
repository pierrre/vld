package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestRegexpMatch(t *testing.T) {
	testValidator(t, RegexpMatch(`^[a-z]+$`), "abc", "abc123")
}

func TestRegexpNotMatch(t *testing.T) {
	testValidator(t, RegexpNotMatch(`^[a-z]+$`), "abc123", "abc")
}
