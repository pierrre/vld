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

// Localization implements [Localizable].
func (vr *GetValidator[In, Out]) Localization() (key string, args []any) {
	return "GetValidator", []any{getFuncName(vr.Func), vr.Validator}
}

func get[In, Out any](getFunc func(In) Out, f func(Out) error) func(In) error {
	return func(v In) error {
		return validateGet(v, getFunc, f)
	}
}

func validateGet[In, Out any](v In, getFunc func(In) Out, f func(Out) error) error {
	return f(getFunc(v))
}

// Parse creates a [ParseValidator].
func Parse[In, Out any](parseFunc func(In) (Out, error), vr Validator[Out]) *ParseValidator[In, Out] {
	return &ParseValidator[In, Out]{
		Func:      parseFunc,
		Validator: vr,
	}
}

// ParseValidator is a [Validator] that parses the value with the function and validates the result.
type ParseValidator[In, Out any] struct {
	Func      func(In) (Out, error)
	Validator Validator[Out]
}

// Validate implements [Validator].
func (vr *ParseValidator[In, Out]) Validate(v In) error {
	pv, err := vr.Func(v)
	if err != nil {
		return &ParseError[In, Out]{
			Err:   err,
			Value: v,
			Func:  vr.Func,
		}
	}
	return vr.Validator.Validate(pv) //nolint:wrapcheck // Not needed.
}

func (vr *ParseValidator[In, Out]) String() string {
	return fmt.Sprintf("Parse(%s, %v)", getFuncName(vr.Func), vr.Validator)
}

// Localization implements [Localizable].
func (vr *ParseValidator[In, Out]) Localization() (key string, args []any) {
	return "ParseValidator", []any{getFuncName(vr.Func), vr.Validator}
}

// ParseError is the error type returned by [ParseValidator] when the parsing fails.
type ParseError[In, Out any] struct {
	Err   error
	Value In
	Func  func(In) (Out, error)
}

func (e *ParseError[In, Out]) Error() string {
	return fmt.Sprintf("parse %#v with %s: %v", e.Value, getFuncName(e.Func), e.Err)
}

// Localization implements [Localizable].
func (e *ParseError[In, Out]) Localization() (key string, args []any) {
	return "ParseError", []any{e.Value, getFuncName(e.Func), e.Err}
}

func (e *ParseError[In, Out]) Unwrap() error {
	return e.Err
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

// Localization implements [Localizable].
func (vr *WrapValidator[T]) Localization() (key string, args []any) {
	return "WrapValidator", []any{vr.Message, vr.Validator}
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
	return fmt.Sprintf("Field(%q, %s, %v)", vr.Name, getFuncName(vr.Func), vr.Validator)
}

// Localization implements [Localizable].
func (vr *FieldValidator[In, Out]) Localization() (key string, args []any) {
	return "FieldValidator", []any{vr.Name, getFuncName(vr.Func), vr.Validator}
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

// Localization implements [Localizable].
func (vr *MessageValidator[T]) Localization() (key string, args []any) {
	return "MessageValidator", []any{vr.Message, vr.Validator}
}

// MessageError is the error type returned by [MessageValidator].
type MessageError struct {
	Message string
}

func (e *MessageError) Error() string {
	return e.Message
}
