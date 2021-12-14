package users

import (
	"fmt"
	"strings"

	"github.com/jrvldam/bookstore_users-api/datasources/mysql/users_db"
	"github.com/jrvldam/bookstore_users-api/logger"
	"github.com/jrvldam/bookstore_users-api/utils/errors"
	"github.com/jrvldam/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, status, password, date_created) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE email=? AND password=? AND status=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("Error when trying to preare get user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)

	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); getErr != nil {
		logger.Error("Error when trying to get user by id", getErr)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("Error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.DateCreated)

	if saveErr != nil {
		logger.Error("Error when trying to save user", saveErr)
		return errors.NewInternalServerError("Database error")
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		logger.Error("Error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("Database error")
	}

	u.Id = userId

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("Error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)

	if err != nil {
		logger.Error("Error when trying to update user", err)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		logger.Error("Error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.Id); err != nil {
		logger.Error("Error when trying to delete user", err)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

func (u *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)

	if err != nil {
		logger.Error("Error when trying to prepare find by status user statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("Error when trying to find by status user", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
			logger.Error("Error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("Database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
	}

	return results, nil
}

func (u *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)

	if err != nil {
		logger.Error("Error when trying to preare find user by email and password statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Email, u.Password, StatusActive)

	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("Invalid users credentials")
		}

		logger.Error("Error when trying to get user by email and password", getErr)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}
