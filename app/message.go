package app

import (
	"chat/model"
)

func (ctx *Context) CreateMessage(m *model.Message) error {
	return ctx.Database.CreateMessage(m)
}
