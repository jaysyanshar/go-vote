package model

import (
	"database/sql"
	"go-vote/util/response"
	"go-vote/util/validator"
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

type LoginUserReq struct {
	Email    string
	Password string
}

type RegisterUserRes struct {
	response.Response `json:"-"`
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
}

type GetProfileUserRes struct {
	response.Response `json:"-"`
	Name              string `json:"name"`
	Email             string `json:"email"`
}

type LoginUserRes struct {
	response.Response `json:"-"`
	AccessToken       string `json:"accessToken"`
	RefreshToken      string `json:"refreshToken"` //todo: implement refresh token
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
	return validateEmailPassword(r.Email, r.Password)
}

func (r *LoginUserReq) Validate() (bool, error) {
	return validateEmailPassword(r.Email, r.Password)
}

func (d *FindUserDb) ToUser() User {
	return User{
		Id:       d.Id,
		Name:     d.Name.String,
		Email:    d.Email,
		Password: d.Password,
	}
}

func validateEmailPassword(email, password string) (bool, error) {
	valid, err := validator.ValidateEmail(email)
	if !valid {
		return false, err
	}
	valid, err = validator.ValidatePassword(password)
	return valid, err
}
