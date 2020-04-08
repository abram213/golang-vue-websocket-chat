package db

import (
	"chat/model"
)

//DatabaseTest struct for testing
type DatabaseTest struct {
	User  model.User
	Users []model.User
}

//User methods
func (db *DatabaseTest) GetUserById(id uint) (*model.User, error) {
	db.User.ID = id
	return &db.User, nil
}

func (db *DatabaseTest) GetUserByUsername(username string) (*model.User, error) {
	db.User.Username = username
	return &db.User, nil
}

func (db *DatabaseTest) CreateUser(user *model.User) error {
	return nil
}

func (db *DatabaseTest) UserExistByUsername(username string) bool {
	return true
}

func (db *DatabaseTest) UserExistByID(username string) bool {
	return true
}

func (db *DatabaseTest) AddUserFriend(userID, friendID uint) error {
	return nil
}

func (db *DatabaseTest) DeleteUserFriend(userID, friendID uint) error {
	return nil
}

func (db *DatabaseTest) FriendshipExist(userID, friendID uint) bool {
	return true
}

func (db *DatabaseTest) GetUsers() (*[]model.User, error) {
	return &db.Users, nil
}

func (db *DatabaseTest) GetUsersExcept(id uint) (*[]model.User, error) {
	return &db.Users, nil
}

func (db *DatabaseTest) GetUserFriends(id uint) (*[]model.User, error) {
	return &db.Users, nil
}

//Util methods
func (db *DatabaseTest) CloseDB() error {
	return nil
}

func (db *DatabaseTest) Migrate(values ...interface{}) {}

func (db *DatabaseTest) DropTables(values ...interface{}) {}
