package users

import (
	"fmt"

	"github.com/jrvldam/bookstore_users-api/datasources/mysql/users_db"
	"github.com/jrvldam/bookstore_users-api/utils/date_utils"
	"github.com/jrvldam/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[u.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", u.Id))
	}

	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated

	return nil
}

func (u *User) Save() *errors.RestErr {
	current := usersDB[u.Id]

	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("User %s already registered", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", u.Id))
	}

	u.DateCreated = date_utils.GetNowString()

	usersDB[u.Id] = u

	return nil
}
