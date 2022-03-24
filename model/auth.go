package model

import (
	"errors"
	"time"
)

type Auth struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"userId"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	IpAddress string `json:"ipAddress"`
	CreatedAt int64  `json:"createdAt"`
	ExpiredAt int64  `json:"expiredAt"`
	IsRevoked bool   `json:"isRevoked"`
}

type InsertAuthDb struct {
	UserId    int64
	IpAddress string
	CreatedAt string
	ExpiredAt string
	IsRevoked bool
}

type UpdateAuthDb struct {
	Id        int64
	ExpiredAt string
}

type FindAuthDb struct {
	Id        int64
	UserId    int64
	Email     int64
	Name      int64
	IpAddress string
	CreatedAt string
	ExpiredAt string
	IsRevoked bool
}

func (a *Auth) Valid() error {
	if a.UserId < 1 {
		return errors.New("invalid user id")
	}
	if a.ExpiredAt <= time.Now().Unix() {
		return errors.New("auth session expired")
	}
	if a.IsRevoked {
		return errors.New("auth session already invoked")
	}
	return nil
}

func (a *Auth) ToUser() User {
	return User{
		Id:       a.UserId,
		Name:     a.Name,
		Email:    a.Email,
		Password: "",
	}
}
