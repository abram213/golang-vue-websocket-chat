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
	UserExist(username string) bool

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