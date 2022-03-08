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

	//Creating Access AccessToken
	claims := jwt.MapClaims{}
	claims[KeyUserId] = user.Id
	claims[KeyUserName] = user.Name
	claims[KeyUserEmail] = user.Email
	claims[KeyExpiredAt] = time.Now().Add(time.Minute * time.Duration(cfg.TokenExpirationMinute)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.AccessSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
