package vld

// And creates a [AndValidator].
func And[T any](vrs ...Validator[T]) *AndValidator[T] {
	return &AndValidator[T]{
		Validators: vrs,
	}
}

// AndValidator is a [Validator] that validates the value with all validators and returns the first error.
type AndValidator[T any] struct {
	Validators []Validator[T]
}

// Validate implements [Validator].
func (vr *AndValidator[T]) Validate(v T) error {
	for _, validator := range vr.Validators {
		err := validator.Validate(v)
		if err != nil {
			return err //nolint:wrapcheck // Not needed.
		}
	}
	return nil
}

func (vr *AndValidator[T]) String() string {
	return buildMultiValidatorString("And", vr.Validators...)
}

// Localize implements [Localizer].
func (vr *AndValidator[T]) Localize(locales ...string) string {
	return localizeMultiValidator("And", vr.Validators, locales...)
}

// Or creates a [OrValidator].
func Or[T any](vrs ...Validator[T]) *OrValidator[T] {
	return &OrValidator[T]{
		Validators: vrs,
	}
}

// OrValidator is a [Validator] that validates the value with all validators and returns nil if any returns nil, or joins all errors otherwise.
type OrValidator[T any] struct {
	Validators []Validator[T]
}

// Validate implements [Validator].
func (vr *OrValidator[T]) Validate(v T) error {
	var errs []error
	for _, validator := range vr.Validators {
		err := validator.Validate(v)
		if err == nil {
			return nil
		}
		errs = append(errs, err)
	}
	return ErrorJoin(errs...)
}

func (vr *OrValidator[T]) String() string {
	return buildMultiValidatorString("Or", vr.Validators...)
}

// Localize implements [Localizer].
func (vr *OrValidator[T]) Localize(locales ...string) string {
	return localizeMultiValidator("Or", vr.Validators, locales...)
}

// All creates an [AllValidator].
func All[T any](vrs ...Validator[T]) *AllValidator[T] {
	return &AllValidator[T]{
		Validators: vrs,
	}
}

// AllValidator is a [Validator] that validates the value with all validators and joins all errors.
type AllValidator[T any] struct {
	Validators []Validator[T]
}

// Validate implements [Validator].
func (vr *AllValidator[T]) Validate(v T) error {
	var errs []error
	for _, validator := range vr.Validators {
		err := validator.Validate(v)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return ErrorJoin(errs...)
}

func (vr *AllValidator[T]) String() string {
	return buildMultiValidatorString("All", vr.Validators...)
}

// Localize implements [Localizer].
func (vr *AllValidator[T]) Localize(locales ...string) string {
	return localizeMultiValidator("All", vr.Validators, locales...)
}
