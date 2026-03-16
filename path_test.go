package vld_test

import (
	"errors"
	"fmt"

	. "github.com/pierrre/vld"
)

func ExamplePath() {
	fmt.Println(Path(nil))
	fmt.Println(Path{
		&FieldPathElem{Field: "User"},
		&IndexPathElem{Index: 0},
		&KeyPathElem{Key: "name"},
		&PointerPathElem{},
	})
	// Output:
	// .
	// .User[0]["name"]*
}

func ExampleFieldPathElem() {
	e := &FieldPathElem{Field: "Test"}
	fmt.Println(e.PathElem())
	fmt.Println(e.String())
	// Output:
	// .Test
	// field "Test"
}

func ExampleIndexPathElem() {
	e := &IndexPathElem{Index: 42}
	fmt.Println(e.PathElem())
	fmt.Println(e.String())
	// Output:
	// [42]
	// index 42
}

func ExampleKeyPathElem() {
	e := &KeyPathElem{Key: "Test"}
	fmt.Println(e.PathElem())
	fmt.Println(e.String())
	// Output:
	// ["Test"]
	// key "Test"
}

func ExamplePointerPathElem() {
	e := &PointerPathElem{}
	fmt.Println(e.PathElem())
	fmt.Println(e.String())
	// Output:
	// *
	// pointer
}

func ExamplePathElemError() {
	e := &FieldPathElem{Field: "Test"}
	err := &PathElemError{
		Err:      errors.New("error"),
		PathElem: e,
	}
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
	// Output:
	// path field "Test": error
	// error
}

func ExampleErrorWrapPathElement() {
	e := &FieldPathElem{Field: "Test"}
	fmt.Println(ErrorWrapPathElement(nil, e))
	fmt.Println(ErrorWrapPathElement(errors.New("error"), e))
	// Output:
	// <nil>
	// path field "Test": error
}

func ExampleGetErrorPath() {
	err := errors.New("error")
	err = ErrorWrapPathElement(err, &FieldPathElem{Field: "Foo"})
	err = ErrorWrapPathElement(err, &FieldPathElem{Field: "Bar"})
	fmt.Println(err)
	fmt.Println(GetErrorPath(err))
	// Output:
	// path field "Bar": path field "Foo": error
	// .Bar.Foo
}
