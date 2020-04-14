package db

import (
	"chat/model"

	"github.com/pkg/errors"
)

func (db *Database) CreateChat(chat *model.Chat) (string, error) {
	return chat.Identifier, errors.Wrap(db.Create(&chat).Error, "unable to create chat")
}

func (db *Database) DeleteChat(id uint) error {
	chat := model.Chat{Model: model.Model{ID: id}}
	db.Model(&chat).Association("Users").Clear()
	//clear messages also
	return errors.Wrap(db.Delete(&chat).Error, "unable to delete chat")
}

func (db *Database) GetChatByID(id uint) (*model.Chat, error) {
	var chat model.Chat
	return &chat, errors.Wrap(db.First(&chat, id).Error, "unable to get chat")
}

func (db *Database) GetChatByIdentifier(i string) (*model.Chat, error) {
	var chat model.Chat
	pDb := db.Preload("Users").Preload("Messages").Preload("Messages.User")
	return &chat, errors.Wrap(pDb.Where("identifier = ?", i).First(&chat).Error, "unable to get chat")
}

func (db *Database) GetUserChats(id uint) (*[]model.Chat, error) {
	user := model.User{Model: model.Model{ID: id}}
	var chats []model.Chat
	return &chats, errors.Wrap(db.Model(&user).Related(&chats, "Chats").Error, "unable to get user chats")
}
