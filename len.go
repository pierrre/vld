package vld

import (
	"errors"
	"fmt"
)

func lenEqual[T any](length int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l != length {
			err := fmt.Errorf("length %d is not equal to %d", l, length)
			err = ErrorWrapLocalization(err, "LenEqual", l, length)
			return err
		}
		return nil
	}
}

func lenMin[T any](minLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l < minLen {
			err := fmt.Errorf("length %d is less than %d", l, minLen)
			err = ErrorWrapLocalization(err, "LenMin", l, minLen)
			return err
		}
		return nil
	}
}

func lenMax[T any](maxLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l > maxLen {
			err := fmt.Errorf("length %d is greater than %d", l, maxLen)
			err = ErrorWrapLocalization(err, "LenMax", l, maxLen)
			return err
		}
		return nil
	}
}

func lenRange[T any](minLen, maxLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l < minLen || l > maxLen {
			err := fmt.Errorf("length %d is not in the range [%d, %d]", l, minLen, maxLen)
			err = ErrorWrapLocalization(err, "LenRange", l, minLen, maxLen)
			return err
		}
		return nil
	}
}

func empty[T any](getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l != 0 {
			err := fmt.Errorf("is not empty (%d)", l)
			err = ErrorWrapLocalization(err, "NotEmpty", l)
			return err
		}
		return nil
	}
}

func notEmpty[T any](getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l == 0 {
			err := errors.New("is empty")
			err = ErrorWrapLocalization(err, "Empty")
			return err
		}
		return nil
	}
}
