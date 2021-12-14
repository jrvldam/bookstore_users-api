package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jrvldam/bookstore_users-api/utils/errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("Error processing request")
}
