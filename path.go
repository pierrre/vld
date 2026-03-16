package vld

import (
	"errors"
	"fmt"
	"strings"
)

// Path describes the path in a data structure.
type Path []PathElem

func (p Path) String() string {
	if len(p) == 0 {
		return "."
	}
	var sb strings.Builder
	for _, elem := range p {
		sb.WriteString(elem.PathElem())
	}
	return sb.String()
}

// PathElem represents an element in a [Path].
type PathElem interface {
	PathElem() string
	String() string
}

// FieldPathElem represents a field in a struct.
type FieldPathElem struct {
	Field string
}

// PathElem implements [PathElem].
func (e *FieldPathElem) PathElem() string {
	return "." + e.Field
}

func (e *FieldPathElem) String() string {
	return fmt.Sprintf("field %q", e.Field)
}

// IndexPathElem represents an index in a slice or array.
type IndexPathElem struct {
	Index int
}

// PathElem implements [PathElem].
func (e *IndexPathElem) PathElem() string {
	return fmt.Sprintf("[%d]", e.Index)
}

func (e *IndexPathElem) String() string {
	return fmt.Sprintf("index %d", e.Index)
}

// KeyPathElem represents a key in a map.
type KeyPathElem struct {
	Key any
}

// PathElem implements [PathElem].
func (e *KeyPathElem) PathElem() string {
	return fmt.Sprintf("[%#v]", e.Key)
}

func (e *KeyPathElem) String() string {
	return fmt.Sprintf("key %#v", e.Key)
}

// PointerPathElem represents a pointer dereference.
type PointerPathElem struct{}

// PathElem implements [PathElem].
func (e *PointerPathElem) PathElem() string {
	return "*"
}

func (e *PointerPathElem) String() string {
	return "pointer"
}

// PathElemError represents an error with a [PathElem].
type PathElemError struct {
	Err      error
	PathElem PathElem
}

func (err *PathElemError) Error() string {
	return fmt.Sprintf("path %v: %v", err.PathElem, err.Err)
}

func (err *PathElemError) Unwrap() error {
	return err.Err
}

// ErrorWrapPathElement wraps the error with the [PathElem].
// See [ErrorWrap] for details.
func ErrorWrapPathElement(err error, e PathElem) error {
	if err == nil {
		return nil
	}
	return ErrorWrap(err, func(err error) error {
		return &PathElemError{
			Err:      err,
			PathElem: e,
		}
	})
}

// GetErrorPath returns the [Path] of the error, by collecting the [PathElemError]s recursively.
func GetErrorPath(err error) Path {
	var path Path
	for {
		pErr, ok := errors.AsType[*PathElemError](err)
		if !ok {
			break
		}
		path = append(path, pErr.PathElem)
		err = pErr.Err
	}
	return path
}
