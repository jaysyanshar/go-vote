package model

import (
	"database/sql"
	"errors"
	"net/mail"
)

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

type RegisterUserReq struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type RegisterUserRes struct {
	Response `json:"-"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type InsertUserDb struct {
	Name     sql.NullString
	Email    string
	Password string
}

type FindUserDb struct {
	Id       int64
	Name     sql.NullString
	Email    string
	Password string
}

func (r *RegisterUserReq) Validate() (bool, error) {
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return false, err
	}
	if r.Password == "" {
		return false, errors.New("password cannot be empty")
	}
	return true, nil
}
