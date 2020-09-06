package app

import (
	"chat/db"
	"chat/model"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	tApp := App{Database: &db.DatabaseTest{}}
	msg := &model.Message{
		Body:           "Text",
		UserID:         1,
		ChatIdentifier: "123",
	}

	if err := tApp.CreateMessage(msg); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
