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

func (db *Database) UserExistByUsername(username string) bool {
	var user model.User
	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return false
	}
	return true
}

func (db *Database) UserExistByID(id uint) bool {
	var user model.User
	if db.First(&user, id).RecordNotFound() {
		return false
	}
	return true
}

func (db *Database) AddUserFriend(userID, friendID uint) error {
	friendship := model.Friendship{userID, friendID}
	return errors.Wrap(db.Create(&friendship).Error, "unable to create friendship")
}

func (db *Database) DeleteUserFriend(userID, friendID uint) error {
	var friendship model.Friendship
	return errors.Wrap(db.Where("user_id = ? and friend_id = ?", userID, friendID).Delete(&friendship).Error, "unable to delete friendship")
}

func (db *Database) GetUsers() (*[]model.User, error) {
	var users []model.User
	return &users, errors.Wrap(db.Find(&users).Error, "unable to get users")
}

func (db *Database) GetUsersExcept(id uint) (*[]model.User, error) {
	var users []model.User
	return &users, errors.Wrap(db.Where("id <> ?", id).Find(&users).Error, "unable to get users")
}

func (db *Database) GetUserFriends(id uint) (*[]model.User, error) {
	var user model.User
	user.ID = id
	var friends []model.User
	return &friends, errors.Wrap(db.Model(&user).Related(&friends, "Friends").Error, "unable to get user friends")
}

func (db *Database) FriendshipExist(userID, friendID uint) bool {
	var friendship model.Friendship
	if db.Where("user_id = ? and friend_id = ?", userID, friendID).First(&friendship).RecordNotFound() {
		return false
	}
	return true
}
