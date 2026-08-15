package vld_test

import (
	"testing"
	"time"

	"github.com/pierrre/assert"
	. "github.com/pierrre/vld"
)

func TestTime(t *testing.T) {
	vr := Time(time.RFC3339, nil, nil)
	testValidator(t, vr, "2006-01-02T15:04:05Z", "invalid")
	err := vr.Validate("invalid")
	terr, _ := assert.ErrorAsType[*TimeError](t, err)
	err = terr.Unwrap()
	assert.Error(t, err)
}

func TestTimeInLocation(t *testing.T) {
	testValidator(t, Time(time.RFC3339, time.UTC, nil), "2006-01-02T15:04:05Z", "invalid")
}

func TestTimeWithValidator(t *testing.T) {
	testValidator(t, Time(time.RFC3339, nil, NotZero[time.Time]()), "2006-01-02T15:04:05Z", "0001-01-01T00:00:00Z", "invalid")
}

func TestTimeInLocationWithValidator(t *testing.T) {
	testValidator(t, Time(time.RFC3339, time.UTC, NotZero[time.Time]()), "2006-01-02T15:04:05Z", "0001-01-01T00:00:00Z", "invalid")
}
