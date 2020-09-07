package main

import (
	"chat/api"
	"chat/app"
	"chat/model"
	"context"
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	port = flag.Int("port", 8080, "The port to bind the web application server to")

	dbURI   = flag.String("db_uri", "test.db", "Path to database file")
	envPath = flag.String("env_path", "config.env", "Path to .env file with config")

	migrate = flag.Bool("migrate", false, "Run db migration")
	refresh = flag.Bool("refresh", false, "Drop all tables before migration")
)

func main() {
	flag.Parse()

	config, err := InitConfig(*dbURI, *port, *envPath)
	if err != nil {
		log.Fatalf("init config error: %v", err)
	}

	app, err := app.New(config.DatabaseURI, config.APIKey)
	if err != nil {
		log.Fatalf("creating app error: %v", err)
	}
	defer app.Close()

	if *migrate {
		app.MigrateDB(*refresh,
			&model.User{},
			&model.Friendship{},
			&model.Tokens{},
			&model.Chat{},
			&model.Message{})
		//os.Exit(0)
	}

	api, err := api.New(app, config.Port)
	if err != nil {
		log.Fatalf("creating api error: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		cancel()
		log.Fatalln("signal caught. shutting down...")
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		serveAPI(ctx, api)
	}()

	wg.Wait()
}

func serveAPI(ctx context.Context, api *api.API) {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	router := chi.NewRouter()
	api.Init(router)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", api.Port),
		Handler:     cors.Handler(router),
		ReadTimeout: 60 * time.Second,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
		}
		close(done)
	}()

	log.Printf("serving api at http://127.0.0.1:%d", api.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
	<-done
}
