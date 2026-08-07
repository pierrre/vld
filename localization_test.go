package vld_test

import (
	"errors"
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/vld"
)

func TestLocalize(t *testing.T) {
	s := Localize("EqualError", []any{1, 2}, "fr")
	assertAutoString(t, s)
	s = Localize("_Unknown", nil, "en")
	assert.Zero(t, s)
	s = Localize("EqualError", []any{1, 2}, "unknown")
	assert.Zero(t, s)
}

func TestLocalizeNoMutateArgs(t *testing.T) {
	lv := &EqualError[int]{Value: 1, Expected: 2}
	args := []any{lv, 2}
	_ = Localize("EqualError", args, "en")
	got, ok := args[0].(*EqualError[int])
	assert.True(t, ok)
	assert.Equal(t, lv, got)
}

func TestLocalizeError(t *testing.T) {
	s := LocalizeError(nil, "en")
	assert.Zero(t, s)
	s = LocalizeError(errors.New("error"), "en")
	assert.Zero(t, s)
	s = LocalizeError(&EqualError[int]{Value: 1, Expected: 2}, "en")
	assertAutoString(t, s)
}
