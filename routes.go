package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"net/http"
)

// create a new user
func CreateUser(z *zap.SugaredLogger, d *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
