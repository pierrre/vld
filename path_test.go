package vld_test

import (
	"errors"
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/assert/assertauto"
	. "github.com/pierrre/vld"
)

func TestPath(t *testing.T) {
	assertAutoString(t, Path(nil).String())
	p := Path{
		&FieldPathElem{Field: "User"},
		&IndexPathElem{Index: 0},
		&KeyPathElem{Key: "name"},
		&PointerPathElem{},
	}
	assertAutoString(t, p.String())
	for _, e := range p {
		assertAutoString(t, e.String())
	}
}

func TestPathError(t *testing.T) {
	e := &FieldPathElem{Field: "Test"}
	err := ErrorWrapPathElem(nil, e)
	assert.Zero(t, err)
	err = ErrorWrapPathElem(errors.New("error"), e)
	assertauto.Equal(t, err)
	p := GetErrorPath(err)
	assertAutoString(t, p.String())
}

func TestGetErrorPathJoined(t *testing.T) {
	err1 := ErrorWrapPathElem(errors.New("error 1"), &IndexPathElem{Index: 0})
	err2 := ErrorWrapPathElem(errors.New("error 2"), &IndexPathElem{Index: 1})
	err := ErrorJoin(err1, err2)
	p := GetErrorPath(err)
	assert.Equal(t, "[0]", p.String())
}
