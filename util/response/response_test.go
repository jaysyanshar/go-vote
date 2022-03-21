package response

import (
	"github.com/stretchr/testify/assert"
	"go-vote/util/test"
	"net/http"
	"testing"
)

func TestMakeResponse(t *testing.T) {
	tcs := []test.Case{
		{
			Name:     "Make Response Expect Status OK",
			Input:    http.StatusOK,
			Expected: Response{Status: http.StatusOK},
			Error:    false,
		},
	}

	handler := func(t *testing.T, tc test.Case) {
		actual := MakeResponse(tc.Input.(int))
		assert.Equal(t, tc.Expected.(Response), actual)
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			handler(t, tc)
		})
	}
}
