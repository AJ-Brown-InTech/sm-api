package main

import (
	"encoding/json"
	"fmt"
	"net/http"
"time"
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func Test(z *zap.SugaredLogger, db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		z.Infof("Connected")
	}
}

// create a new user
func Register(z *zap.SugaredLogger, db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser UserProfile
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()
		err := d.Decode(&newUser)
		if err != nil {
			z.Errorf("Error: %v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// validate data
		switch {
		case len(newUser.Username) > 15:
			z.Error("username too long")
			http.Error(w, "username too long", http.StatusInternalServerError)
			return
		case len(newUser.Username) < 6:
			z.Error("username too short")
			http.Error(w, "username too short", http.StatusInternalServerError)
			return
		default:
			z.Info("username valid")
		}
		verify := emailverifier.NewVerifier()
		res, _ := verify.Verify(newUser.Email)
		if !res.Syntax.Valid {
			z.Error("email not valid")
			http.Error(w, "email not valid", http.StatusInternalServerError)
			return
		}
		if len(newUser.Password) < 12 {
			z.Errorf("Error: the password needs to be 12 characters minimum")
			http.Error(w, "the password needs to be 12 characters minimum", http.StatusInternalServerError)
			return
		}

		// TODO: hash the password
		// create a session and store with other data
		sessionID := uuid.New().String()
		query := fmt.Sprintf(`INSERT INTO users(username, email, password, rating, birthday, session_id, updated_at, created_at) VALUES ('%s', '%s', '%s', %f, '%s', '%s', '%s', '%s')`,
			newUser.Username,
			newUser.Email,
			newUser.Password,
			newUser.Rating,
			newUser.Birthday,
			sessionID,
			time.Now(),
			time.Now(),
		)
		//INSERT INTO users(username, email, password, rating, birthday, session_id) VALUES ('JohnDoe123', 'johndoe@example.com', 'mypassword123', 0, '1996-03-05 00:00:00 +0000 UTC', 'eb2249c7-c7ae-4cb2-a3de-d600cb6e0a51')

		_, err = db.Exec(query)
		if err != nil {
			z.Errorf("Error: creating user: %v", err.Error())
			http.Error(w, "error creating user", http.StatusInternalServerError)
			return
		}

		// return the session_id
		type Session struct {
			ID string `json:"session_id"`
		}
		result := Session{ID: sessionID}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(&result)
		if err != nil {
			z.Error("Error writing response")
			http.Error(w, "error writing respondse", http.StatusInternalServerError)
			return
		}
	}
}
