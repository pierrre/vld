package vld

import (
	"errors"
	"fmt"
	"strings"
)

// LocalizationMessages contains localization messages.
var LocalizationMessages = map[string]map[string]string{
	"And": {
		"en": "Value must satisfy all of the following validators (stop on first failure):",
		"fr": "La valeur doit satisfaire tous les validateurs suivants (arrêt à la première erreur):",
	},
	"Or": {
		"en": "Value must satisfy at least one of the following validators:",
		"fr": "La valeur doit satisfaire au moins un des validateurs suivants:",
	},
	"All": {
		"en": "The value must satisfy all of the following validators:",
		"fr": "La valeur doit satisfaire tous les validateurs suivants:",
	},
	"If": {
		"en": "Value must satisfy the following validator if the condition (%[1]s) is true:",
		"fr": "La valeur doit satisfaire le validateur suivant si la condition (%[1]s) est vraie:",
	},
	"IfElse": {
		"en": "Value must satisfy the first validator if the condition (%[1]s) is true, or the second validator otherwise:",
		"fr": "La valeur doit satisfaire le premier validateur si la condition (%[1]s) est vraie, ou le second validateur sinon:",
	},
	"Switch": {
		"en": "Value must satisfy the validator of the first case whose condition is true:",
		"fr": "La valeur doit satisfaire le validateur du premier cas dont la condition est vraie:",
	},
	"Equal": {
		"en": "Value must be equal to %#[1]v.",
		"fr": "La valeur doit être égale à %#[1]v.",
	},
	"EqualFunc": {
		"en": "Value must be equal to %#[1]v with function %[2]s.",
		"fr": "La valeur doit être égale à %#[1]v avec la fonction %[2]s.",
	},
	"EqualError": {
		"en": "Value %#[1]v is not equal to %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas égale à %#[2]v.",
	},
	"NotEqual": {
		"en": "Value must not be equal to %#[1]v.",
		"fr": "La valeur ne doit pas être égale à %#[1]v.",
	},
	"NotEqualFunc": {
		"en": "Value must not be equal to %#[1]v with function %[2]s.",
		"fr": "La valeur ne doit pas être égale à %#[1]v avec la fonction %[2]s.",
	},
	"NotEqualError": {
		"en": "Value %#[1]v is equal to %#[2]v.",
		"fr": "La valeur %#[1]v est égale à %#[2]v.",
	},
	"In": {
		"en": "Value must be in %#[1]v.",
		"fr": "La valeur doit être dans %#[1]v.",
	},
	"InError": {
		"en": "Value %#[1]v is not in %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas dans %#[2]v.",
	},
	"NotIn": {
		"en": "Value must not be in %#[1]v.",
		"fr": "La valeur ne doit pas être dans %#[1]v.",
	},
	"NotInError": {
		"en": "Value %#[1]v is in %#[2]v.",
		"fr": "La valeur %#[1]v est dans %#[2]v.",
	},
	"LenEqual": {
		"en": "Length must be equal to %[1]d.",
		"fr": "La longueur doit être égale à %[1]d.",
	},
	"LenEqualError": {
		"en": "Length %[1]d is not equal to %[2]d.",
		"fr": "La longueur %[1]d n'est pas égale à %[2]d.",
	},
	"LenMin": {
		"en": "Length must be greater than or equal to %[1]d.",
		"fr": "La longueur doit être supérieure ou égale à %[1]d.",
	},
	"LenMinError": {
		"en": "Length %[1]d is less than %[2]d.",
		"fr": "La longueur %[1]d est inférieure à %[2]d.",
	},
	"LenMax": {
		"en": "Length must be less than or equal to %[1]d.",
		"fr": "La longueur doit être inférieure ou égale à %[1]d.",
	},
	"LenMaxError": {
		"en": "Length %[1]d is greater than %[2]d.",
		"fr": "La longueur %[1]d est supérieure à %[2]d.",
	},
	"LenRange": {
		"en": "Length must be in the range [%[1]d, %[2]d].",
		"fr": "La longueur doit être dans l'intervalle [%[1]d, %[2]d].",
	},
	"LenRangeError": {
		"en": "Length %[1]d is not in the range [%[2]d, %[3]d].",
		"fr": "La longueur %[1]d n'est pas dans l'intervalle [%[2]d, %[3]d].",
	},
	"Empty": {
		"en": "Value must be empty.",
		"fr": "La valeur doit être vide.",
	},
	"EmptyError": {
		"en": "Value is not empty (%[1]d).",
		"fr": "La valeur n'est pas vide (%[1]d).",
	},
	"NotEmpty": {
		"en": "Value must not be empty.",
		"fr": "La valeur ne doit pas être vide.",
	},
	"NotEmptyError": {
		"en": "Value is empty.",
		"fr": "La valeur est vide.",
	},
	"Min": {
		"en": "Value must be greater than or equal to %#[1]v.",
		"fr": "La valeur doit être supérieure ou égale à %#[1]v.",
	},
	"MinFunc": {
		"en": "Value must be greater than or equal to %#[1]v with function %[2]s.",
		"fr": "La valeur doit être supérieure ou égale à %#[1]v avec la fonction %[2]s.",
	},
	"MinError": {
		"en": "Value %#[1]v is less than %#[2]v.",
		"fr": "La valeur %#[1]v est inférieure à %#[2]v.",
	},
	"Max": {
		"en": "Value must be less than or equal to %#[1]v.",
		"fr": "La valeur doit être inférieure ou égale à %#[1]v.",
	},
	"MaxFunc": {
		"en": "Value must be less than or equal to %#[1]v with function %[2]s.",
		"fr": "La valeur doit être inférieure ou égale à %#[1]v avec la fonction %[2]s.",
	},
	"MaxError": {
		"en": "Value %#[1]v is greater than %#[2]v.",
		"fr": "La valeur %#[1]v est supérieure à %#[2]v.",
	},
	"Range": {
		"en": "Value must be in the range [%#[1]v, %#[2]v].",
		"fr": "La valeur doit être dans l'intervalle [%#[1]v, %#[2]v].",
	},
	"RangeFunc": {
		"en": "Value must be in the range [%#[1]v, %#[2]v] with function %[3]s.",
		"fr": "La valeur doit être dans l'intervalle [%#[1]v, %#[2]v] avec la fonction %[3]s.",
	},
	"RangeError": {
		"en": "Value %#[1]v is not in the range [%#[2]v, %#[3]v].",
		"fr": "La valeur %#[1]v n'est pas dans l'intervalle [%#[2]v, %#[3]v].",
	},
	"Less": {
		"en": "Value must be less than %#[1]v.",
		"fr": "La valeur doit être inférieure à %#[1]v.",
	},
	"LessFunc": {
		"en": "Value must be less than %#[1]v with function %[2]s.",
		"fr": "La valeur doit être inférieure à %#[1]v avec la fonction %[2]s.",
	},
	"LessError": {
		"en": "Value %#[1]v is not less than %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas inférieure à %#[2]v.",
	},
	"Greater": {
		"en": "Value must be greater than %#[1]v.",
		"fr": "La valeur doit être supérieure à %#[1]v.",
	},
	"GreaterFunc": {
		"en": "Value must be greater than %#[1]v with function %[2]s.",
		"fr": "La valeur doit être supérieure à %#[1]v avec la fonction %[2]s.",
	},
	"GreaterError": {
		"en": "Value %#[1]v is not greater than %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas supérieure à %#[2]v.",
	},
	"RegexpMatch": {
		"en": "String must match regexp %[1]q.",
		"fr": "La chaîne doit correspondre à l'expression régulière %[1]q.",
	},
	"RegexpMatchError": {
		"en": "String %[1]q does not match regexp %[2]q.",
		"fr": "La chaîne %[1]q ne correspond pas à l'expression régulière %[2]q.",
	},
	"RegexpNotMatch": {
		"en": "String must not match regexp %[1]q.",
		"fr": "La chaîne ne doit pas correspondre à l'expression régulière %[1]q.",
	},
	"RegexpNotMatchError": {
		"en": "String %[1]q matches regexp %[2]q.",
		"fr": "La chaîne %[1]q correspond à l'expression régulière %[2]q.",
	},
	"Positive": {
		"en": "Value must be positive.",
		"fr": "La valeur doit être positive.",
	},
	"PositiveError": {
		"en": "Value %#[1]v is not positive.",
		"fr": "La valeur %#[1]v n'est pas positive.",
	},
	"Negative": {
		"en": "Value must be negative.",
		"fr": "La valeur doit être négative.",
	},
	"NegativeError": {
		"en": "Value %#[1]v is not negative.",
		"fr": "La valeur %#[1]v n'est pas négative.",
	},
	"SliceContains": {
		"en": "Slice must contain %#[1]v.",
		"fr": "Le slice doit contenir %#[1]v.",
	},
	"SliceContainsError": {
		"en": "Slice does not contain %#[1]v.",
		"fr": "Le slice ne contient pas %#[1]v.",
	},
	"SliceNotContains": {
		"en": "Slice must not contain %#[1]v.",
		"fr": "Le slice ne doit pas contenir %#[1]v.",
	},
	"SliceNotContainsError": {
		"en": "Slice contains %#[1]v.",
		"fr": "Le slice contient %#[1]v.",
	},
	"SliceUnique": {
		"en": "Slice must have unique elements.",
		"fr": "Le slice doit avoir des éléments uniques.",
	},
	"SliceUniqueFunc": {
		"en": "Slice must have unique elements with function %[1]s.",
		"fr": "Le slice doit avoir des éléments uniques avec la fonction %[1]s.",
	},
	"SliceUniqueError": {
		"en": "Duplicate %#[1]v (index %[2]d).",
		"fr": "Doublon %#[1]v (indice %[2]d).",
	},
	"StringContains": {
		"en": "String must contain %[1]q.",
		"fr": "La chaîne doit contenir %[1]q.",
	},
	"StringContainsError": {
		"en": "String %[1]q does not contain %[2]q.",
		"fr": "La chaîne %[1]q ne contient pas %[2]q.",
	},
	"StringNotContains": {
		"en": "String must not contain %[1]q.",
		"fr": "La chaîne ne doit pas contenir %[1]q.",
	},
	"StringNotContainsError": {
		"en": "String %[1]q contains %[2]q.",
		"fr": "La chaîne %[1]q contient %[2]q.",
	},
	"StringHasPrefix": {
		"en": "String must begin with %[1]q.",
		"fr": "La chaîne doit commencer par %[1]q.",
	},
	"StringHasPrefixError": {
		"en": "String %[1]q does not begin with %[2]q.",
		"fr": "La chaîne %[1]q ne commence pas par %[2]q.",
	},
	"StringNotHasPrefix": {
		"en": "String must not begin with %[1]q.",
		"fr": "La chaîne ne doit pas commencer par %[1]q.",
	},
	"StringNotHasPrefixError": {
		"en": "String %[1]q begins with %[2]q.",
		"fr": "La chaîne %[1]q commence par %[2]q.",
	},
	"StringHasSuffix": {
		"en": "String must end with %[1]q.",
		"fr": "La chaîne doit se terminer par %[1]q.",
	},
	"StringHasSuffixError": {
		"en": "String %[1]q does not end with %[2]q.",
		"fr": "La chaîne %[1]q ne se termine pas par %[2]q.",
	},
	"StringNotHasSuffix": {
		"en": "String must not end with %[1]q.",
		"fr": "La chaîne ne doit pas se terminer par %[1]q.",
	},
	"StringNotHasSuffixError": {
		"en": "String %[1]q ends with %[2]q.",
		"fr": "La chaîne %[1]q se termine par %[2]q.",
	},
	"BytesEqual": {
		"en": "Bytes must be equal to %[1]q.",
		"fr": "Les bytes doivent être égaux à %[1]q.",
	},
	"BytesEqualError": {
		"en": "Bytes %[1]q is not equal to %[2]q.",
		"fr": "Les bytes %[1]q ne sont pas égaux à %[2]q.",
	},
	"BytesNotEqual": {
		"en": "Bytes must not be equal to %[1]q.",
		"fr": "Les bytes ne doivent pas être égaux à %[1]q.",
	},
	"BytesNotEqualError": {
		"en": "Bytes %[1]q is equal to %[2]q.",
		"fr": "Les bytes %[1]q sont égaux à %[2]q.",
	},
	"BytesContains": {
		"en": "Bytes must contain %[1]q.",
		"fr": "Les bytes doivent contenir %[1]q.",
	},
	"BytesContainsError": {
		"en": "Bytes %[1]q does not contain %[2]q.",
		"fr": "Les bytes %[1]q ne contiennent pas %[2]q.",
	},
	"BytesNotContains": {
		"en": "Bytes must not contain %[1]q.",
		"fr": "Les bytes ne doivent pas contenir %[1]q.",
	},
	"BytesNotContainsError": {
		"en": "Bytes %[1]q contains %[2]q.",
		"fr": "Les bytes %[1]q contiennent %[2]q.",
	},
	"BytesHasPrefix": {
		"en": "Bytes must have prefix %[1]q.",
		"fr": "Les bytes doivent avoir le préfixe %[1]q.",
	},
	"BytesHasPrefixError": {
		"en": "Bytes %[1]q does not have prefix %[2]q.",
		"fr": "Les bytes %[1]q n'ont pas le préfixe %[2]q.",
	},
	"BytesNotHasPrefix": {
		"en": "Bytes must not have prefix %[1]q.",
		"fr": "Les bytes ne doivent pas avoir le préfixe %[1]q.",
	},
	"BytesNotHasPrefixError": {
		"en": "Bytes %[1]q has prefix %[2]q.",
		"fr": "Les bytes %[1]q ont le préfixe %[2]q.",
	},
	"BytesHasSuffix": {
		"en": "Bytes must have suffix %[1]q.",
		"fr": "Les bytes doivent avoir le suffixe %[1]q.",
	},
	"BytesHasSuffixError": {
		"en": "Bytes %[1]q does not have suffix %[2]q.",
		"fr": "Les bytes %[1]q n'ont pas le suffixe %[2]q.",
	},
	"BytesNotHasSuffix": {
		"en": "Bytes must not have suffix %[1]q.",
		"fr": "Les bytes ne doivent pas avoir le suffixe %[1]q.",
	},
	"BytesNotHasSuffixError": {
		"en": "Bytes %[1]q has suffix %[2]q.",
		"fr": "Les bytes %[1]q ont le suffixe %[2]q.",
	},
	"Zero": {
		"en": "Value must be zero.",
		"fr": "La valeur doit être zéro.",
	},
	"ZeroError": {
		"en": "Value %#[1]v is not zero.",
		"fr": "La valeur %#[1]v n'est pas zéro.",
	},
	"NotZero": {
		"en": "Value must not be zero.",
		"fr": "La valeur ne doit pas être zéro.",
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
	"Func": {
		"en": "Function %[1]s.",
		"fr": "Fonction %[1]s.",
	},
}

// GetLocalizationMessage returns the localization message for a given key and locales (by order of preference).
// The bool return value indicates whether a message was found for the given key and locales.
func GetLocalizationMessage(key string, locales ...string) (string, bool) {
	formats, ok := LocalizationMessages[key]
	if !ok {
		return "", false
	}
	for _, locale := range locales {
		format, ok := formats[locale]
		if ok {
			return format, true
		}
	}
	return "", false
}

// Localize returns the localized message for a given key and arguments, using the provided locales (by order of preference).
func Localize(key string, args []any, locales ...string) string {
	format, ok := GetLocalizationMessage(key, locales...)
	if !ok {
		return ""
	}
	return fmt.Sprintf(format, args...)
}

// Localizable is an interface for types that can provide a localization key and arguments.
type Localizable interface {
	Localization() (key string, args []any)
}

// LocalizeLocalizable returns the localized message for a given [Localizable] and locales (by order of preference).
func LocalizeLocalizable(l Localizable, locales ...string) string {
	key, args := l.Localization()
	return Localize(key, args, locales...)
}

// Localizer is an interface for types that can provide a localized message for given locales.
type Localizer interface {
	Localize(locales ...string) string
}

// LocalizableError is an interface for errors that can provide localized messages.
type LocalizableError interface {
	error
	Localizable
}

// LocalizeError returns the localized message for a given error and locales (by order of preference).
func LocalizeError(err error, locales ...string) string {
	if err == nil {
		return ""
	}
	lv, ok := errors.AsType[LocalizableError](err)
	if !ok {
		return ""
	}
	return LocalizeLocalizable(lv, locales...)
}

// LocalizeValidator returns the localized message for a given [Validator] and locales (by order of preference).
func LocalizeValidator[T any](vr Validator[T], locales ...string) string {
	var s string
	switch vr := vr.(type) {
	case Localizer:
		s = vr.Localize(locales...)
	case Localizable:
		s = LocalizeLocalizable(vr, locales...)
	}
	return s
}

func localizeMultiValidator[T any](key string, vrs []Validator[T], locales ...string) string {
	sb := new(strings.Builder)
	sb.WriteString(Localize(key, nil, locales...))
	for _, vr := range vrs {
		sb.WriteString("\n")
		writeStringIndent(sb, LocalizeValidator(vr, locales...))
	}
	return sb.String()
}

// TODO: localization
// PointerOptionalValidator
// PointerRequiredValidator
// OptionalValidator
// RequiredValidator
// TypeOptionalValidator
// TypeRequiredValidator
// GetValidator
// WrapValidator
// FieldValidator
// MessageValidator
// MapEachValidator
// MapEachKeyValidator
// MapEachValueValidator
// MapSortedEachValidator
// MapSortedEachKeyValidator
// MapSortedEachValueValidator
// SliceEachValidator
// SliceEachValueValidator
// SeqEachValidator
// SeqEachValueValidator
// Seq2EachValidator
// Seq2EachKeyValidator
// Seq2EachValueValidator
