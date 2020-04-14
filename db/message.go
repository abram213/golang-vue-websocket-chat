package db

import (
	"chat/model"
	"github.com/pkg/errors"
)

func (db *Database) CreateMessage(m *model.Message) error {
	return errors.Wrap(db.Create(&m).Error, "unable to create message")
}
