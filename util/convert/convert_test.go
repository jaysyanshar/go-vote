package convert

import (
	"github.com/stretchr/testify/assert"
	"go-vote/util/test"
	"testing"
)

func TestStrToInt64(t *testing.T) {
	tcs := []test.Case{
		{
			Name:     "Convert Number Base 10 Expect Success",
			Input:    "100",
			Expected: int64(100),
			Error:    false,
		},
		{
			Name:     "Convert Number Random Expect Success",
			Input:    "1230918",
			Expected: int64(1230918),
			Error:    false,
		},
		{
			Name:     "Convert Big Number Expect Success",
			Input:    "9223372036854775807",
			Expected: int64(9223372036854775807),
			Error:    false,
		},
		{
			Name:     "Convert Text Expect Error",
			Input:    "this is not a number",
			Expected: int64(0),
			Error:    true,
		},
	}

	handler := func(input string, expected int64, expectError bool) {
		actual, err := StrToInt64(input)
		if err != nil {
			assert.True(t, expectError)
			return
		}
		assert.Equal(t, expected, actual)
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(tc.Input.(string), tc.Expected.(int64), tc.Error)
		})
	}
}
