package vld

import (
	"fmt"
	"strings"
)

// If creates a [IfValidator].
func If[T any](cond func(v T) bool, vr Validator[T]) *IfValidator[T] {
	return &IfValidator[T]{
		Condition: cond,
		Validator: vr,
	}
}

// IfValidator is a [Validator] that validates the value if the condition is true.
type IfValidator[T any] struct {
	Condition func(v T) bool
	Validator Validator[T]
}

// Validate implements [Validator].
func (vr *IfValidator[T]) Validate(v T) error {
	if !vr.Condition(v) {
		return nil
	}
	return vr.Validator.Validate(v) //nolint:wrapcheck // Not needed.
}

func (vr *IfValidator[T]) String() string {
	return fmt.Sprintf("If(%s, %v)", getFuncName(vr.Condition), vr.Validator)
}

// Localize implements [Localizer].
func (vr *IfValidator[T]) Localize(locales ...string) string {
	sb := new(strings.Builder)
	sb.WriteString(Localize("If", []any{getFuncName(vr.Condition)}, locales...))
	sb.WriteString("\n")
	writeStringIndent(sb, LocalizeValidator(vr.Validator, locales...))
	return sb.String()
}

// IfElse creates a [IfElseValidator].
func IfElse[T any](cond func(v T) bool, thenVr Validator[T], elseVr Validator[T]) *IfElseValidator[T] {
	return &IfElseValidator[T]{
		Condition: cond,
		Then:      thenVr,
		Else:      elseVr,
	}
}

// IfElseValidator is a [Validator] that validates the value with the Then validator if the condition is true, or with the Else validator if the condition is false.
type IfElseValidator[T any] struct {
	Condition func(v T) bool
	Then      Validator[T]
	Else      Validator[T]
}

// Validate implements [Validator].
func (vr *IfElseValidator[T]) Validate(v T) error {
	if vr.Condition(v) {
		return vr.Then.Validate(v) //nolint:wrapcheck // Not needed.
	}
	return vr.Else.Validate(v) //nolint:wrapcheck // Not needed.
}

func (vr *IfElseValidator[T]) String() string {
	sb := new(strings.Builder)
	sb.WriteString("IfElse(")
	sb.WriteString(getFuncName(vr.Condition))
	sb.WriteString(",\n")
	writeStringIndent(sb, vr.Then.String())
	sb.WriteString(",\n")
	writeStringIndent(sb, vr.Else.String())
	sb.WriteString(",\n)")
	return sb.String()
}

// Localize implements [Localizer].
func (vr *IfElseValidator[T]) Localize(locales ...string) string {
	sb := new(strings.Builder)
	sb.WriteString(Localize("IfElse", []any{getFuncName(vr.Condition)}, locales...))
	sb.WriteString("\n")
	writeStringIndent(sb, LocalizeValidator(vr.Then, locales...))
	sb.WriteString("\n")
	writeStringIndent(sb, LocalizeValidator(vr.Else, locales...))
	return sb.String()
}

// Switch creates a [SwitchValidator].
func Switch[T any](cases ...*SwitchCase[T]) *SwitchValidator[T] {
	return &SwitchValidator[T]{
		Cases: cases,
	}
}

// SwitchValidator is a [Validator] that validates the value with the first validator whose condition returns true.
type SwitchValidator[T any] struct {
	Cases []*SwitchCase[T]
}

// Validate implements [Validator].
func (sv *SwitchValidator[T]) Validate(v T) error {
	for _, c := range sv.Cases {
		if c.Condition(v) {
			return c.Validator.Validate(v) //nolint:wrapcheck // Not needed.
		}
	}
	return nil
}

func (sv *SwitchValidator[T]) String() string {
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

// Localize implements [Localizer].
func (sv *SwitchValidator[T]) Localize(locales ...string) string {
	sb := new(strings.Builder)
	sb.WriteString(Localize("Switch", nil, locales...))
	for _, c := range sv.Cases {
		sb.WriteString("\n")
		writeStringIndent(sb, c.Localize(locales...))
	}
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

// Localize implements [Localizer].
func (sc *SwitchCase[T]) Localize(locales ...string) string {
	sb := new(strings.Builder)
	sb.WriteString(getFuncName(sc.Condition))
	sb.WriteString(": ")
	sb.WriteString(LocalizeValidator(sc.Validator, locales...))
	return sb.String()
}
