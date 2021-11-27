package users

import (
	"strings"

	"github.com/jrvldam/bookstore_users-api/utils/errors"
)

// Domain User
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate given user
func (u *User) Validate() *errors.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if u.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	return nil
}
