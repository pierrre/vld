package vld

import (
	"fmt"
	"iter"
)

// SeqEach creates a [SeqEachValidator].
func SeqEach[T any](vr Validator[KeyValue[int, T]]) *SeqEachValidator[T] {
	return &SeqEachValidator[T]{
		Validator: vr,
	}
}

// SeqEachValidator is a [Validator] that checks each index and element of the [iter.Seq].
type SeqEachValidator[T any] struct {
	Validator Validator[KeyValue[int, T]]
}

// Validate implements [Validator].
func (vr *SeqEachValidator[T]) Validate(s iter.Seq[T]) error {
	return validateSeqEach(s, vr.Validator.Validate)
}

func (vr *SeqEachValidator[T]) String() string {
	return fmt.Sprintf("SeqEach(%v)", vr.Validator)
}

// SeqEachValue creates a [SeqEachValueValidator].
func SeqEachValue[T any](vr Validator[T]) *SeqEachValueValidator[T] {
	return &SeqEachValueValidator[T]{
		Validator: vr,
	}
}

// SeqEachValueValidator is a [Validator] that checks each element of the [iter.Seq].
type SeqEachValueValidator[T any] struct {
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *SeqEachValueValidator[T]) Validate(s iter.Seq[T]) error {
	return validateSeqEach(s, get(KeyValue[int, T].GetValue, vr.Validator.Validate))
}

func (vr *SeqEachValueValidator[T]) String() string {
	return fmt.Sprintf("SeqEachValue(%v)", vr.Validator)
}

func validateSeqEach[T any](s iter.Seq[T], f func(KeyValue[int, T]) error) error {
	var errs []error
	i := 0
	for v := range s {
		err := f(KeyValue[int, T]{Key: i, Value: v})
		if err != nil {
			errs = append(errs, ErrorWrapPathElem(err, &IndexPathElem{Index: i}))
		}
		i++
	}
	return ErrorJoin(errs...)
}

// Seq2Each creates a [Seq2EachValidator].
func Seq2Each[K, V any](vr Validator[KeyValue[K, V]]) *Seq2EachValidator[K, V] {
	return &Seq2EachValidator[K, V]{
		Validator: vr,
	}
}

// Seq2EachValidator is a [Validator] that checks each key and value of the [iter.Seq2].
type Seq2EachValidator[K, V any] struct {
	Validator Validator[KeyValue[K, V]]
}

// Validate implements [Validator].
func (vr *Seq2EachValidator[K, V]) Validate(s iter.Seq2[K, V]) error {
	return validateSeq2Each(s, vr.Validator.Validate)
}

func (vr *Seq2EachValidator[K, V]) String() string {
	return fmt.Sprintf("Seq2Each(%v)", vr.Validator)
}

// Seq2EachKey creates a [Seq2EachKeyValidator].
func Seq2EachKey[K, V any](vr Validator[K]) *Seq2EachKeyValidator[K, V] {
	return &Seq2EachKeyValidator[K, V]{
		Validator: vr,
	}
}

// Seq2EachKeyValidator is a [Validator] that checks each key of the [iter.Seq2].
type Seq2EachKeyValidator[K, V any] struct {
	Validator Validator[K]
}

// Validate implements [Validator].
func (vr *Seq2EachKeyValidator[K, V]) Validate(s iter.Seq2[K, V]) error {
	return validateSeq2Each(s, field("key", KeyValue[K, V].GetKey, vr.Validator.Validate))
}

func (vr *Seq2EachKeyValidator[K, V]) String() string {
	return fmt.Sprintf("Seq2EachKey(%v)", vr.Validator)
}

// Seq2EachValue creates a [Seq2EachValueValidator].
func Seq2EachValue[K, V any](vr Validator[V]) *Seq2EachValueValidator[K, V] {
	return &Seq2EachValueValidator[K, V]{
		Validator: vr,
	}
}

// Seq2EachValueValidator is a [Validator] that checks each value of the [iter.Seq2].
type Seq2EachValueValidator[K, V any] struct {
	Validator Validator[V]
}

// Validate implements [Validator].
func (vr *Seq2EachValueValidator[K, V]) Validate(s iter.Seq2[K, V]) error {
	return validateSeq2Each(s, field("value", KeyValue[K, V].GetValue, vr.Validator.Validate))
}

func (vr *Seq2EachValueValidator[K, V]) String() string {
	return fmt.Sprintf("Seq2EachValue(%v)", vr.Validator)
}

func validateSeq2Each[K, V any](s iter.Seq2[K, V], f func(KeyValue[K, V]) error) error {
	var errs []error
	i := 0
	for k, v := range s {
		err := f(KeyValue[K, V]{Key: k, Value: v})
		if err != nil {
			errs = append(errs, ErrorWrapPathElem(err, &IndexPathElem{Index: i}))
		}
		i++
	}
	return ErrorJoin(errs...)
}
