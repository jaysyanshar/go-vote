package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"go-vote/config"
	"go-vote/model"
	"go-vote/util/constant"
	"time"
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

func ParseToken(token, secret string) (*model.Auth, error) {
	claims := &model.Auth{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func CreateToken(user model.User) (string, error) {
	cfg := config.Get()
	return createToken(user, cfg.AccessSecret, time.Now().Add(time.Minute*time.Duration(cfg.TokenExpirationMinute)).Unix())
}

func CreateRefreshToken(user model.User) (string, error) {
	cfg := config.Get()
	return createToken(user, cfg.RefreshSecret, time.Now().Add(24*time.Hour*time.Duration(cfg.RefreshTokenExpirationDay)).Unix())
}

func createToken(user model.User, secret string, expiredAt int64) (string, error) {
	claims := &model.Auth{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: time.Now().Unix(),
		ExpiredAt: expiredAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
