package main

import (

	"time"


	"github.com/AJ-Brown-InTech/sm-api/pkg/config"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	// "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var (
	Port, Environment, Version, DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_DATABASE string
)

func init() {

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
  Environment = file.Env //TODO: not reading fix struct 
  Port = file.Server.Port

  logrus.Infof("ANOTHERTEST:%v",Environment)

  if Port == "" || Environment == "" || Version == "" || DB_HOST == "" || DB_PORT == "" || DB_USERNAME == "" || DB_PASSWORD == "" || DB_DATABASE == "" {
		panic("Enviroment varaibles aren't properly set, please revise or contact admin.")
  }

}

func main() {
	// Initialize logger
	//logger, _ := zap.NewProduction()
	//log := logger.Sugar()
	loc, _ := time.LoadLocation("America/Chicago")
	logrus.Infof("ANOTHERTEST:%v",loc)



	// // Initialize router and add handlers
	// r := chi.NewRouter()
	// router := chi.NewRouter()
	// router.Use(middleware.Logger)
	// //router.Use(middleware.Recoverer)
	// // Routes
	// router.Get("/", Test(log, db))
	// router.Post("/user/create", Register(log, db))
	// router.Post("/user/login", UserLogin(log, db))
	// router.Get("/user/profile/{id}", GetUserBySessionId(log, db))
	// router.Post("/user/update/{id}", UpdateUser(log, db))
	// router.Post("/user/follower/add/{id}", AddFollower(log, db))
	// router.Post("/user/follower/remove/{id}", RemoveFollower(log,db))
	// r.Post("/user/delete", routes.DeleteUser)
	

	// r.Mount("/api", router)
	// log.Infof(
	// 	"Libra Version: %s | Listening on port:%s | Time: %s",
	// 	Version,
	// 	Port,
	// 	time.Now().In(loc),
	// )
	// err = http.ListenAndServe(":"+Port, r)
	// if err != nil {

	// 	log.Panicf("Server launch error", err.Error())
	// 	panic("Server launch error")
	// }
}
