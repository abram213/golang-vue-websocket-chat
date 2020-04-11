package api

import (
	"chat/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"chat/app"

	"github.com/go-chi/chi"
)

type chatInput struct {
	Title string        `json:"title"`
	Users []*model.User `json:"users"`
}

type chatResponse struct {
	Identifier string `json:"identifier"`
}

func (a *API) chatsRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(TokenAuth(a.App.Auth))
	r.Method("GET", "/", a.handler(a.GetChats))
	r.Method("POST", "/", a.handler(a.CreateChat))
	r.Route("/{id}", func(r chi.Router) {
		r.Method("GET", "/", a.handler(a.GetChat))
		r.Method("DELETE", "/", a.handler(a.DeleteChat))
	})
	return r
}

func (a *API) CreateChat(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input chatInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	chat := &model.Chat{
		Title: input.Title,
		Users: input.Users,
	}
	identifier, err := ctx.CreateChat(chat)
	if err != nil {
		return err
	}
	cResp := chatResponse{identifier}
	json.NewEncoder(w).Encode(cResp)
	return nil
}

func (a *API) GetChat(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	chat, err := ctx.GetChat(id)
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(chat)
	return nil
}

func (a *API) GetChats(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	chats, err := ctx.GetChats()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(chats)
	return nil
}

func (a *API) DeleteChat(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	err = ctx.DeleteChat(uint(uid))
	if err != nil {
		return err
	}
	return nil
}
