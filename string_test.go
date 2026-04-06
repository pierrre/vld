package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestStringLenEqual(t *testing.T) {
	testValidator(t, StringLenEqual(3), "abc", "abcd")
}

func TestStringLenMin(t *testing.T) {
	testValidator(t, StringLenMin(3), "abc", "ab")
}

func TestStringLenMax(t *testing.T) {
	testValidator(t, StringLenMax(3), "abc", "abcd")
}

func TestStringLenRange(t *testing.T) {
	testValidator(t, StringLenRange(2, 4), "abc", "a", "abcde")
}

func TestStringEmpty(t *testing.T) {
	testValidator(t, StringEmpty(), "", "a")
}

func TestStringNotEmpty(t *testing.T) {
	testValidator(t, StringNotEmpty(), "a", "")
}

func TestStringRunesEqual(t *testing.T) {
	testValidator(t, StringRunesEqual(3), "abc", "abcd")
}

func TestStringRunesMin(t *testing.T) {
	testValidator(t, StringRunesMin(3), "abc", "ab")
}

func TestStringRunesMax(t *testing.T) {
	testValidator(t, StringRunesMax(3), "abc", "abcd")
}

func TestStringRunesRange(t *testing.T) {
	testValidator(t, StringRunesRange(2, 4), "abc", "a", "abcde")
}

func TestStringContains(t *testing.T) {
	testValidator(t, StringContains("b"), "abc", "acd")
}

func TestStringNotContains(t *testing.T) {
	testValidator(t, StringNotContains("b"), "acd", "abc")
}

func TestStringHasPrefix(t *testing.T) {
	testValidator(t, StringHasPrefix("a"), "abc", "bac")
}

func TestStringNotHasPrefix(t *testing.T) {
	testValidator(t, StringNotHasPrefix("a"), "bac", "abc")
}

func TestStringHasSuffix(t *testing.T) {
	testValidator(t, StringHasSuffix("c"), "abc", "acb")
}

func TestStringNotHasSuffix(t *testing.T) {
	testValidator(t, StringNotHasSuffix("c"), "acb", "abc")
}
