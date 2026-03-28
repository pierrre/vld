package vld

import (
	"fmt"
)

// Get creates a [GetValidator].
func Get[In, Out any](getFunc func(In) Out, vr Validator[Out]) *GetValidator[In, Out] {
	return &GetValidator[In, Out]{
		Func:      getFunc,
		Validator: vr,
	}
}

// GetValidator is a [Validator] that validates the value returned by the get function.
type GetValidator[In, Out any] struct {
	Func      func(In) Out
	Validator Validator[Out]
}

// Validate implements [Validator].
func (vr *GetValidator[In, Out]) Validate(v In) error {
	return validateGet(v, vr.Func, vr.Validator.Validate)
}

func (vr *GetValidator[In, Out]) String() string {
	return fmt.Sprintf("Get(%s, %v)", getFuncName(vr.Func), vr.Validator)
}

func get[In, Out any](getFunc func(In) Out, f func(Out) error) func(In) error {
	return func(v In) error {
		return validateGet(v, getFunc, f)
	}
}

func validateGet[In, Out any](v In, getFunc func(In) Out, f func(Out) error) error {
	return f(getFunc(v))
}

// Wrap creates a [WrapValidator].
func Wrap[T any](msg string, vr Validator[T]) *WrapValidator[T] {
	return &WrapValidator[T]{
		Message:   msg,
		Validator: vr,
	}
}

// WrapValidator is a [Validator] that validates the value and wraps any error with the message.
type WrapValidator[T any] struct {
	Message   string
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *WrapValidator[T]) Validate(v T) error {
	err := vr.Validator.Validate(v)
	if err != nil {
		return ErrorWrapMessage(err, vr.Message)
	}
	return nil
}

func (vr *WrapValidator[T]) String() string {
	return fmt.Sprintf("Wrap(%q, %v)", vr.Message, vr.Validator)
}

// Field creates a [FieldValidator].
func Field[In, Out any](name string, getFunc func(In) Out, vr Validator[Out]) *FieldValidator[In, Out] {
	return &FieldValidator[In, Out]{
		Name:      name,
		Func:      getFunc,
		Validator: vr,
	}
}

// FieldValidator is a [Validator] that validates a field returned by the get function.
type FieldValidator[In, Out any] struct {
	Name      string
	Func      func(In) Out
	Validator Validator[Out]
}

// Validate implements [Validator].
func (vr *FieldValidator[In, Out]) Validate(v In) error {
	return validateField(v, vr.Name, vr.Func, vr.Validator.Validate)
}

func (vr *FieldValidator[In, Out]) String() string {
	return fmt.Sprintf("Field(%q, %v)", vr.Name, vr.Validator)
}

func field[In, Out any](name string, getFunc func(In) Out, f func(Out) error) func(In) error {
	return func(v In) error {
		return validateField(v, name, getFunc, f)
	}
}

func validateField[In, Out any](v In, name string, getFunc func(In) Out, f func(Out) error) error {
	err := f(getFunc(v))
	if err != nil {
		return ErrorWrapPathElem(err, &FieldPathElem{
			Field: name,
		})
	}
	return nil
}

// Message creates a [MessageValidator].
func Message[T any](msg string, vr Validator[T]) *MessageValidator[T] {
	return &MessageValidator[T]{
		Message:   msg,
		Validator: vr,
	}
}

// MessageValidator is a [Validator] that validates the value and overrides the error message of the underlying validator.
type MessageValidator[T any] struct {
	Message   string
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *MessageValidator[T]) Validate(v T) error {
	err := vr.Validator.Validate(v)
	if err != nil {
		return &MessageError{Message: vr.Message}
	}
	return nil
}

func (vr *MessageValidator[T]) String() string {
	return fmt.Sprintf("Message(%q, %v)", vr.Message, vr.Validator)
}

// MessageError is the error type returned by [MessageValidator].
type MessageError struct {
	Message string
}

func (e *MessageError) Error() string {
	return e.Message
}
