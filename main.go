package main

import (
	"context"
	"encoding/base64"
	"net/http"
	"time"
	"github.com/AJ-Brown-InTech/sm-api/pkg/config"
	"github.com/AJ-Brown-InTech/sm-api/pkg/database"
m   "github.com/AJ-Brown-InTech/sm-api/pkg/models"
api	"github.com/AJ-Brown-InTech/sm-api/pkg/router"
	"github.com/akyoto/cache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	Port, Environment, Version, DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_DATABASE string
)

var GlobalCache  *cache.Cache

func RequestMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd := m.Request{}
		rd.From = r.Header.Get("Request")
		rd.TraceId = r.Header.Get("Trace")
		if rd.TraceId == "" {
			rd.TraceId = uuid.New().String()
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "Request", rd)))
	})
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd := r.Context().Value("Request").(m.Request) 

		sessionID, found := GlobalCache.Get("session_id")
		if !found {
			logrus.WithFields(logrus.Fields{"Request": rd, "Error": "user session is not present"})
			http.Error(w, "user session is not present", http.StatusInternalServerError)
			return
		}

		session, ok := sessionID.(*m.Session)
		if !ok {
			logrus.WithFields(logrus.Fields{"RequestData": rd}).Error("Invalid session data in cache")
			http.Error(w, "Invalid session data in cache", http.StatusInternalServerError)
			return
		}

		userEmail, err := base64.StdEncoding.DecodeString(session.Token)
		if err != nil {
			logrus.WithFields(logrus.Fields{"RequestData": rd, "Error": err.Error()}).Error("Session token decryption failed")
			http.Error(w, "Session token decryption failed", http.StatusInternalServerError)
			return
		}

		logrus.Infof("Session valid for %s", userEmail)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "session_id", &session)))
	})
}

// func GetCurrentSession(db *sqlx.DB, c *cache.Cache) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		rd := r.Context().Value("RequestData").(RequestData)
// 		// check the user cache
// 		sessionId, found := c.Get("session_id")
// 		if !found {
// 			logrus.WithFields(logrus.Fields{"RequestData": rd, "Error": "No user session available, create a new user session"}).Error("No user session available, create a new user session")
// 			http.Error(w, "No user session available, create a new user session", http.StatusInternalServerError)
// 			return
// 		}
// 		session, ok := sessionId.(*authentication.Session)
// 		if !ok {
// 			logrus.WithFields(logrus.Fields{"RequestData": rd}).Error("Invalid session data in cache")
// 			http.Error(w, "Invalid session data in cache", http.StatusInternalServerError)
// 			return
// 		}

// 		logrus.WithFields(logrus.Fields{"RequestData": rd}).Infof("Successfully fetched user session")

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		err := json.NewEncoder(w).Encode(session)
// 		if err != nil {
// 			logrus.WithFields(logrus.Fields{
// 				"RequestData": rd,
// 				"Error":       err,
// 			}).Error("Error writing response")
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }


func init() { // ! could throw error if running test

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



func main() {

	loc, _ := time.LoadLocation("America/Chicago")

	//db initialization
	db, err := database.DatabaseInstance(DB_HOST,DB_PORT,DB_USERNAME,DB_PASSWORD,DB_DATABASE)
	if err != nil {
		logrus.Fatalf("Could not initialize database, %v", err)
		return 
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatalf("Could not connect to database, %v", err)
		return
	}

	// Global Cache
	c := cache.New(time.Hour * 1)
	GlobalCache = c 

	// router initialization
	r := chi.NewRouter() 
	router := chi.NewRouter()

	//middleware callstack
	r.Use(middleware.Recoverer)
	r.Use(RequestMiddlware)
	r.Use(SessionMiddleware)

	// Routes
	//TODO : login & register ar not protected routes
	r.Post("/user/", api.RegisterUserAccount(db, c))
	//r.Post("/user/login", api.UserLogin(db, c))
	//  r.Get("/user/profile/{id}", handle.GetUserBySessionId(db, c))
	//  r.Post("/user/update/{id}", handle.UpdateUser(db, c))
	//  r.Post("/user/follower/add/{id}", handle.AddFollower(db, c))
	//  r.Post("/user/follower/remove/{id}", handle.RemoveFollower(db, c))
	 //r.Post("/user/delete", handle.DeleteUser)
	 
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

	err = server.ListenAndServe()
	if err != nil {
		logrus.Panicf("Server launch error: %v", err)
		panic("Server launch error")
	}
}
