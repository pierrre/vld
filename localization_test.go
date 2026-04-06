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

func TestLocalizeError(t *testing.T) {
	s := LocalizeError(nil, "en")
	assert.Zero(t, s)
	s = LocalizeError(errors.New("error"), "en")
	assert.Zero(t, s)
	s = LocalizeError(&EqualError[int]{Value: 1, Expected: 2}, "en")
	assertAutoString(t, s)
}
