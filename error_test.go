package trnmnterr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorf(t *testing.T) {
	tt := []struct {
		name         string
		startErr     error
		isTrnmtError bool
		format       string
		arg          []interface{}
		want         string
	}{
		{
			"Tournament error",
			New(errors.New("test"), NotFound),
			true,
			"test2",
			[]interface{}{},
			"test2: test",
		},
		{
			"Default Error",
			errors.New("test3"),
			false,
			"db %v",
			[]interface{}{"postgres"},
			"db postgres: test3",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			got := Errorf(tc.startErr, tc.format, tc.arg...)
			trmt := trnmntError{}
			assert.Equal(errors.As(got, &trmt), tc.isTrnmtError, "assert tournament and default error")
			assert.Equal(got.Error(), tc.want, "error message should be equal")
		})

	}

}
