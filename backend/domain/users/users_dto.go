package users

import (
	"main/utils/errors"
	"strings"
)

type User struct {
	ID        int64  `json:"ID"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
