package main

import (
	"database/sql"
	"net/http"
	// "github.com/AJ-Brown-InTech/sm-api/middleware"
	// "github.com/AJ-Brown-InTech/sm-api/models"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

// create a new user
func Register(w http.ResponseWriter, r *http.Request, z *zap.SugaredLogger, d *sql.DB) error {
	// user := &models.UserLogin{
	// 	Username: "dog",
	// 	Email: "nnheo@example.com",

	// }
	err := d.Ping()
	if err != nil {
		return err
	}
	z.Infof("helo kitty: %v", err)
	w.Write([]byte("hello"))
	return nil
}
