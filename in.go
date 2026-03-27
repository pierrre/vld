package vld

import (
	"fmt"
)

// In returns a [Validator] that checks if the value is in a list of values.
func In[T comparable](vs ...T) Validator[T] {
	m := make(map[T]struct{}, len(vs))
	for _, v := range vs {
		m[v] = struct{}{}
	}
	return WithStringFunc(func() string { return fmt.Sprintf("In(%#v)", vs) }, func(v T) error {
		_, ok := m[v]
		if !ok {
			return &InError[T]{
				Value:  v,
				Values: vs,
			}
		}
		return nil
	})
}

// InError is the error type returned by [In].
type InError[T comparable] struct {
	Value  T
	Values []T
}

// Error implements [error].
func (e *InError[T]) Error() string {
	return fmt.Sprintf("%#v is not in %#v", e.Value, e.Values)
}

// Localization implements [LocalizableError].
func (e *InError[T]) Localization() (key string, args []any) {
	return "InError", []any{e.Value, e.Values}
}

// NotIn returns a [Validator] that checks if the value is not in a list of values.
func NotIn[T comparable](vs ...T) Validator[T] {
	m := make(map[T]struct{}, len(vs))
	for _, v := range vs {
		m[v] = struct{}{}
	}
	return WithStringFunc(func() string { return fmt.Sprintf("NotIn(%#v)", vs) }, func(v T) error {
		_, ok := m[v]
		if ok {
			return &NotInError[T]{
				Value:  v,
				Values: vs,
			}
		}
		return nil
	})
}

// NotInError is the error type returned by [NotIn].
type NotInError[T comparable] struct {
	Value  T
	Values []T
}

// Error implements [error].
func (e *NotInError[T]) Error() string {
	return fmt.Sprintf("%#v is in %#v", e.Value, e.Values)
}

// Localization implements [LocalizableError].
func (e *NotInError[T]) Localization() (key string, args []any) {
	return "NotInError", []any{e.Value, e.Values}
}
