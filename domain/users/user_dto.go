package users

import (
	"strings"

	"github.com/jrvldam/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

// Domain User
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
}

// Validate given user
func (u *User) Validate() *errors.RestErr {
	u.FirstName = strings.TrimSpace(strings.ToLower(u.FirstName))
	u.LastName = strings.TrimSpace(strings.ToLower(u.LastName))
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if u.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	u.Password = strings.TrimSpace(u.Password)

	if u.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	return nil
}
