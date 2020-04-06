package db

import "chat/model"

//DatabaseTest struct for testing
type DatabaseTest struct {
	User model.User
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

func (db *DatabaseTest) UserExist(username string) bool {
	return true
}

//Util methods
func (db *DatabaseTest) CloseDB() error {
	return nil
}

func (db *DatabaseTest) Migrate(values ...interface{}) {}

func (db *DatabaseTest) DropTables(values ...interface{}) {}
