package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (a *API) wsRouter() http.Handler {
	r := chi.NewRouter()
	go a.App.Hub.Run()
	r.HandleFunc("/conn", func(w http.ResponseWriter, r *http.Request) {
		a.App.Hub.ServeWs(a.App, w, r)
	})
	return r
}
