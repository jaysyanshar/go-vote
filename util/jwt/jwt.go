package jwt

import (
	"github.com/golang-jwt/jwt"
	"go-vote/config"
	"go-vote/model"
	"time"
)

type Identity struct {
	AuthId    int64
	UserId    int64
	UserEmail string
	UserName  string
	IpAddress string
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

func GetRefreshTokenExpiration() string {
	return time.Unix(getRefreshTokenExpiration(), 0).String()
}

func CreateToken(i Identity) (string, error) {
	cfg := config.Get()
	return createToken(i, cfg.AccessSecret, getAccessTokenExpiration())
}

func CreateRefreshToken(i Identity) (string, error) {
	cfg := config.Get()
	return createToken(i, cfg.RefreshSecret, getRefreshTokenExpiration())
}

func createToken(i Identity, secret string, expiredAt int64) (string, error) {
	claims := &model.Auth{
		UserId:    i.UserId,
		Email:     i.UserEmail,
		Name:      i.UserName,
		IpAddress: i.IpAddress,
		CreatedAt: time.Now().Unix(),
		ExpiredAt: expiredAt,
		IsRevoked: false,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getAccessTokenExpiration() int64 {
	cfg := config.Get()
	return time.Now().Add(time.Minute * time.Duration(cfg.TokenExpirationMinute)).Unix()
}

func getRefreshTokenExpiration() int64 {
	cfg := config.Get()
	return time.Now().Add(24 * time.Hour * time.Duration(cfg.RefreshTokenExpirationDay)).Unix()
}
