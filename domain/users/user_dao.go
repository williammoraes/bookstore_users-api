package users

import (
	"fmt"
	"log"

	"github.com/williammoraes/bookstore_users-api/datasources/mysql/users_db"
	dateutils "github.com/williammoraes/bookstore_users-api/utils/date_utils"
	"github.com/williammoraes/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		log.Fatalln(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewBadResquesError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DataCreated = result.DataCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadResquesError(fmt.Sprintf("email %s already registred", user.Email))
		}
		return errors.NewBadResquesError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DataCreated = dateutils.GetNowString()
	usersDB[user.Id] = user
	return nil
}
