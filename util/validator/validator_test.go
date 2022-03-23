package validator

import (
	"github.com/stretchr/testify/assert"
	"go-vote/util/test"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tcs := []test.Case{
		{
			Name:     "Complete Email Expect True",
			Input:    "test@mail.com",
			Expected: true,
			Error:    false,
		},
		{
			Name:     "Name Only Expect False",
			Input:    "test",
			Expected: false,
			Error:    false,
		},
		{
			Name:     "Domain Only Expect False",
			Input:    "mail.com",
			Expected: false,
			Error:    false,
		},
		{
			Name:     "Name With @ Only Expect False",
			Input:    "test@",
			Expected: false,
			Error:    false,
		},
		{
			Name:     "Domain With @ Only Expect False",
			Input:    "@mail.com",
			Expected: false,
			Error:    false,
		},
		{
			Name:     "Empty Expect False",
			Input:    "",
			Expected: false,
			Error:    false,
		},
	}

	handler := func(input string, expected bool) {
		actual, _ := ValidateEmail(input)
		assert.Equal(t, expected, actual)
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(tc.Input.(string), tc.Expected.(bool))
		})
	}
}
