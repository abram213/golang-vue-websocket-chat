package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

func serveFiles(r chi.Router, path string) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	workDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir := http.Dir(filepath.Join(workDir, "dist"))
	fs := http.StripPrefix(path, http.FileServer(dir))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(fmt.Sprintf("%s", dir) + r.RequestURI); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(workDir, "dist/index.html"))
		} else {
			fs.ServeHTTP(w, r)
		}
	}))
}
