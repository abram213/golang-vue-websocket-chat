package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"chat/app"

	"github.com/go-chi/chi"
)

type UserResponse struct {
	ID uint `json:"id"`
}

func (a *API) usersRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(TokenAuth(a.App.Auth))
	r.Method("GET", "/info", a.handler(a.GetUserInfo))
	r.Method("GET", "/", a.handler(a.GetUsers))
	r.Route("/{id:^[0-9]*$}", func(r chi.Router) {
		r.Method("GET", "/", a.handler(a.GetUser))
	})
	r.Route("/friends", func(r chi.Router) {
		r.Method("GET", "/", a.handler(a.GetUserFriends))
		r.Route("/{id:^[0-9]*$}", func(r chi.Router) {
			r.Method("GET", "/add", a.handler(a.AddFriend))
			r.Method("DELETE", "/delete", a.handler(a.DeleteFriend))
		})
	})
	return r
}

func (a *API) GetUser(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	user, err := ctx.GetUserById(uint(uid))
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(user)
	return nil
}

func (a *API) AddFriend(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	err = ctx.AddFriendByID(uint(uid))
	if err != nil {
		return err
	}
	return nil
}

func (a *API) DeleteFriend(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	err = ctx.DeleteFriendByID(uint(uid))
	if err != nil {
		return err
	}
	return nil
}

func (a *API) GetUsers(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	users, err := ctx.GetUsers()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(users)
	return nil
}

func (a *API) GetUserFriends(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	friends, err := ctx.GetUserFriends()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(friends)
	return nil
}

func (a *API) GetUserInfo(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	user, err := ctx.GetUserById(ctx.User.ID)
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(user)
	return nil
}
