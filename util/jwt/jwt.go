package jwt

import (
	"github.com/golang-jwt/jwt"
	"go-vote/config"
	"go-vote/model"
	"time"
)

const (
	KeyUserId    = "userId"
	KeyUserName  = "userName"
	KeyUserEmail = "userEmail"
	KeyExpiredAt = "expiredAt"
)

func CreateToken(user model.User) (string, error) {
	cfg := config.Get()

	claims := jwt.MapClaims{}
	claims[KeyUserId] = user.Id
	claims[KeyUserName] = user.Name
	claims[KeyUserEmail] = user.Email
	claims[KeyExpiredAt] = time.Now().Add(time.Minute * time.Duration(cfg.TokenExpirationMinute)).Unix()

	return createToken(claims, cfg.AccessSecret)
}

func CreateRefreshToken(user model.User) (string, error) {
	cfg := config.Get()

	claims := jwt.MapClaims{}
	claims[KeyUserId] = user.Id
	claims[KeyExpiredAt] = time.Now().Add(24 * time.Hour * time.Duration(cfg.RefreshTokenExpirationDay)).Unix()

	return createToken(claims, cfg.RefreshSecret)
}

func createToken(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
