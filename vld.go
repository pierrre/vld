// Package vld provides a validation library.
package vld

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
	return "ValidatorFunc"
}

// WithStringFunc returns a [Validator] with a custom string representation.
// It validates the value with the given function.
func WithStringFunc[T any](getString func() string, f func(T) error) Validator[T] {
	return &withStringFunc[T]{
		getString: getString,
		f:         f,
	}
}

type withStringFunc[T any] struct {
	getString func() string
	f         func(T) error
}

func (vr *withStringFunc[T]) Validate(v T) error {
	return vr.f(v)
}

func (vr *withStringFunc[T]) String() string {
	return vr.getString()
}
