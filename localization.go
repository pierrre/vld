package vld

import (
	"errors"
	"fmt"
)

// LocalizedMessages contains localized messages.
var LocalizedMessages = map[string]map[string]string{
	"Equal": {
		"en": "Value %#[1]v is not equal to %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas égale à %#[2]v.",
	},
	"NotEqual": {
		"en": "Value %#[1]v is equal to %#[2]v.",
		"fr": "La valeur %#[1]v est égale à %#[2]v.",
	},
	"In": {
		"en": "Value %#[1]v is not in %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas dans %#[2]v.",
	},
	"NotIn": {
		"en": "Value %#[1]v is in %#[2]v.",
		"fr": "La valeur %#[1]v est dans %#[2]v.",
	},
	"LenEqual": {
		"en": "Length %[1]d is not equal to %[2]d.",
		"fr": "La longueur %[1]d n'est pas égale à %[2]d.",
	},
	"LenMin": {
		"en": "Length %[1]d is less than %[2]d.",
		"fr": "La longueur %[1]d est inférieure à %[2]d.",
	},
	"LenMax": {
		"en": "Length %[1]d is greater than %[2]d.",
		"fr": "La longueur %[1]d est supérieure à %[2]d.",
	},
	"LenRange": {
		"en": "Length %[1]d is not in the range [%[2]d, %[3]d].",
		"fr": "La longueur %[1]d n'est pas dans l'intervalle [%[2]d, %[3]d].",
	},
	"NotEmpty": {
		"en": "Value is not empty (%[1]d).",
		"fr": "La valeur n'est pas vide (%[1]d).",
	},
	"Empty": {
		"en": "Value is empty.",
		"fr": "La valeur est vide.",
	},
	"Min": {
		"en": "Value %#[1]v is less than %#[2]v.",
		"fr": "La valeur %#[1]v est inférieure à %#[2]v.",
	},
	"Max": {
		"en": "Value %#[1]v is greater than %#[2]v.",
		"fr": "La valeur %#[1]v est supérieure à %#[2]v.",
	},
	"Range": {
		"en": "Value %#[1]v is not in the range [%#[2]v, %#[3]v].",
		"fr": "La valeur %#[1]v n'est pas dans l'intervalle [%#[2]v, %#[3]v].",
	},
	"Less": {
		"en": "Value %#[1]v is not less than %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas inférieure à %#[2]v.",
	},
	"Greater": {
		"en": "Value %#[1]v is not greater than %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas supérieure à %#[2]v.",
	},
	"RegexpMatch": {
		"en": "String %[1]q does not match regexp %[2]q.",
		"fr": "La chaîne %[1]q ne correspond pas à l'expression régulière %[2]q.",
	},
	"RegexpNotMatch": {
		"en": "String %[1]q matches regexp %[2]q.",
		"fr": "La chaîne %[1]q correspond à l'expression régulière %[2]q.",
	},
	"Positive": {
		"en": "Value %#[1]v is not positive.",
		"fr": "La valeur %#[1]v n'est pas positive.",
	},
	"Negative": {
		"en": "Value %#[1]v is not negative.",
		"fr": "La valeur %#[1]v n'est pas négative.",
	},
	"SliceContains": {
		"en": "Slice does not contain %#[1]v.",
		"fr": "Le slice ne contient pas %#[1]v.",
	},
	"SliceNotContains": {
		"en": "Slice contains %#[1]v.",
		"fr": "Le slice contient %#[1]v.",
	},
	"SliceUnique": {
		"en": "Duplicate %#[1]v (index %[2]d).",
		"fr": "Doublon %#[1]v (indice %[2]d).",
	},
	"StringContains": {
		"en": "String %[1]q does not contain %[2]q.",
		"fr": "La chaîne %[1]q ne contient pas %[2]q.",
	},
	"StringNotContains": {
		"en": "String %[1]q contains %[2]q.",
		"fr": "La chaîne %[1]q contient %[2]q.",
	},
	"StringHasPrefix": {
		"en": "String %[1]q does not begin with %[2]q.",
		"fr": "La chaîne %[1]q ne commence pas par %[2]q.",
	},
	"StringNotHasPrefix": {
		"en": "String %[1]q begins with %[2]q.",
		"fr": "La chaîne %[1]q commence par %[2]q.",
	},
	"StringHasSuffix": {
		"en": "String %[1]q does not end with %[2]q.",
		"fr": "La chaîne %[1]q ne se termine pas par %[2]q.",
	},
	"StringNotHasSuffix": {
		"en": "String %[1]q ends with %[2]q.",
		"fr": "La chaîne %[1]q se termine par %[2]q.",
	},
	"Zero": {
		"en": "Value %#[1]v is not zero.",
		"fr": "La valeur %#[1]v n'est pas zéro.",
	},
	"NotZero": {
		"en": "Value is zero.",
		"fr": "La valeur est zéro.",
	},
	"Required": {
		"en": "Value is required.",
		"fr": "La valeur est requise.",
	},
	"PointerRequired": {
		"en": "Pointer is nil.",
		"fr": "Le pointeur est nil.", //nolint:misspell // "pointeur" is correct in French
	},
	"TypeRequired": {
		"en": "Type %[1]T cannot be converted to %[2]T.",
		"fr": "Le type %[1]T ne peut pas être converti en %[2]T.",
	},
}

// LocalizedError is an interface for errors that can provide localized messages.
type LocalizedError interface {
	error
	LocalizationKey() string
	LocalizationArgs() []any
}

// ErrorWrapLocalization wraps the error with a localized message.
func ErrorWrapLocalization(err error, key string, args ...any) error {
	if err == nil {
		return nil
	}
	return ErrorWrap(err, func(err error) error {
		return &localizedError{
			error: err,
			key:   key,
			args:  args,
		}
	})
}

type localizedError struct {
	error
	key  string
	args []any
}

func (e *localizedError) LocalizationKey() string {
	return e.key
}

func (e *localizedError) LocalizationArgs() []any {
	return e.args
}

// GetErrorLocalization returns the localized message for a given error and locales (by order of preference).
func GetErrorLocalization(err error, locales ...string) string {
	if err == nil {
		return ""
	}
	lv, ok := errors.AsType[LocalizedError](err)
	if !ok {
		return ""
	}
	key := lv.LocalizationKey()
	formats, ok := LocalizedMessages[key]
	if !ok {
		return ""
	}
	for _, locale := range locales {
		format, ok := formats[locale]
		if ok {
			return fmt.Sprintf(format, lv.LocalizationArgs()...)
		}
	}
	return ""
}
