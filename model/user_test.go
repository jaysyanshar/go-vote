package model

import (
	"github.com/stretchr/testify/assert"
	"go-vote/util/test"
	"testing"
)

func TestLoginUserReq_Validate(t *testing.T) {
	tcs := []test.Case{
		{
			Name: "Expect Success",
			Input: &RegisterUserReq{
				Name:     "Jays",
				Email:    "jays@mail.com",
				Password: "123456",
			},
			Expected: true,
			Error:    false,
		},
		{
			Name: "Expect Error Password",
			Input: &RegisterUserReq{
				Name:     "Jays",
				Email:    "jays@mail.com",
				Password: "",
			},
			Expected: false,
			Error:    false,
		},
	}

	handler := func(input *RegisterUserReq, expected bool) {
		actual, _ := input.Validate()
		assert.Equal(t, expected, actual)
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(tc.Input.(*RegisterUserReq), tc.Expected.(bool))
		})
	}
}
