package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"chat/app"

	"github.com/go-chi/chi"
)

type userInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID uint `json:"id"`
}

func (a *API) usersRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(TokenAuth(a.App.Auth))
	r.Route("/{id:^[0-9]*$}", func(r chi.Router) {
		r.Method("GET", "/", a.handler(a.GetUser))
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
