package vld

import (
	"errors"
	"fmt"
)

func lenEqual[T any](length int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l != length {
			return fmt.Errorf("length %d is not equal to %d", l, length)
		}
		return nil
	}
}

func lenMin[T any](minLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l < minLen {
			return fmt.Errorf("length %d is less than %d", l, minLen)
		}
		return nil
	}
}

func lenMax[T any](maxLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l > maxLen {
			return fmt.Errorf("length %d is greater than %d", l, maxLen)
		}
		return nil
	}
}

func lenRange[T any](minLen, maxLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l < minLen || l > maxLen {
			return fmt.Errorf("length %d is not in the range [%d, %d]", l, minLen, maxLen)
		}
		return nil
	}
}

func empty[T any](getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l != 0 {
			return fmt.Errorf("is not empty (%d)", l)
		}
		return nil
	}
}

func notEmpty[T any](getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l == 0 {
			return errors.New("is empty")
		}
		return nil
	}
}
