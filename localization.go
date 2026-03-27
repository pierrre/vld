package vld

import (
	"errors"
	"fmt"
)

// LocalizationMessages contains localization messages.
var LocalizationMessages = map[string]map[string]string{
	"EqualError": {
		"en": "Value %#[1]v is not equal to %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas égale à %#[2]v.",
	},
	"NotEqualError": {
		"en": "Value %#[1]v is equal to %#[2]v.",
		"fr": "La valeur %#[1]v est égale à %#[2]v.",
	},
	"InError": {
		"en": "Value %#[1]v is not in %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas dans %#[2]v.",
	},
	"NotInError": {
		"en": "Value %#[1]v is in %#[2]v.",
		"fr": "La valeur %#[1]v est dans %#[2]v.",
	},
	"LenEqualError": {
		"en": "Length %[1]d is not equal to %[2]d.",
		"fr": "La longueur %[1]d n'est pas égale à %[2]d.",
	},
	"LenMinError": {
		"en": "Length %[1]d is less than %[2]d.",
		"fr": "La longueur %[1]d est inférieure à %[2]d.",
	},
	"LenMaxError": {
		"en": "Length %[1]d is greater than %[2]d.",
		"fr": "La longueur %[1]d est supérieure à %[2]d.",
	},
	"LenRangeError": {
		"en": "Length %[1]d is not in the range [%[2]d, %[3]d].",
		"fr": "La longueur %[1]d n'est pas dans l'intervalle [%[2]d, %[3]d].",
	},
	"NotEmptyError": {
		"en": "Value is not empty (%[1]d).",
		"fr": "La valeur n'est pas vide (%[1]d).",
	},
	"EmptyError": {
		"en": "Value is empty.",
		"fr": "La valeur est vide.",
	},
	"MinError": {
		"en": "Value %#[1]v is less than %#[2]v.",
		"fr": "La valeur %#[1]v est inférieure à %#[2]v.",
	},
	"MaxError": {
		"en": "Value %#[1]v is greater than %#[2]v.",
		"fr": "La valeur %#[1]v est supérieure à %#[2]v.",
	},
	"RangeError": {
		"en": "Value %#[1]v is not in the range [%#[2]v, %#[3]v].",
		"fr": "La valeur %#[1]v n'est pas dans l'intervalle [%#[2]v, %#[3]v].",
	},
	"LessError": {
		"en": "Value %#[1]v is not less than %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas inférieure à %#[2]v.",
	},
	"GreaterError": {
		"en": "Value %#[1]v is not greater than %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas supérieure à %#[2]v.",
	},
	"RegexpMatchError": {
		"en": "String %[1]q does not match regexp %[2]q.",
		"fr": "La chaîne %[1]q ne correspond pas à l'expression régulière %[2]q.",
	},
	"RegexpNotMatchError": {
		"en": "String %[1]q matches regexp %[2]q.",
		"fr": "La chaîne %[1]q correspond à l'expression régulière %[2]q.",
	},
	"PositiveError": {
		"en": "Value %#[1]v is not positive.",
		"fr": "La valeur %#[1]v n'est pas positive.",
	},
	"NegativeError": {
		"en": "Value %#[1]v is not negative.",
		"fr": "La valeur %#[1]v n'est pas négative.",
	},
	"SliceContainsError": {
		"en": "Slice does not contain %#[1]v.",
		"fr": "Le slice ne contient pas %#[1]v.",
	},
	"SliceNotContainsError": {
		"en": "Slice contains %#[1]v.",
		"fr": "Le slice contient %#[1]v.",
	},
	"SliceUniqueError": {
		"en": "Duplicate %#[1]v (index %[2]d).",
		"fr": "Doublon %#[1]v (indice %[2]d).",
	},
	"StringContainsError": {
		"en": "String %[1]q does not contain %[2]q.",
		"fr": "La chaîne %[1]q ne contient pas %[2]q.",
	},
	"StringNotContainsError": {
		"en": "String %[1]q contains %[2]q.",
		"fr": "La chaîne %[1]q contient %[2]q.",
	},
	"StringHasPrefixError": {
		"en": "String %[1]q does not begin with %[2]q.",
		"fr": "La chaîne %[1]q ne commence pas par %[2]q.",
	},
	"StringNotHasPrefixError": {
		"en": "String %[1]q begins with %[2]q.",
		"fr": "La chaîne %[1]q commence par %[2]q.",
	},
	"StringHasSuffixError": {
		"en": "String %[1]q does not end with %[2]q.",
		"fr": "La chaîne %[1]q ne se termine pas par %[2]q.",
	},
	"StringNotHasSuffixError": {
		"en": "String %[1]q ends with %[2]q.",
		"fr": "La chaîne %[1]q se termine par %[2]q.",
	},
	"BytesEqualError": {
		"en": "Bytes %[1]q is not equal to %[2]q.",
		"fr": "Les bytes %[1]q ne sont pas égaux à %[2]q.",
	},
	"BytesNotEqualError": {
		"en": "Bytes %[1]q is equal to %[2]q.",
		"fr": "Les bytes %[1]q sont égaux à %[2]q.",
	},
	"BytesContainsError": {
		"en": "Bytes %[1]q does not contain %[2]q.",
		"fr": "Les bytes %[1]q ne contiennent pas %[2]q.",
	},
	"BytesNotContainsError": {
		"en": "Bytes %[1]q contains %[2]q.",
		"fr": "Les bytes %[1]q contiennent %[2]q.",
	},
	"BytesHasPrefixError": {
		"en": "Bytes %[1]q does not have prefix %[2]q.",
		"fr": "Les bytes %[1]q n'ont pas le préfixe %[2]q.",
	},
	"BytesNotHasPrefixError": {
		"en": "Bytes %[1]q has prefix %[2]q.",
		"fr": "Les bytes %[1]q ont le préfixe %[2]q.",
	},
	"BytesHasSuffixError": {
		"en": "Bytes %[1]q does not have suffix %[2]q.",
		"fr": "Les bytes %[1]q n'ont pas le suffixe %[2]q.",
	},
	"BytesNotHasSuffixError": {
		"en": "Bytes %[1]q has suffix %[2]q.",
		"fr": "Les bytes %[1]q ont le suffixe %[2]q.",
	},
	"ZeroError": {
		"en": "Value %#[1]v is not zero.",
		"fr": "La valeur %#[1]v n'est pas zéro.",
	},
	"NotZeroError": {
		"en": "Value is zero.",
		"fr": "La valeur est zéro.",
	},
	"RequiredError": {
		"en": "Value is required.",
		"fr": "La valeur est requise.",
	},
	"PointerRequiredError": {
		"en": "Pointer is nil.",
		"fr": "Le pointeur est nil.", //nolint:misspell // "pointeur" is correct in French
	},
	"TypeRequiredError": {
		"en": "Type %[1]T cannot be converted to %[2]T.",
		"fr": "Le type %[1]T ne peut pas être converti en %[2]T.",
	},
}

// GetLocalizedMessage returns the localized message for a given key and arguments, using the provided locales (by order of preference).
func GetLocalizedMessage(key string, args []any, locales ...string) string {
	formats, ok := LocalizationMessages[key]
	if !ok {
		return ""
	}
	for _, locale := range locales {
		format, ok := formats[locale]
		if ok {
			return fmt.Sprintf(format, args...)
		}
	}
	return ""
}

// Localizable is an interface for types that can provide a localization key and arguments.
type Localizable interface {
	Localization() (key string, args []any)
}

// GetLocalizableMessage returns the localized message for a given [Localizable] and locales (by order of preference).
func GetLocalizableMessage(l Localizable, locales ...string) string {
	key, args := l.Localization()
	return GetLocalizedMessage(key, args, locales...)
}

// LocalizableError is an interface for errors that can provide localized messages.
type LocalizableError interface {
	error
	Localizable
}

// GetErrorLocalizedMessage returns the localized message for a given error and locales (by order of preference).
func GetErrorLocalizedMessage(err error, locales ...string) string {
	if err == nil {
		return ""
	}
	lv, ok := errors.AsType[LocalizableError](err)
	if !ok {
		return ""
	}
	return GetLocalizableMessage(lv, locales...)
}
