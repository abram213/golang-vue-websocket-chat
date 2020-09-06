package api

import (
	"chat/app"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"strconv"
	"time"
)

type API struct {
	App  *app.App
	Port int
}

func New(a *app.App, port int) (api *API, err error) {
	api = &API{
		App:  a,
		Port: port,
	}
	return api, nil
}

func (a *API) Init(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	serveFiles(r, "/")
	r.Mount("/ws", a.wsRouter())
	r.Route("/api", func(r chi.Router) {
		r.Mount("/auth", a.authRouter())
		r.Mount("/users", a.usersRouter())
		r.Mount("/chats", a.chatsRouter())
	})
}

func (a *API) handler(f func(*app.Context, http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)

		ctx := a.App.NewContext()

		if _, claims, ok := TokenFromContext(r.Context()); ok {
			//check user_id before
			userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
			if err != nil {
				http.Error(w, fmt.Sprintf("bad user_id, err: %v", err), 401)
				return
			}
			user, err := a.App.GetUserById(uint(userID))
			if err != nil {
				http.Error(w, fmt.Sprintf("problem with user, err: %v", err), 401)
				return
			}
			ctx = ctx.WithUser(user)
		}

		w.Header().Set("Content-Type", "application/json")

		if err := f(ctx, w, r); err != nil {
			if verr, ok := err.(*app.ValidationError); ok {
				data, err := json.Marshal(verr)
				if err == nil {
					w.WriteHeader(http.StatusBadRequest)
					_, err = w.Write(data)
				}

				if err != nil {
					log.Println(err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			} else if uerr, ok := err.(*app.AuthError); ok {
				data, err := json.Marshal(uerr)
				if err == nil {
					w.WriteHeader(http.StatusUnauthorized)
					_, err = w.Write(data)
				}

				if err != nil {
					log.Println(err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			} else {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}
	})
}
