package app

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"chat/model"
)

func (a *App) GetUserById(id uint) (*model.User, error) {
	return a.Database.GetUserById(id)
}

/*  */

func (ctx *Context) CreateUser(u *model.User) error {
	if err := ctx.validateUser(u); err != nil {
		return err
	}
	if err := u.HashPassword(); err != nil {
		return err
	}
	return ctx.Database.CreateUser(u)
}

func (ctx *Context) GetUserById(id uint) (*model.User, error) {
	return ctx.Database.GetUserById(id)
}

func (ctx *Context) validateUser(c *model.User) *ValidationError {
	if err := validation.ValidateStruct(c,
		validation.Field(&c.Username, validation.Required, validation.Length(5, 50), validation.By(uniqueByUsername(ctx))),
		validation.Field(&c.Password, validation.Required, validation.Length(8, 30)),
	); err != nil {
		return &ValidationError{err.Error()}
	}
	return nil
}

func uniqueByUsername(ctx *Context) validation.RuleFunc {
	return func(value interface{}) error {
		username, _ := value.(string)
		if ctx.Database.UserExistByUsername(username) {
			return errors.New("already exist user with such username")
		}
		return nil
	}
}

func (ctx *Context) AddFriendByID(id uint) error {
	if !ctx.Database.UserExistByID(id) {
		return errors.New("no such user")
	}

	userID := ctx.User.ID
	if ctx.Database.FriendshipExist(userID, id) {
		return errors.New("friendship already exist")
	}
	if userID == id {
		return errors.New("you can`t add yorself to friends")
	}
	return ctx.Database.AddUserFriend(userID, id)
}

func (ctx *Context) DeleteFriendByID(id uint) error {
	userID := ctx.User.ID
	return ctx.Database.DeleteUserFriend(userID, id)
}

func (ctx *Context) GetUsers() (*[]model.User, error) {
	return ctx.Database.GetUsersExcept(ctx.User.ID)
}

func (ctx *Context) GetUserFriends() (*[]model.User, error) {
	return ctx.Database.GetUserFriends(ctx.User.ID)
}
