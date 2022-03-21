package crypto

import (
	"github.com/stretchr/testify/assert"
	"go-vote/util/test"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tcs := []test.Case{
		{
			Name:     "Expect True Number",
			Input:    "123456",
			Expected: "123456",
			Error:    false,
		},
		{
			Name:     "Expect True Alphanumeric",
			Input:    "test123",
			Expected: "test123",
			Error:    false,
		},
		{
			Name:     "Expect True Empty",
			Input:    "",
			Expected: "",
			Error:    false,
		},
	}

	handler := func(input, expected string) {
		actual, _ := HashPassword(input)
		assert.True(t, CheckPasswordHash(expected, actual))
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(tc.Input.(string), tc.Expected.(string))
		})
	}
}
