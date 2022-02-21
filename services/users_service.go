package services

import (
	"github.com/williammoraes/bookstore_users-api/domain/users"
	"github.com/williammoraes/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
