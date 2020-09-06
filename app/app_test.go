package app

import (
	"os"
	"testing"
)

func TestNewApp(t *testing.T) {
	dbURI := "temp.db"
	app, err := New(dbURI, "test_api_key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer os.Remove(dbURI)
	defer app.Close()

	_, err = os.Stat(dbURI)
	if os.IsNotExist(err) {
		t.Errorf("Unexpected error: %v", err)
	}

	ctx := app.NewContext()
	if app.Database != ctx.Database {
		t.Errorf("ctx Database and app Database must be similar\nappDB: %v\nctxDB: %v", app.Database, ctx.Database)
	}
}

func TestNewAppDBConnErr(t *testing.T) {
	dbURI := "test/*db" //invalid db file
	_, err := New(dbURI, "test_api_key")
	if err == nil {
		t.Errorf("Expected error, but got %v", err)
	}
}
