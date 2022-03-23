package model

import (
	"errors"
	"time"
)

type Auth struct {
	Id        int64  `json:"id"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int64  `json:"createdAt"`
	ExpiredAt int64  `json:"expiredAt"`
}

func (a *Auth) Valid() error {
	if a.Id < 1 {
		return errors.New("invalid user id")
	}
	if a.ExpiredAt <= time.Now().Unix() {
		return errors.New("auth session expired")
	}
	return nil
}

func (a *Auth) ToUser() User {
	return User{
		Id:       a.Id,
		Name:     a.Name,
		Email:    a.Email,
		Password: "",
	}
}
