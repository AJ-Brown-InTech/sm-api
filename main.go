package main

import (
	"context"
	"net/http"
	"time"

	"github.com/AJ-Brown-InTech/sm-api/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	// "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	//"github.com/akyoto/cache"
)

var (
	Port, Environment, Version, DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_DATABASE string
)

type Middleware interface{
	ProccessMiddleware(args any) (http.Handler, error)
}

type RequestMiddleware struct {
	TracerId string
	From string
}

type SessionMiddleware struct {
	SessionToken string
	SessionExpiration string
}

func(r RequestMiddleware) ProccessMiddleware(next http.HandlerFunc) (http.HandlerFunc, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd := RequestMiddleware{}
		if rd.TracerId == "" {
			rd.TracerId = uuid.New().String()
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "TracerId", rd.TracerId)))
	}), nil
}

func(s SessionMiddleware) ProccessMiddleware(next http.HandlerFunc) (http.HandlerFunc, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, req)
	}), nil
} 

type localCache struct {
	Token SessionMiddleware
}





func init() { // could throw error if running test

	file, err := config.SetupEnv()
	if err != nil {
		logrus.Errorf("Error reading config file: %v", err)
		return
	}

	DB_DATABASE = file.Database.DBName
	DB_PASSWORD = file.Database.Password
	DB_USERNAME = file.Database.Username
	DB_PORT = file.Database.Port
	DB_HOST = file.Database.Host
	Version = file.Version.Release
	Environment = file.Enviroment.Env 
	Port = file.Server.Port

	if Port == "" || Environment == "" || Version == "" || DB_HOST == "" || DB_PORT == "" || DB_USERNAME == "" || DB_PASSWORD == "" || DB_DATABASE == "" {
		panic("Enviroment varaibles aren't properly set, please revise")
	}
}

//Todo: add cache & db
func main() {
	
	loc, _ := time.LoadLocation("America/Chicago")

	// router initialization
	r := chi.NewRouter() 
	router := chi.NewRouter()

	//middleware callstack
	r.Use(middleware.Recoverer)
	
	
	// // Routes
	// router.Get("/", Test(log, db))
	// router.Post("/user/create", Register(log, db))
	// router.Post("/user/login", UserLogin(log, db))
	// router.Get("/user/profile/{id}", GetUserBySessionId(log, db))
	// router.Post("/user/update/{id}", UpdateUser(log, db))
	// router.Post("/user/follower/add/{id}", AddFollower(log, db))
	// router.Post("/user/follower/remove/{id}", RemoveFollower(log,db))
	// r.Post("/user/delete", routes.DeleteUser)
	 
	router.Mount("/api", r)

	 logrus.Infof(
	 	"\n Libra Version: %s\n Listening on port:%s\n Current Time: %s",
	 	Version,
	 	Port,
	 	time.Now().In(loc),
	 )

	server := &http.Server{
		Addr:         ":" + Port,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	err := server.ListenAndServe()
	if err != nil {
		logrus.Panicf("Server launch error: %v", err)
		panic("Server launch error")
	}
}
