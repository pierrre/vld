package vld

import (
	"fmt"
	"iter"
)

// SeqEach returns a [Validator] that checks each index and element of the [iter.Seq].
func SeqEach[T any](vr Validator[KeyValue[int, T]]) Validator[iter.Seq[T]] {
	return WithStringFunc(func() string { return fmt.Sprintf("SeqEach(%v)", vr) }, seqEach(vr.Validate))
}

// SeqEachValue returns a [Validator] that checks each element of the [iter.Seq].
func SeqEachValue[T any](vr Validator[T]) Validator[iter.Seq[T]] {
	return WithStringFunc(func() string { return fmt.Sprintf("SeqEachValue(%v)", vr) }, seqEach(get(KeyValue[int, T].GetValue, vr.Validate)))
}

func seqEach[T any](f func(KeyValue[int, T]) error) func(iter.Seq[T]) error {
	return func(s iter.Seq[T]) error {
		var errs []error
		i := 0
		for v := range s {
			err := f(KeyValue[int, T]{Key: i, Value: v})
			if err != nil {
				errs = append(errs, ErrorWrapMessagef(err, "%d", i))
			}
			i++
		}
		return ErrorJoin(errs...)
	}
}

// Seq2Each returns a [Validator] that checks each key and value of the [iter.Seq2].
func Seq2Each[K, V any](vr Validator[KeyValue[K, V]]) Validator[iter.Seq2[K, V]] {
	return WithStringFunc(func() string { return fmt.Sprintf("Seq2Each(%v)", vr) }, seq2Each(vr.Validate))
}

// Seq2EachKey returns a [Validator] that checks each key of the [iter.Seq2].
func Seq2EachKey[K, V any](vr Validator[K]) Validator[iter.Seq2[K, V]] {
	return WithStringFunc(func() string { return fmt.Sprintf("Seq2EachKey(%v)", vr) }, seq2Each(get(KeyValue[K, V].GetKey, vr.Validate)))
}

// Seq2EachValue returns a [Validator] that checks each value of the [iter.Seq2].
func Seq2EachValue[K, V any](vr Validator[V]) Validator[iter.Seq2[K, V]] {
	return WithStringFunc(func() string { return fmt.Sprintf("Seq2EachValue(%v)", vr) }, seq2Each(get(KeyValue[K, V].GetValue, vr.Validate)))
}

func seq2Each[K, V any](f func(KeyValue[K, V]) error) func(iter.Seq2[K, V]) error {
	return func(s iter.Seq2[K, V]) error {
		var errs []error
		i := 0
		for k, v := range s {
			err := f(KeyValue[K, V]{Key: k, Value: v})
			if err != nil {
				errs = append(errs, ErrorWrapMessagef(err, "%d", i))
			}
			i++
		}
		return ErrorJoin(errs...)
	}
}
