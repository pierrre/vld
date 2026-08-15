package vld_test

import (
	"testing"

	. "github.com/pierrre/vld"
)

func TestBytesEqual(t *testing.T) {
	testValidator(t, BytesEqual([]byte("abc")), []byte("abc"), []byte("abd"))
}

func TestBytesNotEqual(t *testing.T) {
	testValidator(t, BytesNotEqual([]byte("abc")), []byte("abd"), []byte("abc"))
}

func TestBytesContains(t *testing.T) {
	testValidator(t, BytesContains([]byte("b")), []byte("abc"), []byte("acd"))
}

func TestBytesNotContains(t *testing.T) {
	testValidator(t, BytesNotContains([]byte("b")), []byte("acd"), []byte("abc"))
}

func TestBytesHasPrefix(t *testing.T) {
	testValidator(t, BytesHasPrefix([]byte("a")), []byte("abc"), []byte("bac"))
}

func TestBytesNotHasPrefix(t *testing.T) {
	testValidator(t, BytesNotHasPrefix([]byte("a")), []byte("bac"), []byte("abc"))
}

func TestBytesHasSuffix(t *testing.T) {
	testValidator(t, BytesHasSuffix([]byte("c")), []byte("abc"), []byte("acb"))
}

func TestBytesNotHasSuffix(t *testing.T) {
	testValidator(t, BytesNotHasSuffix([]byte("c")), []byte("acb"), []byte("abc"))
}

func TestBytesUTF8(t *testing.T) {
	testValidator(t, BytesUTF8(), []byte("abc"), []byte("abc\xff"))
}
