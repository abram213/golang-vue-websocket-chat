package db

import (
	"chat/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

type DataLayer interface {
	//User methods
	GetUserById(id uint) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UserExistByUsername(username string) bool
	UserExistByID(id uint) bool
	AddUserFriend(userID, friendID uint) error
	DeleteUserFriend(userID, friendID uint) error
	FriendshipExist(userID, friendID uint) bool
	GetUsers() (*[]model.User, error)
	GetUsersExcept(id uint) (*[]model.User, error)
	GetUserFriends(id uint) (*[]model.User, error)

	//Chat methods
	CreateChat(chat *model.Chat) (string, error)
	GetChatByID(id uint) (*model.Chat, error)
	GetChatByIdentifier(i string) (*model.Chat, error)
	GetUserChats(id uint) (*[]model.Chat, error)
	DeleteChat(id uint) error

	//Util methods
	Migrate(values ...interface{})
	DropTables(values ...interface{})
	CloseDB() error
}

type Database struct {
	*gorm.DB
}

func New(config *Config) (*Database, error) {
	db, err := gorm.Open("sqlite3", config.DatabaseURI)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &Database{db}, nil
}

func (db *Database) CloseDB() error {
	return db.Close()
}

func (db *Database) Migrate(values ...interface{}) {
	db.AutoMigrate(values...)
}

func (db *Database) DropTables(values ...interface{}) {
	db.DropTableIfExists(values...)
}
