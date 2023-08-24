package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/dzwiedz90/go-loadenvconf/loadenvconf"
	"github.com/dzwiedz90/helpdesk-service-frontend/config"
	"github.com/dzwiedz90/helpdesk-service-frontend/logs"
	"github.com/dzwiedz90/helpdesk-service-frontend/pkg/users"
	"github.com/dzwiedz90/helpdesk-service-frontend/service"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func main() {
	cfg := &config.Config{}
	loadEnvConfig(cfg)

	httpAddress := cfg.HTTPAddress + ":" + cfg.HTTPPort

	logs.InfoLogger(fmt.Sprintf("Starting application helpdesk-service-frontend on port %s", httpAddress))

	usersClient := users.NewClient(cfg.UsersGRPCPort, cfg.UsersGRPCAddress)

	repo := service.NewRepo(cfg, usersClient)
	service.NewHandlers(repo)

	timeout, err := strconv.Atoi(cfg.Timeout)
	if err != nil {
		message := fmt.Sprintf("Could not read timeout value form config: %v", err)
		logs.ErrorLogger(message)
		log.Fatal(message)
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
		message := fmt.Sprintf("Failed to start server: %v", err)
		logs.ErrorLogger(message)
		log.Fatal(message)
	}
}

func loadEnvConfig(cfg *config.Config) {
	_, err := loadenvconf.LoadEnvConfig(".env", cfg)
	if err != nil {
		message := fmt.Sprintf("Could not load config: %v", err)
		logs.ErrorLogger(message)
		log.Fatal(message)
	}

	evaluateLogFilesSize()
	logs.NewLoggers(cfg)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg.ErrorLog = errorLog

	logs.InfoLogger("Config loaded!")
}

func routes(app *config.Config) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Route("/users", func(mux chi.Router) {
		mux.Post("/user.create", service.Repo.CreateUser)
		mux.Get("/user.get", service.Repo.GetUser)
		mux.Get("/user.get_all", service.Repo.GetAllUsers)
		mux.Get("/user.update", service.Repo.UpdateUser)
		mux.Get("/user.delete", service.Repo.DeleteUser)
	})

	return mux
}

func evaluateLogFilesSize() error {
	maxFileSize := int64(20 * 1024 * 1024)

	consoleLog := "logs/console.log"
	errorLog := "logs/error.log"

	// Check console.log
	fileInfo, err := os.Stat(consoleLog)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Error when getting file information: %v", err))
		return err
	}
	fileSize := fileInfo.Size()

	// Check error.log
	fileInfo2, err := os.Stat(errorLog)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Error when getting file information: %v", err))
		return err
	}
	fileSize2 := fileInfo2.Size()

	// Check if file size greater than 20MB and prune it if so
	if fileSize > maxFileSize {
		logs.InfoLogger("console.log over 20MB - cleaning")
		file, err := os.Create(consoleLog)
		if err != nil {
			logs.ErrorLogger(fmt.Sprintf("Error during console.log cleaning: %v", err))
			panic(err)
		}
		defer file.Close()
	}
	if fileSize2 > maxFileSize {
		logs.InfoLogger("error.log over 20MB - cleaning")
		file2, err := os.Create(errorLog)
		if err != nil {
			logs.ErrorLogger(fmt.Sprintf("Error during error.log cleaning: %v", err))
			panic(err)
		}
		defer file2.Close()
	}

	return nil
}
