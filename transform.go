package vld

import (
	"errors"
	"fmt"
)

// Get returns a [Validator] that validates the value returned by the get function.
func Get[In, Out any](getFunc func(In) Out, vr Validator[Out]) Validator[In] {
	return WithStringFunc(func() string { return fmt.Sprintf("Get(%T => %T, %v)", *new(In), *new(Out), vr) }, get(getFunc, vr.Validate))
}

func get[In, Out any](getFunc func(In) Out, f func(Out) error) func(In) error {
	return func(v In) error {
		return f(getFunc(v))
	}
}

// Wrap returns a [Validator] that wraps the error with the message.
func Wrap[T any](msg string, vr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Wrap(%q, %v)", msg, vr) }, wrap(msg, vr.Validate))
}

func wrap[T any](msg string, f func(T) error) func(T) error {
	return func(v T) error {
		err := f(v)
		if err != nil {
			return ErrorWrapMessage(err, msg)
		}
		return nil
	}
}

// Field returns a [Validator] that validates a field returned by the get function.
func Field[In, Out any](name string, getFunc func(In) Out, vr Validator[Out]) Validator[In] {
	return WithStringFunc(func() string { return fmt.Sprintf("Field(%q, %v)", name, vr) }, field(name, getFunc, vr.Validate))
}

func field[In, Out any](name string, getFunc func(In) Out, f func(Out) error) func(In) error {
	return func(v In) error {
		err := f(getFunc(v))
		if err != nil {
			return ErrorWrapPathElement(err, &FieldPathElem{
				Field: name,
			})
		}
		return nil
	}
}

// Message returns a [Validator] that overrides the error message of the validator.
func Message[T any](msg string, vr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("Message(%q, %v)", msg, vr) }, func(v T) error {
		err := vr.Validate(v)
		if err != nil {
			return errors.New(msg)
		}
		return nil
	})
}
