package jwt

import (
	"github.com/stretchr/testify/assert"
	"go-vote/model"
	"go-vote/util/test"
	"testing"
)

func TestCreateToken(t *testing.T) {
	handler := func(t *testing.T, c test.Case) {
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
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(t, tc)
		})
	}
}
