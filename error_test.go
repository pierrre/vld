package vld_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/assert/assertauto"
	. "github.com/pierrre/vld"
)

func TestErrorWrap(t *testing.T) {
	wrapFunc := func(err error) error {
		return fmt.Errorf("test: %w", err)
	}
	assert.Zero(t, ErrorWrap(nil, wrapFunc))
	assertauto.Equal(t, ErrorWrap(errors.New("error"), wrapFunc))
	assertauto.Equal(t, ErrorWrap(errors.Join(errors.New("error 1"), errors.New("error 2")), wrapFunc))
}

func TestErrorWrapMessage(t *testing.T) {
	assert.Zero(t, ErrorWrapMessage(nil, "test"))
	assertauto.Equal(t, ErrorWrapMessage(errors.New("error"), "test"))
}

func TestErrorWrapMessagef(t *testing.T) {
	assert.Zero(t, ErrorWrapMessagef(nil, "test %d", 1))
	assertauto.Equal(t, ErrorWrapMessagef(errors.New("error"), "test %d", 1))
}

func TestErrorJoin(t *testing.T) {
	assert.Zero(t, ErrorJoin())
	assert.Zero(t, ErrorJoin(nil))
	assertauto.Equal(t, ErrorJoin(errors.Join(errors.New("error 1"), errors.New("error 2")), errors.New("error 3")))
}

func TestGetErrors(t *testing.T) {
	assert.SliceEmpty(t, GetErrors(nil))
	assertauto.Equal(t, GetErrors(errors.New("error")))
	assertauto.Equal(t, GetErrors(errors.Join(errors.New("error1"), errors.New("error2"))))
}
