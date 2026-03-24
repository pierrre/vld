package vld

import (
	"bytes"
	"fmt"
)

// BytesEqual returns a [Validator] that checks if the byte slice is equal to the specified byte slice.
func BytesEqual(s []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesEqual(%q)", s) }, func(b []byte) error {
		if !bytes.Equal(b, s) {
			err := fmt.Errorf("%q is not equal to %q", b, s)
			err = ErrorWrapLocalization(err, "BytesEqual", b, s)
			return err
		}
		return nil
	})
}

// BytesNotEqual returns a [Validator] that checks if the byte slice is not equal to the specified byte slice.
func BytesNotEqual(s []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesNotEqual(%q)", s) }, func(b []byte) error {
		if bytes.Equal(b, s) {
			err := fmt.Errorf("%q is equal to %q", b, s)
			err = ErrorWrapLocalization(err, "BytesNotEqual", b, s)
			return err
		}
		return nil
	})
}

// BytesContains returns a [Validator] that checks if the byte slice contains the specified byte slice.
func BytesContains(sub []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesContains(%q)", sub) }, func(b []byte) error {
		if !bytes.Contains(b, sub) {
			err := fmt.Errorf("%q does not contain %q", b, sub)
			err = ErrorWrapLocalization(err, "BytesContains", b, sub)
			return err
		}
		return nil
	})
}

// BytesNotContains returns a [Validator] that checks if the byte slice does not contain the specified byte slice.
func BytesNotContains(sub []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesNotContains(%q)", sub) }, func(b []byte) error {
		if bytes.Contains(b, sub) {
			err := fmt.Errorf("%q contains %q", b, sub)
			err = ErrorWrapLocalization(err, "BytesNotContains", b, sub)
			return err
		}
		return nil
	})
}

// BytesHasPrefix returns a [Validator] that checks if the byte slice has the specified prefix.
func BytesHasPrefix(prefix []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesHasPrefix(%q)", prefix) }, func(b []byte) error {
		if !bytes.HasPrefix(b, prefix) {
			err := fmt.Errorf("%q does not have prefix %q", b, prefix)
			err = ErrorWrapLocalization(err, "BytesHasPrefix", b, prefix)
			return err
		}
		return nil
	})
}

// BytesNotHasPrefix returns a [Validator] that checks if the byte slice does not have the specified prefix.
func BytesNotHasPrefix(prefix []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesNotHasPrefix(%q)", prefix) }, func(b []byte) error {
		if bytes.HasPrefix(b, prefix) {
			err := fmt.Errorf("%q has prefix %q", b, prefix)
			err = ErrorWrapLocalization(err, "BytesNotHasPrefix", b, prefix)
			return err
		}
		return nil
	})
}

// BytesHasSuffix returns a [Validator] that checks if the byte slice has the specified suffix.
func BytesHasSuffix(suffix []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesHasSuffix(%q)", suffix) }, func(b []byte) error {
		if !bytes.HasSuffix(b, suffix) {
			err := fmt.Errorf("%q does not have suffix %q", b, suffix)
			err = ErrorWrapLocalization(err, "BytesHasSuffix", b, suffix)
			return err
		}
		return nil
	})
}

// BytesNotHasSuffix returns a [Validator] that checks if the byte slice does not have the specified suffix.
func BytesNotHasSuffix(suffix []byte) Validator[[]byte] {
	return WithStringFunc(func() string { return fmt.Sprintf("BytesNotHasSuffix(%q)", suffix) }, func(b []byte) error {
		if bytes.HasSuffix(b, suffix) {
			err := fmt.Errorf("%q has suffix %q", b, suffix)
			err = ErrorWrapLocalization(err, "BytesNotHasSuffix", b, suffix)
			return err
		}
		return nil
	})
}
