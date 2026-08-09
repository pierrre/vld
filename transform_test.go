package vld_test

import (
	"errors"
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

func TestMessageErrorUnwrap(t *testing.T) {
	vr := Message("override", Equal(1))
	err := vr.Validate(2)
	msgErr, _ := assert.ErrorAsType[*MessageError](t, err)
	assert.Equal(t, "override", msgErr.Message)
	eqErr, _ := assert.ErrorAsType[*EqualError[int]](t, err)
	assert.Equal(t, eqErr.Value, 2)
	assert.Equal(t, eqErr.Expected, 1)
	assert.ErrorEqual(t, err, "override")
}

func TestMessageErrorIs(t *testing.T) {
	errSentinel := errors.New("sentinel")
	vr := Message("override", ValidatorFunc[int](func(v int) error {
		return errSentinel
	}))
	err := vr.Validate(1)
	assert.ErrorIs(t, err, errSentinel)
}

func TestMessageErrorPath(t *testing.T) {
	type User struct {
		Name string
	}
	vr := Message("override", Field("Name", func(u User) string { return u.Name }, StringLenMax(5)))
	err := vr.Validate(User{Name: "Charlie"})
	path := GetErrorPath(err)
	assert.Equal(t, ".Name", path.String())
}

func TestMessageErrorLocalize(t *testing.T) {
	vr := Message("override", Equal(1))
	err := vr.Validate(2)
	assert.Equal(t, "override", err.Error())
	assert.Equal(t, "override", LocalizeError(err, "en"))
	assert.Equal(t, "override", LocalizeError(err, "fr"))
	assert.Equal(t, err.Error(), LocalizeError(err, "en"))
}
