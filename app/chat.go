package app

import (
	"chat/model"
	"fmt"
	"strings"
)

func (ctx *Context) CreateChat(c *model.Chat) (string, error) {
	var usernames []string
	for _, user := range c.Users {
		if !ctx.Database.UserExistByID(user.ID) {
			return "", fmt.Errorf("no user with id: %v", user.ID)
		}
		usernames = append(usernames, user.Username)
	}
	if c.Title == "" {
		if len(usernames) > 3 {
			usernames = usernames[:3]
		}
		c.Title = strings.Join(usernames, ", ")
	}
	c.UserID = ctx.User.ID
	c.Users = append(c.Users, ctx.User)
	return ctx.Database.CreateChat(c)
}

func (ctx *Context) GetChat(identifier string) (*model.Chat, error) {
	return ctx.Database.GetChatByIdentifier(identifier)
}

func (ctx *Context) DeleteChat(id uint) error {
	return ctx.Database.DeleteChat(id)
}

func (ctx *Context) GetChats() (*[]model.Chat, error) {
	return ctx.Database.GetUserChats(ctx.User.ID)
}
