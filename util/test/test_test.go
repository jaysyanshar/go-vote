package test

import (
	"github.com/stretchr/testify/assert"
	"go-vote/config"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	t.Run("Init Test Expect Created", func(t *testing.T) {
		assert.Equal(t, cfg, config.Get())
	})
}
