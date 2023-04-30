package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/AfterShip/email-verifier"
	"github.com/Masterminds/squirrel"

	//"github.com/go-chi/chi/v5"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

// create a new user
func CreateUser(z *zap.SugaredLogger, d *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newUser = &UserProfile{}
		err := json.NewDecoder(r.Body).Decode(newUser)
		if err != nil {
			z.Errorf("Error: %v", err)
		}

		if len(newUser.Username) > 15 || len(newUser.Username) < 6 {
			z.Errorf("Error: the username needs to be 6 characters minimum")
			return
		}

		if ok := emailverifier.IsAddressValid(newUser.Email); !ok {
			z.Errorf("Error: email address not is valid")
			return
		}

		if len(newUser.Password) < 12 {
			z.Errorf("Error: the password needs to be 12 characters minimum")
			return
		}

		sessionId, _ := CreateSession(w, r, newUser.Username, z)
		_, err = d.Exec(`Insert into users(username, email, password, rating, session_id) Values($1, $2, $3, $4, $5)`, newUser.Username, newUser.Email, newUser.Password, newUser.Rating, sessionId)
		if err != nil {
			z.Errorf("Error: Database error = %w", err)
		}

		var createdUser UserProfile
		query := fmt.Sprintf("Select * from user where username = %s", newUser.Username)
		_, err = d.Exec(query, createdUser) //(&createdUser, query)
		if err != nil {
			z.Errorf("Error: Database retrieving user = %w", err)
		}

		var createdUser UserProfile
		val, err := json.Marshal()
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode()
		// user := &models.UserLogin{
		// 	Username: "dog",
		// 	Email: "nnheo@example.com",
		// }
		// err := d.Ping()
		// if err != nil {
		// 	return err
		// }
		// z.Infof("helo kitty: %v", err)
		// w.Write([]byte("hello"))

	}
}
