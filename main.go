package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/dzwiedz90/go-loadenvconf/loadenvconf"
	"github.com/dzwiedz90/helpdesk-service-frontend/config"
	"github.com/dzwiedz90/helpdesk-service-frontend/pkg/users"
	"github.com/dzwiedz90/helpdesk-service-frontend/service"
)

func main() {
	cfg := &config.Config{}
	loadEnvConfig(cfg)

	httpAddress := cfg.HTTPAddress + ":" + cfg.HTTPPort

	fmt.Printf("Starting application helpdesk-service-frontend on port %s\n", httpAddress)

	usersClient := users.NewClient(cfg.UsersGRPCPort, cfg.UsersGRPCAddress)

	repo := service.NewRepo(cfg, usersClient)
	service.NewHandlers(repo)

	timeout, err := strconv.Atoi(cfg.Timeout)
	if err != nil {
		log.Fatal("Could not read timeout value form config", err)
	}
	serverTimeout := time.Duration(timeout) * time.Second

	srv := &http.Server{
		Addr:         httpAddress,
		Handler:      routes(cfg),
		ReadTimeout:  serverTimeout,
		WriteTimeout: serverTimeout,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func loadEnvConfig(cfg *config.Config) {
	_, err := loadenvconf.LoadEnvConfig(".env", cfg)
	if err != nil {
		log.Fatal("Could not load config", err)
	}
	log.Println("Config loaded!")
}

func routes(app *config.Config) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Route("/users", func(mux chi.Router) {
		mux.Post("/user.create", service.Repo.CreateUser)
	})

	return mux
}
