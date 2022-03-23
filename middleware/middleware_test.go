package middleware

import (
	"github.com/stretchr/testify/assert"
	"go-vote/config"
	"go-vote/util/test"
	"testing"
)

func TestGetAccessConfig(t *testing.T) {
	test.Init()
	cfg := config.Get()
	actual := GetAccessConfig()
	assert.Equal(t, []byte(cfg.AccessSecret), actual.SigningKey)
}

func TestGetRefreshConfig(t *testing.T) {
	test.Init()
	cfg := config.Get()
	actual := GetRefreshConfig()
	assert.Equal(t, []byte(cfg.RefreshSecret), actual.SigningKey)
}
