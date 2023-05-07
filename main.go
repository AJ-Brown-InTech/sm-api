package main

import (
	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

var Port, Environment string
var Version = "1.0.0"

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
	db, err := sqlx.Open("sqlite3", "database.db")
	if err != nil {
		log.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Error("error pinging database: %v", err.Error())
	}
	log.Infof("Connected to database at %s...", loc)

	// Initialize router and add handlers
	router := chi.NewRouter()
	//router.Use(middleware.Logger)
	//router.Use(middleware.Recoverer)
	// Routes
	router.Get("/", Test(log, db))
	router.Post("/user/create", Register(log, db))

	// r.Post("/user/update", routes.UpdateUser)
	// r.Post("/user/delete", routes.DeleteUser)
	// r.Post("/user/login", routes.Login)
	// r.Post("/user/logout", routes.Logout)
	// r.Post("/user/profile", routes.Profile)
	log.Infof("Libra Version: %s | Listening on port:%s | Time: %s", Version, Port, time.Now().In(loc))
	http.ListenAndServe(":"+Port, router)
}
