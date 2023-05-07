package main

import (
	"encoding/json"
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"net/http"
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
			return
		}
		// validate data
		switch {
		case len(newUser.Username) > 15:
			z.Error("username too long")
			return
		case len(newUser.Username) < 6:
			z.Error("username too short")
			return
		default:
			z.Info("username valid")
		}
		verify := emailverifier.NewVerifier()
		res, _ := verify.Verify(newUser.Email)
		if !res.Syntax.Valid {
			z.Error("email not valid")
			return
		}
		if len(newUser.Password) < 12 {
			z.Errorf("Error: the password needs to be 12 characters minimum")
			return
		}

		// TODO: hash the password
		// create a session and store with other data
		sessionID := uuid.New().String()
		query := fmt.Sprintf(
			"INSERT INTO users (username, email, password, rating, birthday, session_id) VALUES ('%s', '%s', '%s', %d, '%s', '%s')",
			newUser.Username,
			newUser.Email,
			newUser.Password,
			newUser.Rating,
			newUser.Birthday,
			sessionID,
		)
		_, err = db.Exec(query)
		if err != nil {
			z.Errorf("Error: creating user  = %w", err)
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
			return
		}
	}
}
