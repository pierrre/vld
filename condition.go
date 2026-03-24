package vld

import (
	"fmt"
	"strings"
)

// If returns a [Validator] that validates the value if the condition is true.
func If[T any](cond func(v T) bool, vr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("If(%s, %v)", getFuncName(cond), vr) }, func(v T) error {
		if !cond(v) {
			return nil
		}
		return vr.Validate(v)
	})
}

// IfElse returns a [Validator] that validates the value with the thenVr validator if the condition is true, or with the elseVr validator if the condition is false.
func IfElse[T any](cond func(v T) bool, thenVr Validator[T], elseVr Validator[T]) Validator[T] {
	return WithStringFunc(func() string { return fmt.Sprintf("IfElse(%s, %v, %v)", getFuncName(cond), thenVr, elseVr) }, func(v T) error {
		if cond(v) {
			return thenVr.Validate(v)
		}
		return elseVr.Validate(v)
	})
}

// Switch returns a new [SwitchValidator] with the given cases.
func Switch[T any](cases ...*SwitchCase[T]) SwitchValidator[T] {
	return SwitchValidator[T]{
		Cases: cases,
	}
}

// SwitchValidator is a [Validator] that validates the value with the first validator whose condition returns true.
type SwitchValidator[T any] struct {
	Cases []*SwitchCase[T]
}

// Validate implements [Validator].
func (sv SwitchValidator[T]) Validate(v T) error {
	for _, c := range sv.Cases {
		if c.Condition(v) {
			return c.Validator.Validate(v) //nolint:wrapcheck // Not needed.
		}
	}
	return nil
}

func (sv SwitchValidator[T]) String() string {
	sb := new(strings.Builder)
	sb.WriteString("Switch(")
	if len(sv.Cases) > 0 {
		sb.WriteString("\n")
		for _, c := range sv.Cases {
			writeStringIndent(sb, c.String())
			sb.WriteString(",\n")
		}
	}
	sb.WriteString(")")
	return sb.String()
}

// Case returns a new [SwitchCase] with the given condition and [Validator].
func Case[T any](cond func(v T) bool, vr Validator[T]) *SwitchCase[T] {
	return &SwitchCase[T]{
		Condition: cond,
		Validator: vr,
	}
}

// SwitchCase is a case for [SwitchValidator].
type SwitchCase[T any] struct {
	Condition func(v T) bool
	Validator Validator[T]
}

func (sc *SwitchCase[T]) String() string {
	return fmt.Sprintf("Case(%s, %v)", getFuncName(sc.Condition), sc.Validator)
}
