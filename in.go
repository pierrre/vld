package vld

import (
	"fmt"
)

// In returns a [Validator] that checks if the value is in a list of values.
func In[T comparable](vs ...T) Validator[T] {
	m := make(map[T]struct{}, len(vs))
	for _, v := range vs {
		m[v] = struct{}{}
	}
	return WithStringFunc(func() string { return fmt.Sprintf("In(%#v)", vs) }, func(v T) error {
		_, ok := m[v]
		if !ok {
			return fmt.Errorf("%#v is not in %#v", v, vs)
		}
		return nil
	})
}

// NotIn returns a [Validator] that checks if the value is not in a list of values.
func NotIn[T comparable](vs ...T) Validator[T] {
	m := make(map[T]struct{}, len(vs))
	for _, v := range vs {
		m[v] = struct{}{}
	}
	return WithStringFunc(func() string { return fmt.Sprintf("NotIn(%#v)", vs) }, func(v T) error {
		_, ok := m[v]
		if ok {
			return fmt.Errorf("%#v is in %#v", v, vs)
		}
		return nil
	})
}
