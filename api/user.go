package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"chat/app"
	"chat/model"

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
	r.Group(func(r chi.Router) {
		r.Method("POST", "/", a.handler(a.CreateUser))
	})
	r.Group(func(r chi.Router) {
		r.Use(TokenAuth(a.App.Auth))
		r.Route("/{id:^[0-9]*$}", func(r chi.Router) {
			r.Method("GET", "/", a.handler(a.GetUser))
		})
	})
	return r
}

func (a *API) CreateUser(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input userInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	user := &model.User{
		Username: input.Username,
		Password: input.Password,
	}
	if err := ctx.CreateUser(user); err != nil {
		return err
	}

	json.NewEncoder(w).Encode(user)
	return nil
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
