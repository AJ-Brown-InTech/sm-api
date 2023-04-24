package main

import (
	"context"
	"database/sql"
	//"github.com/AJ-Brown-InTech/sm-api/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

var port, enviroment string
var Version = "1.0.0"

func LoadEnv() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	enviroment = os.Getenv("ENVIRONMENT")
	if enviroment == "" {
		enviroment = "development"
	}
}

func Libre_Db() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, test TEXT)")
	if err != nil {
		panic(err)
	}

	return db, nil
}

func main() {
	// set enviroment
	LoadEnv()
	// load database
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, test TEXT)")
	if err != nil {
		panic(err)
	}

	//init logger
	logger, _ := zap.NewProduction()
	log := logger.Sugar()
	loc, _ := time.LoadLocation("America/Chicago")
	//init router and add handlers
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))
	//Todo: pass db, logger, and cache to handlers
	//Routes
	r.Post("/user/create", func(w http.ResponseWriter, r *http.Request) {
	Register(w, r, log, db)
	})
	// r.Post("/user/update", routes.UpdateUser)
	// r.Post("/user/delete", routes.DeleteUser)
	// r.Post("/user/login", routes.Login)
	// r.Post("/user/logout", routes.Logout)
	// r.Post("/user/profile", routes.Profile)
	// r.

	log.Infof("port:%s", port)
	log.Infof("Libra Version: %s | Listening on port:%s | Time: %s", Version, port, time.Now().In(loc))
	http.ListenAndServe(":"+port, r)
}

func Register(w http.ResponseWriter, r *http.Request, log *zap.SugaredLogger, db *sql.DB) {
	panic("unimplemented")
}
