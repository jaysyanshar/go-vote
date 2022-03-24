package jwt

import (
	"github.com/stretchr/testify/assert"
	"go-vote/config"
	"go-vote/model"
	"go-vote/util/test"
	"testing"
)

func TestCreateToken(t *testing.T) {
	test.Init()
	ip := "::1"
	handler := func(t *testing.T, c test.Case) {
		defer test.HandlePanic(t, c.Error)
		actual, err := CreateToken(c.Input.(model.User), ip)
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
	ip := "::1"
	handler := func(t *testing.T, c test.Case) {
		defer test.HandlePanic(t, c.Error)
		actual, err := CreateRefreshToken(c.Input.(model.User), ip)
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

func TestParseToken(t *testing.T) {
	test.Init()
	cfg := config.Get()
	ip := "::1"

	handler := func(t *testing.T, c test.Case) {
		defer test.HandlePanic(t, c.Error)
		input := c.Input.(model.User)
		token, _ := CreateToken(input, ip)
		actual, err := ParseToken(token, cfg.AccessSecret)
		if err != nil {
			assert.True(t, c.Error)
			return
		}
		assert.Equal(t, input.Id, actual.Id)
		assert.Equal(t, input.Name, actual.Name)
		assert.Equal(t, input.Email, actual.Email)
		assert.Greater(t, actual.CreatedAt, int64(0))
		assert.Greater(t, actual.ExpiredAt, int64(0))
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
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(t, tc)
		})
	}
}
