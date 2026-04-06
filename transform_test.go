package vld_test

import (
	"strconv"
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/vld"
)

func TestGet(t *testing.T) {
	type User struct {
		Name string
	}
	testValidator(t,
		Get(func(u User) string { return u.Name }, StringLenMax(5)),
		User{Name: "Alice"},
		User{Name: "Charlie"},
	)
}

func TestParse(t *testing.T) {
	vr := Parse(strconv.Atoi, Equal(1))
	testValidator(t,
		vr,
		"1", "2", "a",
	)
	err := vr.Validate("a")
	perr, _ := assert.ErrorAsType[*ParseError[string, int]](t, err)
	err = perr.Unwrap()
	assert.Error(t, err)
}

func TestWrap(t *testing.T) {
	testValidator(t, Wrap("message", Equal(1)), 1, 2)
}

func TestField(t *testing.T) {
	type User struct {
		Name string
	}
	testValidator(t,
		Field("Name", func(u User) string { return u.Name }, StringLenMax(5)),
		User{Name: "Alice"},
		User{Name: "Charlie"},
	)
}

func TestMessage(t *testing.T) {
	testValidator(t, Message("message", Equal(1)), 1, 2)
}
