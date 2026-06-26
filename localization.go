package vld

import (
	"errors"
	"fmt"
	"strings"
)

// LocalizationMessages contains localization messages.
var LocalizationMessages = map[string]map[string]string{
	"AndValidator": {
		"en": "Value must satisfy all of the following validators (stop on first failure):",
		"fr": "La valeur doit satisfaire tous les validateurs suivants (arrêt à la première erreur):",
	},
	"OrValidator": {
		"en": "Value must satisfy at least one of the following validators:",
		"fr": "La valeur doit satisfaire au moins un des validateurs suivants:",
	},
	"AllValidator": {
		"en": "Value must satisfy all of the following validators:",
		"fr": "La valeur doit satisfaire tous les validateurs suivants:",
	},
	"IfValidator": {
		"en": "Value must satisfy the following validator if the condition (%[1]s) is true:",
		"fr": "La valeur doit satisfaire le validateur suivant si la condition (%[1]s) est vraie:",
	},
	"IfElseValidator": {
		"en": "Value must satisfy the first validator if the condition (%[1]s) is true, or the second validator otherwise:",
		"fr": "La valeur doit satisfaire le premier validateur si la condition (%[1]s) est vraie, ou le second validateur sinon:",
	},
	"SwitchValidator": {
		"en": "Value must satisfy the validator of the first case whose condition is true:",
		"fr": "La valeur doit satisfaire le validateur du premier cas dont la condition est vraie:",
	},
	"GetValidator": {
		"en": "Value returned by function %[1]s must satisfy the following validator: %[2]s",
		"fr": "La valeur retournée par la fonction %[1]s doit satisfaire le validateur suivant: %[2]s",
	},
	"ParseValidator": {
		"en": "Parsing value with function %[1]s must succeed and the parsed value must satisfy the following validator: %[2]s",
		"fr": "Le parsing de la valeur avec la fonction %[1]s doit réussir et la valeur parsée doit satisfaire le validateur suivant: %[2]s",
	},
	"ParseError": {
		"en": "Parsing value %#[1]v with function %[2]s failed: %[3]s",
		"fr": "Le parsing de la valeur %#[1]v avec la fonction %[2]s a échoué: %[3]s",
	},
	"WrapValidator": {
		"en": "Value must satisfy the following validator, with error wrapped by message %[1]q: %[2]s",
		"fr": "La valeur doit satisfaire le validateur suivant, avec l'erreur enveloppée par le message %[1]q: %[2]s",
	},
	"FieldValidator": {
		"en": "Value of field %[1]q returned by function %[2]s must satisfy the following validator: %[3]s",
		"fr": "La valeur du champ %[1]q retournée par la fonction %[2]s doit satisfaire le validateur suivant: %[3]s",
	},
	"MessageValidator": {
		"en": "Value must satisfy the following validator, with error message overridden by %[1]q: %[2]s",
		"fr": "La valeur doit satisfaire le validateur suivant, avec le message d'erreur remplacé par %[1]q: %[2]s",
	},
	"ValidatorFunc": {
		"en": "Function %[1]s.",
		"fr": "Fonction %[1]s.",
	},
	"NoOpValidator": {
		"en": "No validation.",
		"fr": "Pas de validation.",
	},
	"TypeOptionalValidator": {
		"en": "Value must satisfy the following validator if it is of type %[1]T: %[2]s",
		"fr": "La valeur doit satisfaire le validateur suivant si elle est de type %[1]T: %[2]s",
	},
	"TypeRequiredValidator": {
		"en": "Value must be of type %[1]T and satisfy the following validator: %[2]s",
		"fr": "La valeur doit être de type %[1]T et satisfaire le validateur suivant: %[2]s",
	},
	"TypeRequiredError": {
		"en": "Type %[1]T cannot be converted to %[2]T.",
		"fr": "Le type %[1]T ne peut pas être converti en %[2]T.",
	},
	"ZeroValidator": {
		"en": "Value must be zero.",
		"fr": "La valeur doit être zéro.",
	},
	"ZeroError": {
		"en": "Value %#[1]v is not zero.",
		"fr": "La valeur %#[1]v n'est pas zéro.",
	},
	"NotZeroValidator": {
		"en": "Value must not be zero.",
		"fr": "La valeur ne doit pas être zéro.",
	},
	"NotZeroError": {
		"en": "Value is zero.",
		"fr": "La valeur est zéro.",
	},
	"OptionalValidator": {
		"en": "Value must be zero or satisfy the following validator: %[1]s",
		"fr": "La valeur doit être zéro ou satisfaire le validateur suivant: %[1]s",
	},
	"RequiredValidator": {
		"en": "Value must not be zero and must satisfy the following validator: %[1]s",
		"fr": "La valeur ne doit pas être zéro et doit satisfaire le validateur suivant: %[1]s",
	},
	"RequiredError": {
		"en": "Value is required.",
		"fr": "La valeur est requise.",
	},
	"PointerOptionalValidator": {
		"en": "Pointer must be nil or the dereferenced value must satisfy the following validator: %[1]s",
		"fr": "Le pointer doit être nil ou la valeur pointée doit satisfaire le validateur suivant: %[1]s",
	},
	"PointerRequiredValidator": {
		"en": "Pointer must not be nil and the dereferenced value must satisfy the following validator: %[1]s",
		"fr": "Le pointer ne doit pas être nil et la valeur pointée doit satisfaire le validateur suivant: %[1]s",
	},
	"PointerRequiredError": {
		"en": "Pointer is nil.",
		"fr": "Le pointeur est nil.", //nolint:misspell // "pointeur" is correct in French
	},
	"EqualValidator": {
		"en": "Value must be equal to %#[1]v.",
		"fr": "La valeur doit être égale à %#[1]v.",
	},
	"EqualFuncValidator": {
		"en": "Value must be equal to %#[1]v with function %[2]s.",
		"fr": "La valeur doit être égale à %#[1]v avec la fonction %[2]s.",
	},
	"EqualError": {
		"en": "Value %#[1]v is not equal to %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas égale à %#[2]v.",
	},
	"NotEqualValidator": {
		"en": "Value must not be equal to %#[1]v.",
		"fr": "La valeur ne doit pas être égale à %#[1]v.",
	},
	"NotEqualFuncValidator": {
		"en": "Value must not be equal to %#[1]v with function %[2]s.",
		"fr": "La valeur ne doit pas être égale à %#[1]v avec la fonction %[2]s.",
	},
	"NotEqualError": {
		"en": "Value %#[1]v is equal to %#[2]v.",
		"fr": "La valeur %#[1]v est égale à %#[2]v.",
	},
	"CmpEqualValidator": {
		"en": "Value must be equal to %#[1]v using Compare method.",
		"fr": "La valeur doit être égale à %#[1]v en utilisant la méthode Compare.",
	},
	"CmpNotEqualValidator": {
		"en": "Value must not be equal to %#[1]v using Compare method.",
		"fr": "La valeur ne doit pas être égale à %#[1]v en utilisant la méthode Compare.",
	},
	"CmpMinValidator": {
		"en": "Value must be greater than or equal to %#[1]v using Compare method.",
		"fr": "La valeur doit être supérieure ou égale à %#[1]v en utilisant la méthode Compare.",
	},
	"CmpMaxValidator": {
		"en": "Value must be less than or equal to %#[1]v using Compare method.",
		"fr": "La valeur doit être inférieure ou égale à %#[1]v en utilisant la méthode Compare.",
	},
	"CmpRangeValidator": {
		"en": "Value must be in the range [%#[1]v, %#[2]v] using Compare method.",
		"fr": "La valeur doit être dans l'intervalle [%#[1]v, %#[2]v] en utilisant la méthode Compare.",
	},
	"CmpLessValidator": {
		"en": "Value must be less than %#[1]v using Compare method.",
		"fr": "La valeur doit être inférieure à %#[1]v en utilisant la méthode Compare.",
	},
	"CmpGreaterValidator": {
		"en": "Value must be greater than %#[1]v using Compare method.",
		"fr": "La valeur doit être supérieure à %#[1]v en utilisant la méthode Compare.",
	},
	"InValidator": {
		"en": "Value must be in %#[1]v.",
		"fr": "La valeur doit être dans %#[1]v.",
	},
	"InError": {
		"en": "Value %#[1]v is not in %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas dans %#[2]v.",
	},
	"NotInValidator": {
		"en": "Value must not be in %#[1]v.",
		"fr": "La valeur ne doit pas être dans %#[1]v.",
	},
	"NotInError": {
		"en": "Value %#[1]v is in %#[2]v.",
		"fr": "La valeur %#[1]v est dans %#[2]v.",
	},
	"LenEqualValidator": {
		"en": "Length must be equal to %[1]d.",
		"fr": "La longueur doit être égale à %[1]d.",
	},
	"LenEqualError": {
		"en": "Length %[1]d is not equal to %[2]d.",
		"fr": "La longueur %[1]d n'est pas égale à %[2]d.",
	},
	"LenMinValidator": {
		"en": "Length must be greater than or equal to %[1]d.",
		"fr": "La longueur doit être supérieure ou égale à %[1]d.",
	},
	"LenMinError": {
		"en": "Length %[1]d is less than %[2]d.",
		"fr": "La longueur %[1]d est inférieure à %[2]d.",
	},
	"LenMaxValidator": {
		"en": "Length must be less than or equal to %[1]d.",
		"fr": "La longueur doit être inférieure ou égale à %[1]d.",
	},
	"LenMaxError": {
		"en": "Length %[1]d is greater than %[2]d.",
		"fr": "La longueur %[1]d est supérieure à %[2]d.",
	},
	"LenRangeValidator": {
		"en": "Length must be in the range [%[1]d, %[2]d].",
		"fr": "La longueur doit être dans l'intervalle [%[1]d, %[2]d].",
	},
	"LenRangeError": {
		"en": "Length %[1]d is not in the range [%[2]d, %[3]d].",
		"fr": "La longueur %[1]d n'est pas dans l'intervalle [%[2]d, %[3]d].",
	},
	"EmptyValidator": {
		"en": "Value must be empty.",
		"fr": "La valeur doit être vide.",
	},
	"EmptyError": {
		"en": "Value is not empty (%[1]d).",
		"fr": "La valeur n'est pas vide (%[1]d).",
	},
	"NotEmptyValidator": {
		"en": "Value must not be empty.",
		"fr": "La valeur ne doit pas être vide.",
	},
	"NotEmptyError": {
		"en": "Value is empty.",
		"fr": "La valeur est vide.",
	},
	"MinValidator": {
		"en": "Value must be greater than or equal to %#[1]v.",
		"fr": "La valeur doit être supérieure ou égale à %#[1]v.",
	},
	"MinFuncValidator": {
		"en": "Value must be greater than or equal to %#[1]v with function %[2]s.",
		"fr": "La valeur doit être supérieure ou égale à %#[1]v avec la fonction %[2]s.",
	},
	"MinError": {
		"en": "Value %#[1]v is less than %#[2]v.",
		"fr": "La valeur %#[1]v est inférieure à %#[2]v.",
	},
	"MaxValidator": {
		"en": "Value must be less than or equal to %#[1]v.",
		"fr": "La valeur doit être inférieure ou égale à %#[1]v.",
	},
	"MaxFuncValidator": {
		"en": "Value must be less than or equal to %#[1]v with function %[2]s.",
		"fr": "La valeur doit être inférieure ou égale à %#[1]v avec la fonction %[2]s.",
	},
	"MaxError": {
		"en": "Value %#[1]v is greater than %#[2]v.",
		"fr": "La valeur %#[1]v est supérieure à %#[2]v.",
	},
	"RangeValidator": {
		"en": "Value must be in the range [%#[1]v, %#[2]v].",
		"fr": "La valeur doit être dans l'intervalle [%#[1]v, %#[2]v].",
	},
	"RangeFuncValidator": {
		"en": "Value must be in the range [%#[1]v, %#[2]v] with function %[3]s.",
		"fr": "La valeur doit être dans l'intervalle [%#[1]v, %#[2]v] avec la fonction %[3]s.",
	},
	"RangeError": {
		"en": "Value %#[1]v is not in the range [%#[2]v, %#[3]v].",
		"fr": "La valeur %#[1]v n'est pas dans l'intervalle [%#[2]v, %#[3]v].",
	},
	"LessValidator": {
		"en": "Value must be less than %#[1]v.",
		"fr": "La valeur doit être inférieure à %#[1]v.",
	},
	"LessFuncValidator": {
		"en": "Value must be less than %#[1]v with function %[2]s.",
		"fr": "La valeur doit être inférieure à %#[1]v avec la fonction %[2]s.",
	},
	"LessError": {
		"en": "Value %#[1]v is not less than %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas inférieure à %#[2]v.",
	},
	"GreaterValidator": {
		"en": "Value must be greater than %#[1]v.",
		"fr": "La valeur doit être supérieure à %#[1]v.",
	},
	"GreaterFuncValidator": {
		"en": "Value must be greater than %#[1]v with function %[2]s.",
		"fr": "La valeur doit être supérieure à %#[1]v avec la fonction %[2]s.",
	},
	"GreaterError": {
		"en": "Value %#[1]v is not greater than %#[2]v.",
		"fr": "La valeur %#[1]v n'est pas supérieure à %#[2]v.",
	},
	"RegexpMatchValidator": {
		"en": "String must match regexp %[1]q.",
		"fr": "La chaîne doit correspondre à l'expression régulière %[1]q.",
	},
	"RegexpMatchError": {
		"en": "String %[1]q does not match regexp %[2]q.",
		"fr": "La chaîne %[1]q ne correspond pas à l'expression régulière %[2]q.",
	},
	"RegexpNotMatchValidator": {
		"en": "String must not match regexp %[1]q.",
		"fr": "La chaîne ne doit pas correspondre à l'expression régulière %[1]q.",
	},
	"RegexpNotMatchError": {
		"en": "String %[1]q matches regexp %[2]q.",
		"fr": "La chaîne %[1]q correspond à l'expression régulière %[2]q.",
	},
	"PositiveValidator": {
		"en": "Value must be positive.",
		"fr": "La valeur doit être positive.",
	},
	"PositiveError": {
		"en": "Value %#[1]v is not positive.",
		"fr": "La valeur %#[1]v n'est pas positive.",
	},
	"NegativeValidator": {
		"en": "Value must be negative.",
		"fr": "La valeur doit être négative.",
	},
	"NegativeError": {
		"en": "Value %#[1]v is not negative.",
		"fr": "La valeur %#[1]v n'est pas négative.",
	},
	"SliceContainsValidator": {
		"en": "Slice must contain %#[1]v.",
		"fr": "Le slice doit contenir %#[1]v.",
	},
	"SliceContainsError": {
		"en": "Slice does not contain %#[1]v.",
		"fr": "Le slice ne contient pas %#[1]v.",
	},
	"SliceNotContainsValidator": {
		"en": "Slice must not contain %#[1]v.",
		"fr": "Le slice ne doit pas contenir %#[1]v.",
	},
	"SliceNotContainsError": {
		"en": "Slice contains %#[1]v.",
		"fr": "Le slice contient %#[1]v.",
	},
	"SliceEachValidator": {
		"en": "Each index/element of the slice must satisfy the following validator: %[1]s",
		"fr": "Chaque index/élément du slice doit satisfaire le validateur suivant: %[1]s",
	},
	"SliceEachValueValidator": {
		"en": "Each element of the slice must satisfy the following validator: %[1]s",
		"fr": "Chaque élément du slice doit satisfaire le validateur suivant: %[1]s",
	},
	"SliceUniqueValidator": {
		"en": "Slice must have unique elements.",
		"fr": "Le slice doit avoir des éléments uniques.",
	},
	"SliceUniqueByValidator": {
		"en": "Slice must have unique elements with function %[1]s.",
		"fr": "Le slice doit avoir des éléments uniques avec la fonction %[1]s.",
	},
	"SliceUniqueError": {
		"en": "Duplicate %#[1]v (index %[2]d).",
		"fr": "Doublon %#[1]v (indice %[2]d).",
	},
	"MapEachValidator": {
		"en": "Each key/value of the map must satisfy the following validator: %[1]s",
		"fr": "Chaque clé/valeur de la map doit satisfaire le validateur suivant: %[1]s",
	},
	"MapEachKeyValidator": {
		"en": "Each key of the map must satisfy the following validator: %[1]s",
		"fr": "Chaque clé de la map doit satisfaire le validateur suivant: %[1]s",
	},
	"MapEachValueValidator": {
		"en": "Each value of the map must satisfy the following validator: %[1]s",
		"fr": "Chaque valeur de la map doit satisfaire le validateur suivant: %[1]s",
	},
	"MapSortedEachValidator": {
		"en": "Each key/value of the map must satisfy the following validator (in sorted key order): %[1]s",
		"fr": "Chaque clé/valeur de la map doit satisfaire le validateur suivant (dans l'ordre croissant des clés): %[1]s",
	},
	"MapSortedEachKeyValidator": {
		"en": "Each key of the map must satisfy the following validator (in sorted key order): %[1]s",
		"fr": "Chaque clé de la map doit satisfaire le validateur suivant (dans l'ordre croissant des clés): %[1]s",
	},
	"MapSortedEachValueValidator": {
		"en": "Each value of the map must satisfy the following validator (in sorted key order): %[1]s",
		"fr": "Chaque valeur de la map doit satisfaire le validateur suivant (dans l'ordre croissant des clés): %[1]s",
	},
	"SeqEachValidator": {
		"en": "Each index/element of the sequence must satisfy the following validator: %[1]s",
		"fr": "Chaque index/élément de la séquence doit satisfaire le validateur suivant: %[1]s",
	},
	"SeqEachValueValidator": {
		"en": "Each element of the sequence must satisfy the following validator: %[1]s",
		"fr": "Chaque élément de la séquence doit satisfaire le validateur suivant: %[1]s",
	},
	"Seq2EachValidator": {
		"en": "Each key/value of the sequence must satisfy the following validator: %[1]s",
		"fr": "Chaque clé/valeur de la séquence doit satisfaire le validateur suivant: %[1]s",
	},
	"Seq2EachKeyValidator": {
		"en": "Each key of the sequence must satisfy the following validator: %[1]s",
		"fr": "Chaque clé de la séquence doit satisfaire le validateur suivant: %[1]s",
	},
	"Seq2EachValueValidator": {
		"en": "Each value of the sequence must satisfy the following validator: %[1]s",
		"fr": "Chaque valeur de la séquence doit satisfaire le validateur suivant: %[1]s",
	},
	"StringContainsValidator": {
		"en": "String must contain %[1]q.",
		"fr": "La chaîne doit contenir %[1]q.",
	},
	"StringContainsError": {
		"en": "String %[1]q does not contain %[2]q.",
		"fr": "La chaîne %[1]q ne contient pas %[2]q.",
	},
	"StringNotContainsValidator": {
		"en": "String must not contain %[1]q.",
		"fr": "La chaîne ne doit pas contenir %[1]q.",
	},
	"StringNotContainsError": {
		"en": "String %[1]q contains %[2]q.",
		"fr": "La chaîne %[1]q contient %[2]q.",
	},
	"StringHasPrefixValidator": {
		"en": "String must begin with %[1]q.",
		"fr": "La chaîne doit commencer par %[1]q.",
	},
	"StringHasPrefixError": {
		"en": "String %[1]q does not begin with %[2]q.",
		"fr": "La chaîne %[1]q ne commence pas par %[2]q.",
	},
	"StringNotHasPrefixValidator": {
		"en": "String must not begin with %[1]q.",
		"fr": "La chaîne ne doit pas commencer par %[1]q.",
	},
	"StringNotHasPrefixError": {
		"en": "String %[1]q begins with %[2]q.",
		"fr": "La chaîne %[1]q commence par %[2]q.",
	},
	"StringHasSuffixValidator": {
		"en": "String must end with %[1]q.",
		"fr": "La chaîne doit se terminer par %[1]q.",
	},
	"StringHasSuffixError": {
		"en": "String %[1]q does not end with %[2]q.",
		"fr": "La chaîne %[1]q ne se termine pas par %[2]q.",
	},
	"StringNotHasSuffixValidator": {
		"en": "String must not end with %[1]q.",
		"fr": "La chaîne ne doit pas se terminer par %[1]q.",
	},
	"StringNotHasSuffixError": {
		"en": "String %[1]q ends with %[2]q.",
		"fr": "La chaîne %[1]q se termine par %[2]q.",
	},
	"BytesEqualValidator": {
		"en": "Bytes must be equal to %[1]q.",
		"fr": "Les bytes doivent être égaux à %[1]q.",
	},
	"BytesEqualError": {
		"en": "Bytes %[1]q are not equal to %[2]q.",
		"fr": "Les bytes %[1]q ne sont pas égaux à %[2]q.",
	},
	"BytesNotEqualValidator": {
		"en": "Bytes must not be equal to %[1]q.",
		"fr": "Les bytes ne doivent pas être égaux à %[1]q.",
	},
	"BytesNotEqualError": {
		"en": "Bytes %[1]q are equal to %[2]q.",
		"fr": "Les bytes %[1]q sont égaux à %[2]q.",
	},
	"BytesContainsValidator": {
		"en": "Bytes must contain %[1]q.",
		"fr": "Les bytes doivent contenir %[1]q.",
	},
	"BytesContainsError": {
		"en": "Bytes %[1]q do not contain %[2]q.",
		"fr": "Les bytes %[1]q ne contiennent pas %[2]q.",
	},
	"BytesNotContainsValidator": {
		"en": "Bytes must not contain %[1]q.",
		"fr": "Les bytes ne doivent pas contenir %[1]q.",
	},
	"BytesNotContainsError": {
		"en": "Bytes %[1]q contain %[2]q.",
		"fr": "Les bytes %[1]q contiennent %[2]q.",
	},
	"BytesHasPrefixValidator": {
		"en": "Bytes must have prefix %[1]q.",
		"fr": "Les bytes doivent avoir le préfixe %[1]q.",
	},
	"BytesHasPrefixError": {
		"en": "Bytes %[1]q do not have prefix %[2]q.",
		"fr": "Les bytes %[1]q n'ont pas le préfixe %[2]q.",
	},
	"BytesNotHasPrefixValidator": {
		"en": "Bytes must not have prefix %[1]q.",
		"fr": "Les bytes ne doivent pas avoir le préfixe %[1]q.",
	},
	"BytesNotHasPrefixError": {
		"en": "Bytes %[1]q have prefix %[2]q.",
		"fr": "Les bytes %[1]q ont le préfixe %[2]q.",
	},
	"BytesHasSuffixValidator": {
		"en": "Bytes must have suffix %[1]q.",
		"fr": "Les bytes doivent avoir le suffixe %[1]q.",
	},
	"BytesHasSuffixError": {
		"en": "Bytes %[1]q do not have suffix %[2]q.",
		"fr": "Les bytes %[1]q n'ont pas le suffixe %[2]q.",
	},
	"BytesNotHasSuffixValidator": {
		"en": "Bytes must not have suffix %[1]q.",
		"fr": "Les bytes ne doivent pas avoir le suffixe %[1]q.",
	},
	"BytesNotHasSuffixError": {
		"en": "Bytes %[1]q have suffix %[2]q.",
		"fr": "Les bytes %[1]q ont le suffixe %[2]q.",
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
	for i, a := range args {
		s, ok := localizeValue(a, locales...)
		if ok {
			args[i] = s
		}
	}
	return fmt.Sprintf(format, args...)
}

func localizeValue(v any, locales ...string) (string, bool) {
	switch v := v.(type) {
	case Localizer:
		return v.Localize(locales...), true
	case Localizable:
		return LocalizeLocalizable(v, locales...), true
	}
	return "", false
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
	s, _ := localizeValue(vr, locales...)
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
