package middleware

import (
	"github.com/labstack/echo/v4/middleware"
	"go-vote/config"
	"go-vote/model"
	"go-vote/util/constant"
)

func GetAccessConfig() middleware.JWTConfig {
	cfg := config.Get()
	return middleware.JWTConfig{
		SigningKey: []byte(cfg.AccessSecret),
		Claims:     &model.Auth{},
	}
}

func GetRefreshConfig() middleware.JWTConfig {
	cfg := config.Get()
	return middleware.JWTConfig{
		SigningKey:  []byte(cfg.RefreshSecret),
		TokenLookup: "cookie:" + constant.HttpCookieRefreshToken,
		Claims:      &model.Auth{},
	}
}
