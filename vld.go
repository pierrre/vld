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
