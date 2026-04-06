package vld_test

import (
	"fmt"
	"testing"

	"github.com/pierrre/assert/assertauto"
	. "github.com/pierrre/vld"
)

func TestValidatorFunc(t *testing.T) {
	testValidator(t, ValidatorFunc[int](Equal(1).Validate), 1, 2)
}

func TestNoOp(t *testing.T) {
	testValidator(t, NoOp[int](), 1)
}

var testLocales = []string{"en", "fr"}

func testValidator[T any](t *testing.T, vr Validator[T], vs ...T) {
	t.Helper()
	assertAutoString(t, vr.String())
	for _, l := range testLocales {
		assertAutoString(t, fmt.Sprintf("%s: %s", l, LocalizeValidator(vr, l)))
	}
	for _, v := range vs {
		assertauto.Equal(t, v)
		assertauto.AllocsPerRun(t, 100, func() {
			_ = vr.Validate(v)
		})
		err := vr.Validate(v)
		if err == nil {
			assertAutoString(t, "valid")
			continue
		}
		assertAutoString(t, err.Error())
		for _, l := range testLocales {
			assertAutoString(t, fmt.Sprintf("%s: %s", l, LocalizeError(err, l)))
		}
	}
}

func assertAutoString(t *testing.T, s string) {
	t.Helper()
	assertauto.Equal(t, s, assertauto.ValueStringer(func(a any) string {
		s, _ := a.(string)
		return s
	}))
}
