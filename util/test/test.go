package test

import (
	"github.com/stretchr/testify/assert"
	"go-vote/config"
	"testing"
)

type Case struct {
	Name     string
	Input    interface{}
	Expected interface{}
	Error    bool
}

var cfg = &config.Config{
	ServicePort:             0,
	AccessSecret:            "123456",
	TokenExpirationMinute:   1,
	DbDriver:                "",
	DbSource:                "",
	DbConnMaxLifetimeMinute: 0,
	DbMaxOpenIdleConn:       0,
	RedisAddress:            "",
	RedisPassword:           "",
	RedisDb:                 0,
}

func Init() {
	config.Set(cfg)
}

func HandlePanic(t *testing.T, expected bool) {
	if r := recover(); r != nil {
		assert.True(t, expected)
	}
}
