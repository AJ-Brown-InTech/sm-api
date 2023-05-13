package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

var (
	Port, Environment string
	Version           = "1.0.0"
)

func init() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "80"
	}
	Environment = os.Getenv("ENVIRONMENT")
	if Environment == "" {
		Environment = "development"
	}
}

func main() {
	// Initialize logger
	logger, _ := zap.NewProduction()
	log := logger.Sugar()
	loc, _ := time.LoadLocation("America/Chicago")

	// Load database
	db, err := sqlx.Open("sqlite3", "./database/libra.db")
	if err != nil {
		log.Errorf("Error connecting to database: %v", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Error("error pinging database: %v", err.Error())
		return
	}
	log.Infof("Connected to database at %s...", loc)

	// Initialize router and add handlers
	r := chi.NewRouter()
	router := chi.NewRouter()
	//router.Use(middleware.Logger)
	//router.Use(middleware.Recoverer)
	// Routes
	router.Get("/", Test(log, db))
	router.Post("/user/create", Register(log, db))
	router.Get("/user/{id}", GetUserBySessionId(log, db))
	router.Post("/user/login", UserLogin(log, db))
	// r.Post("/user/delete", routes.DeleteUser)
	// r.Post("/user/login", routes.Login)
	// r.Post("/user/logout", routes.Logout)
	// r.Post("/user/profile", routes.Profile)
	r.Mount("/api", router)
	log.Infof(
		"Libra Version: %s | Listening on port:%s | Time: %s",
		Version,
		Port,
		time.Now().In(loc),
	)
	err = http.ListenAndServe(":"+Port, r)
	if err != nil {

		log.Panicf("Server launch error", err.Error())
		panic("Server launch error")
	}
}
