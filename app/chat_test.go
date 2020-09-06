package app

import (
	"chat/db"
	"chat/model"
	"fmt"
	"testing"
)

func TestCreateChat(t *testing.T) {
	tApp := App{Database: &db.DatabaseTest{Users: []model.User{
		{Model: model.Model{ID: 1}},
		{Model: model.Model{ID: 2}},
		{Model: model.Model{ID: 3}},
		{Model: model.Model{ID: 4}},
	}}}
	user := &model.User{
		Model: model.Model{ID: 5},
	}
	ctx := tApp.NewContext().WithUser(user)

	chat := &model.Chat{
		Title: "",
		Users: []*model.User{
			{Model: model.Model{ID: 1}},
			{Model: model.Model{ID: 2}},
			{Model: model.Model{ID: 3}},
			{Model: model.Model{ID: 4}},
		},
	}

	if _, err := ctx.CreateChat(chat); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCreateChatError(t *testing.T) {
	tApp := App{Database: &db.DatabaseTest{}}
	user := &model.User{
		Model: model.Model{ID: 5},
	}
	ctx := tApp.NewContext().WithUser(user)
	var chatUserID uint = 1
	chat := &model.Chat{
		Title: "",
		Users: []*model.User{
			{Model: model.Model{ID: chatUserID}},
		},
	}

	errStr := fmt.Sprintf("no user with id: %v", chatUserID)

	_, err := ctx.CreateChat(chat)
	if err == nil {
		t.Errorf("Expected error, got %v", err)
		return
	}
	if err.Error() != errStr {
		t.Errorf("Expected error: %v, got: %v", errStr, err.Error())
	}
}

func TestGetChat(t *testing.T) {
	tApp := App{Database: &db.DatabaseTest{}}
	tCtx := tApp.NewContext()

	identifier := "chat_id"
	chat, err := tCtx.GetChat(identifier)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if chat.Identifier != identifier {
		t.Errorf("Chat identifier should be: %v, got: %v", identifier, chat.Identifier)
	}
}

func TestDeleteChat(t *testing.T) {
	tApp := App{Database: &db.DatabaseTest{}}
	tCtx := tApp.NewContext()

	var chatID uint = 1
	if err := tCtx.DeleteChat(chatID); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetChats(t *testing.T) {
	tApp := App{Database: &db.DatabaseTest{}}
	user := &model.User{
		Model: model.Model{ID: 5},
	}
	tCtx := tApp.NewContext().WithUser(user)

	if _, err := tCtx.GetChats(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
