package app

import (
	"chat/model"
)

func (app *App) CreateMessage(m *model.Message) error {
	return app.Database.CreateMessage(m)
}
