package vld_test

import (
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/vld"
)

func TestIf(t *testing.T) {
	testValidator(t,
		If(func(v int) bool { return v != 0 }, Min(1)),
		0, 5, -1,
	)
}

func TestIfElse(t *testing.T) {
	testValidator(t,
		IfElse(func(v int) bool { return v > 0 }, Max(10), Min(-10)),
		5, 11, -5, -11,
	)
}

func TestSwitch(t *testing.T) {
	testValidator(t,
		Switch(
			Case(func(v int) bool { return v > 0 }, Max(10)),
			Case(func(v int) bool { return v < 0 }, Min(-10)),
		),
		5, -5, 0, 11, -11,
	)
}

func TestSwitchNoMatch(t *testing.T) {
	sv := Switch(
		Case(func(v int) bool { return v > 0 }, Max(10)),
		Case(func(v int) bool { return v < 0 }, Min(-10)),
	)
	assert.Zero(t, sv.Validate(0))
}
