package routes

import (
	"database/sql"
	"net/http"

	// "github.com/AJ-Brown-InTech/sm-api/middleware"
	// "github.com/AJ-Brown-InTech/sm-api/models"
	"go.uber.org/zap"
)

// create a new user
func CreateUser(w http.ResponseWriter, r *http.Request, z *zap.SugaredLogger, db *sql.DB) error {
	// user := &models.UserLogin{
	// 	Username: "dog",
	// 	Email: "nnheo@example.com",

	// }

	w.Write([]byte("hello"))
	return nil
}
