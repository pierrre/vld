package vld

import (
	"fmt"
	"slices"
)

// In creates a [InValidator].
func In[T comparable](vs ...T) *InValidator[T] {
	return &InValidator[T]{
		Values: vs,
	}
}

// InValidator is a [Validator] that checks if the value is in a list of values.
type InValidator[T comparable] struct {
	Values []T
}

// Validate implements [Validator].
func (vr *InValidator[T]) Validate(v T) error {
	if !slices.Contains(vr.Values, v) {
		return &InError[T]{
			Value:  v,
			Values: vr.Values,
		}
	}
	return nil
}

func (vr *InValidator[T]) String() string {
	return fmt.Sprintf("In(%#v)", vr.Values)
}

// InError is the error type returned by [InValidator].
type InError[T comparable] struct {
	Value  T
	Values []T
}

func (e *InError[T]) Error() string {
	return fmt.Sprintf("%#v is not in %#v", e.Value, e.Values)
}

// Localization implements [LocalizableError].
func (e *InError[T]) Localization() (key string, args []any) {
	return "InError", []any{e.Value, e.Values}
}

// NotIn creates a [NotInValidator].
func NotIn[T comparable](vs ...T) *NotInValidator[T] {
	return &NotInValidator[T]{
		Values: vs,
	}
}

// NotInValidator is a [Validator] that checks if the value is not in a list of values.
type NotInValidator[T comparable] struct {
	Values []T
}

// Validate implements [Validator].
func (vr *NotInValidator[T]) Validate(v T) error {
	if slices.Contains(vr.Values, v) {
		return &NotInError[T]{
			Value:  v,
			Values: vr.Values,
		}
	}
	return nil
}

func (vr *NotInValidator[T]) String() string {
	return fmt.Sprintf("NotIn(%#v)", vr.Values)
}

// NotInError is the error type returned by [NotInValidator].
type NotInError[T comparable] struct {
	Value  T
	Values []T
}

func (e *NotInError[T]) Error() string {
	return fmt.Sprintf("%#v is in %#v", e.Value, e.Values)
}

// Localization implements [LocalizableError].
func (e *NotInError[T]) Localization() (key string, args []any) {
	return "NotInError", []any{e.Value, e.Values}
}
