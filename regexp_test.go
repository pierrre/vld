package vld_test

import (
	"regexp"
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/vld"
)

func TestRegexpMatch(t *testing.T) {
	testValidator(t, RegexpMatch(regexp.MustCompile(`^[a-z]+$`)), "abc", "abc123")
}

func TestRegexpMatchString(t *testing.T) {
	testValidator(t, RegexpMatch(`^[a-z]+$`), "abc", "abc123")
}

func TestRegexpMatchPanicNil(t *testing.T) {
	assert.Panics(t, func() {
		_ = RegexpMatch[*regexp.Regexp](nil)
	})
}

func TestRegexpNotMatch(t *testing.T) {
	testValidator(t, RegexpNotMatch(regexp.MustCompile(`^[a-z]+$`)), "abc123", "abc")
}

func TestRegexpNotMatchString(t *testing.T) {
	testValidator(t, RegexpNotMatch(`^[a-z]+$`), "abc123", "abc")
}

func TestRegexpNotMatchPanicNil(t *testing.T) {
	assert.Panics(t, func() {
		_ = RegexpNotMatch[*regexp.Regexp](nil)
	})
}
