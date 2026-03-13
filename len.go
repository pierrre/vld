package vld

import (
	"errors"
	"fmt"
)

func lenEqual[T any](name string, length int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l != length {
			return fmt.Errorf("%s %d is not equal to %d", name, l, length)
		}
		return nil
	}
}

func lenMin[T any](name string, minLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l < minLen {
			return fmt.Errorf("%s %d is less than %d", name, l, minLen)
		}
		return nil
	}
}

func lenMax[T any](name string, maxLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l > maxLen {
			return fmt.Errorf("%s %d is greater than %d", name, l, maxLen)
		}
		return nil
	}
}

func lenRange[T any](name string, minLen, maxLen int, getLen func(v T) int) func(T) error {
	return func(v T) error {
		l := getLen(v)
		if l < minLen || l > maxLen {
			return fmt.Errorf("%s %d is not in the range [%d, %d]", name, l, minLen, maxLen)
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
