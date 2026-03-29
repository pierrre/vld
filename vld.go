// Package vld provides a validation library.
package vld

import (
	"fmt"
	"strings"
)

// Validator is an interface that validates a value of type T.
type Validator[T any] interface {
	Validate(v T) error
	String() string
}

// ValidatorFunc is a function that implements [Validator].
type ValidatorFunc[T any] func(T) error

// Validate implements [Validator].
func (f ValidatorFunc[T]) Validate(v T) error {
	return f(v)
}

func (f ValidatorFunc[T]) String() string {
	return fmt.Sprintf("ValidatorFunc(%s)", getFuncName(f))
}

// Localization implements [Localizable].
func (f ValidatorFunc[T]) Localization() (key string, args []any) {
	return "Func", []any{getFuncName(f)}
}

func buildMultiValidatorString[T any](name string, vrs ...Validator[T]) string {
	sb := new(strings.Builder)
	sb.WriteString(name)
	sb.WriteString("(")
	if len(vrs) > 0 {
		sb.WriteString("\n")
		for _, vr := range vrs {
			writeStringIndent(sb, vr.String())
			sb.WriteString(",\n")
		}
	}
	sb.WriteString(")")
	return sb.String()
}
