package vld_test

import (
	"net/url"
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/vld"
)

func TestURL(t *testing.T) {
	vr := URL(nil, nil)
	testValidator(t, vr, "https://example.com/path?query=value#fragment", "http://[::1")
	err := vr.Validate("http://[::1")
	uerr, _ := assert.ErrorAsType[*URLError](t, err)
	err = uerr.Unwrap()
	assert.Error(t, err)
}

func TestURLWithParser(t *testing.T) {
	testValidator(t, URL(url.ParseRequestURI, nil), "https://example.com/path", "http://[::1")
}

func TestURLWithValidator(t *testing.T) {
	testValidator(t, URL(nil, NotZero[url.URL]()), "https://example.com/path", "", "http://[::1")
}

func TestURLWithParserAndValidator(t *testing.T) {
	testValidator(t, URL(url.ParseRequestURI, NotZero[url.URL]()), "https://example.com/path", "", "http://[::1")
}
