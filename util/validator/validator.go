package validator

import (
	"errors"
	"net/mail"
)

type Validator interface {
	Validate() (bool, error)
}

func ValidateEmail(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ValidatePassword(password string) (bool, error) {
	if password == "" {
		return false, errors.New("password cannot be empty")
	}
	return true, nil
}
