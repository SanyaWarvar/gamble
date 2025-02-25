package models

import (
	"regexp"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

const (
	UsernamePattern = "^[-a-zA-Z0-9_#$&*]+$"
	PasswordPattern = "^[-a-zA-Z0-9_#$&*]+$"
	EmailPattern    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	UsernameMaxLen  = 32
	UsernameMinLen  = 4
	PasswordMaxLen  = 32
	PasswordMinLen  = 8
)

func (u *User) IsValid() bool {
	matched, err := regexp.Match(EmailPattern, []byte(u.Email))
	if err != nil || !matched {
		return false
	}
	matched, err = regexp.Match(UsernamePattern, []byte(u.Username))
	usernameLen := len([]rune(u.Username))
	passwordLen := len([]rune(u.Password))
	if err != nil || !matched {
		return false
	}

	matched, err = regexp.Match(PasswordPattern, []byte(u.Password))
	if err != nil || !matched {
		return false
	}

	if (usernameLen <= UsernameMaxLen && usernameLen >= UsernameMinLen) &&
		(passwordLen <= PasswordMaxLen && passwordLen >= PasswordMinLen) {
		return true
	}
	return false
}
