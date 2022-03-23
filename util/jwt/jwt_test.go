package jwt

import (
	"github.com/stretchr/testify/assert"
	"go-vote/model"
	"go-vote/util/test"
	"testing"
)

func TestCreateToken(t *testing.T) {
	test.Init()
	handler := func(t *testing.T, c test.Case) {
		defer test.HandlePanic(t, c.Error)
		actual, err := CreateToken(c.Input.(model.User))
		if err != nil {
			assert.True(t, c.Error)
			return
		}
		assert.True(t, actual != "")
	}

	tcs := []test.Case{
		{
			Name: "Expect Success",
			Input: model.User{
				Id:       1,
				Name:     "User",
				Email:    "user@mail.com",
				Password: "",
			},
			Expected: nil,
			Error:    false,
		},
		{
			Name:     "Expect Error",
			Input:    nil,
			Expected: nil,
			Error:    true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(t, tc)
		})
	}
}

func TestCreateRefreshToken(t *testing.T) {
	test.Init()
	handler := func(t *testing.T, c test.Case) {
		defer test.HandlePanic(t, c.Error)
		actual, err := CreateRefreshToken(c.Input.(model.User))
		if err != nil {
			assert.True(t, c.Error)
			return
		}
		assert.True(t, actual != "")
	}

	tcs := []test.Case{
		{
			Name: "Expect Success",
			Input: model.User{
				Id:       1,
				Name:     "User",
				Email:    "user@mail.com",
				Password: "",
			},
			Expected: nil,
			Error:    false,
		},
		{
			Name:     "Expect Error",
			Input:    nil,
			Expected: nil,
			Error:    true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(t, tc)
		})
	}
}
