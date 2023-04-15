package main

import (
	//"fmt"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	//"net/url"
	"net/http"

	"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
)
var port, enviroment string
var Version = "1.0.0"
func LoadEnv(){
	port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	enviroment = os.Getenv("ENVIRONMENT")
	if enviroment == "" {
		enviroment = "development"
	}
}


func main() {
	//set enviroment
	LoadEnv()
	//init logger
	logger, _ := zap.NewProduction()
	log := logger.Sugar()
	loc, _ := time.LoadLocation("America/Chicago")
	//init router and add handlers
	r :=chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))
	//Todo: pass db, logger, and cache to handlers
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	
	log.Infof("port:%s", port,)
	log.Infof("Libra Version: %s | Listening on port:%s | Time: %s", Version, port, time.Now().In(loc))
  http.ListenAndServe(":"+port, r)
}

