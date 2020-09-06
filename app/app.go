package app

import (
	"chat/db"
	"fmt"

	"github.com/go-chi/jwtauth"
)

type App struct {
	Database db.DataLayer
	Auth     *jwtauth.JWTAuth
	Hub      *Hub
}

func (a *App) NewContext() *Context {
	return &Context{
		Database: a.Database,
	}
}

// Creates new App entity.
// App contains connection to Database, JWTAuth and WebSocket Hub.
func New(dbURI, apiKey string) (app *App, err error) {
	app = &App{
		Auth: jwtauth.New("HS256", []byte(apiKey), nil),
		Hub:  NewHub(),
	}
	app.Database, err = db.New(dbURI)
	if err != nil {
		return nil, err
	}

	return app, err
}

// Close close current app Database connection
func (a *App) Close() error {
	return a.Database.CloseDB()
}

// MigrateDB run auto migration for given models to app Database.
// If refresh flag = true, all tables will be dropped before migration.
func (a *App) MigrateDB(refresh bool, values ...interface{}) {
	if refresh {
		fmt.Println("dropping tables...")
		a.Database.DropTables(values)
	}

	a.Database.Migrate(values)
	fmt.Println("migration success!")
}

type ValidationError struct {
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

type AuthError struct {
	Message string `json:"message"`
}

func (e *AuthError) Error() string {
	return e.Message
}
