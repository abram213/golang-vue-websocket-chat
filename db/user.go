package db

import (
	"chat/model"

	"github.com/pkg/errors"
)

func (db *Database) GetUserById(id uint) (*model.User, error) {
	var user model.User
	return &user, errors.Wrap(db.First(&user, id).Error, "unable to get user")
}

func (db *Database) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	return &user, errors.Wrap(db.Where("username = ?", username).First(&user).Error, "unable to get user")
}

func (db *Database) CreateUser(user *model.User) error {
	return errors.Wrap(db.Create(&user).Error, "unable to create user")
}

func (db *Database) UserExist(username string) bool {
	var user model.User
	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return false
	}
	return true
}
