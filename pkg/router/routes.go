package main

import (
	//"encoding/base64"
	//"encoding/json"
	//"fmt"
	"net/http"
	//"time"
	//"github.com/go-playground/validator/v10"
	//m "github.com/AJ-Brown-InTech/sm-api/pkg/models"
	//sq "github.com/Masterminds/squirrel"
	"github.com/akyoto/cache"
	//"github.com/go-chi/chi/v5"
	//"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	//"github.com/sirupsen/logrus"
)

func Register(db *sqlx.DB, c *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var newUser User
		// d := json.NewDecoder(r.Body)
		// d.DisallowUnknownFields()
		// err := d.Decode(&newUser)
		// if err != nil {
		// 	logrus.Errorf("Error: %v", err.Error())
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// validate data
		// switch {
		// case len(newUser.Username) > 15:
		// 	logrus.Error("username too long")
		// 	http.Error(w, "username too long", http.StatusBadRequest)
		// 	return
		// case len(newUser.Username) < 6:
		// 	z.Error("username too short")
		// 	http.Error(w, "username too short", http.StatusBadRequest)
		// 	return
		// default:
		// 	z.Info("username valid")
		// }
		// verify := emailverifier.NewVerifier()
		// res, _ := verify.Verify(newUser.Email)
		// if !res.Syntax.Valid {
		// 	z.Error("email not valid")
		// 	http.Error(w, "email not valid", http.StatusBadRequest)
		// 	return
		// }
		// if len(newUser.Password) < 18 {
		// 	z.Errorf("Error: the password needs to be 12 characters minimum")
		// 	http.Error(w, "the password needs to be 12 characters minimum", http.StatusBadRequest)
		// 	return
		// }

		
		//encodedPassword := base64.StdEncoding.EncodeToString([]byte(newUser.Password))
		// create a session and store with other data
		// sessionID := uuid.New().String()
		// query := fmt.Sprintf(
		// 	`INSERT INTO users(username, email, password, fullname, bio, whoareyou, avatar, account_rating, post_rating, follower_count, following_count, post_count, location, session_id, birthday, updated_at, created_at, active)
		// 	VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %v, %f, %f, %d, %d, %d, '%s', '%s', '%s', '%v', '%v', %t)`,
		// 	newUser.Username,
		// 	newUser.Email,
		// 	encodedPassword,
		// 	newUser.FullName.String,
		// 	newUser.Bio.String,
		// 	newUser.WhoAreYou.String,
		// 	newUser.Avatar,
		// 	newUser.AccountRating,
		// 	newUser.PostRating,
		// 	newUser.FollowerCount,
		// 	newUser.FollowingCount,
		// 	newUser.PostCount,
		// 	newUser.Location.String,
		// 	sessionID,
		// 	newUser.Birthday,
		// 	time.Now(),
		// 	time.Now(),
		// 	newUser.Active,
		// )
		
		// INSERT INTO users(username, email, password, rating, birthday, session_id) VALUES ('JohnDoe123', 'johndoe@example.com', 'mypassword123', 0, '1996-03-05 00:00:00 +0000 UTC', 'eb2249c7-c7ae-4cb2-a3de-d600cb6e0a51')

		// _, err = db.Exec(query)
		// if err != nil {
		// 	z.Errorf("Error: creating user: %v", err.Error())
		// 	http.Error(w, "error creating user", http.StatusInternalServerError)
		// 	return
		// }

		// // return the session_id
		// type Session struct {
		// 	ID string `json:"session_id"`
		// }
		// result := Session{ID: sessionID}
		// w.Header().Set("Content-type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// err = json.NewEncoder(w).Encode(&result)
		// if err != nil {
		// 	z.Error("Error writing response")
		// 	http.Error(w, "error writing respondse", http.StatusInternalServerError)
		// 	return
		// }
	}
}

		//Retrieve all followers of a user
		// SELECT users.username
		// FROM users
		// JOIN followers ON followers.follower_share_key = users.share_key
		// WHERE followers.followed_share_key = 'user_share_key';

		// -- Retrieve all users that a user is following
		// SELECT users.username
		// FROM users
		// JOIN followers ON followers.followed_share_key = users.share_key
		// WHERE followers.follower_share_key = 'user_share_key';
// func GetUserBySessionId(db *sqlx.DB, c *cache.Cache) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "id")
// 		query := fmt.Sprintf(`Select * from users 
// 		JOIN followers ON followers.follower = users.share_key
// 		JOIN followers ON followers.followed = users.share_key
// 		WHERE session_id = '%s'`, id)

// 		var user User
// 		err := db.Get(&user, query)
// 		if err != nil {
// 			logrus.Errorf("error fetching user from db :%v", err.Error())
// 			http.Error(w, "error fetching user from database", http.StatusInternalServerError)
// 			return
// 		}
		
// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-type", "application/json")
// 		err = json.NewEncoder(w).Encode(&user)
// 		if err != nil {
// 			logrus.Error("error writing response")
// 			http.Error(w, "error writing response", http.StatusInternalServerError)
// 		}
// 	}
// }

// func UserLogin(db *sqlx.DB, c *cache.Cache) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var creds Login
// 		decoder := json.NewDecoder(r.Body)
// 		decoder.DisallowUnknownFields()
// 		err := decoder.Decode(&creds)
// 		if err != nil {
// 			logrus.Errorf("error reading in request body: %v", err.Error())
// 			http.Error(w, "request body error", http.StatusBadRequest)
// 			return
// 		}

// 		superString := base64.StdEncoding.EncodeToString([]byte(creds.Password))
// 		var query string
// 		if creds.Email.Valid {
// 			query = fmt.Sprintf(
// 				`Select email, password from users where email = '%s' and password = '%s'`,
// 				creds.Email.String,
// 				superString,
// 			)
// 		}
// 		if creds.Username.Valid {
// 			query = fmt.Sprintf(
// 				`Select username, password from users where username = '%s' and password = '%s'`,
// 				creds.Username.String,
// 				superString,
// 			)
// 		}
// 		_, err = db.Exec(query)
// 		if err != nil {
// 			logrus.Errorf("error fetching user from databse", err.Error())
// 			http.Error(w, "error fetching user from database", http.StatusInternalServerError)
// 			return
// 		}

// 		sessionId := uuid.New().String()
// 		var updateStatement string
// 		if creds.Email.Valid {
// 			updateStatement = fmt.Sprintf(
// 				`update users set session_id = '%s' where email = '%s'`,
// 				sessionId,
// 				creds.Email.String,
// 			)
// 		}
// 		if creds.Username.Valid {
// 			updateStatement = fmt.Sprintf(
// 				`update users set session_id = '%s' where username = '%s'`,
// 				sessionId,
// 				creds.Username.String,
// 			)
// 		}

// 		_, err = db.Exec(updateStatement)
// 		if err != nil {
// 			logrus.Error("error updating user session")
// 			http.Error(w, "error updating user session", http.StatusInternalServerError)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-type", "application/json")
// 		_, err = w.Write([]byte(fmt.Sprintf(`{"session_id":"%s"}`, sessionId)))
// 		if err != nil {
// 			logrus.Error("error wrting response")
// 			http.Error(w, "error wrting response", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

// //add a payload validator
// func UpdateUser(db *sqlx.DB, c *cache.Cache) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var user map[string]interface{}
// 		err := json.NewDecoder(r.Body).Decode(&user)
// 		if err != nil {
// 			logrus.Errorf("error reading request body: %v", err)
// 			http.Error(w, "error reading request body", http.StatusBadRequest)
// 			return
// 		}

// 		statementBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
// 		update:= statementBuilder.Update("users").SetMap(user).
// 		Set("updated_at", time.Now()).
// 		Where(sq.Eq{"session_id": chi.URLParam(r, "id")})
// 		_, err = update.Exec()
// 		if err != nil {
// 			logrus.Errorf("error updating user: %v", err.Error())
// 			http.Error(w, "error updating user", http.StatusInternalServerError)
// 			return
// 		}

// 		var updatedUser User
// 		err = db.Get(&updatedUser, fmt.Sprintf(`Select * from users where session_id = '%s'`, chi.URLParam(r,"id")))
// 		if err!= nil {
// 			logrus.Errorf("error fetching user from db :%v", err.Error())
// 			http.Error(w, "error fetching user from database", http.StatusInternalServerError)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-type", "application/json")
// 		err = json.NewEncoder(w).Encode(&updatedUser)
// 		if err != nil {
// 			logrus.Error("error writing response")
// 			http.Error(w, "error writing response", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

// func AddFollower(db *sqlx.DB, c *cache.Cache) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "id")
// 		// Decode the request body into a struct
// 		type payload struct {Followed string `json:"followed" db:"followed"`}
// 		var AddUser payload
// 	 	err := json.NewDecoder(r.Body).Decode(&AddUser)
// 		if err != nil {
// 			logrus.Errorf("error reading request body: %v", err)
// 			http.Error(w, "error reading request body", http.StatusBadRequest)
// 			return
// 		}

// 		//retreive the current user
// 		var userSharekey map[string]string
// 		err = db.Get(&userSharekey, fmt.Sprintf(`Select share_key from users where session_id = '%s'`, id))
// 		if err!= nil {
//             logrus.Errorf("error fetching user from db :%v", err.Error())
//             http.Error(w, "error fetching user from database", http.StatusInternalServerError)
//             return
//         }
// 		_, err = db.Exec(fmt.Sprintf(`insert into followers(follower, followed) VALUES (%s, %s)`, userSharekey["share_key"], AddUser.Followed))
// 		if err != nil {
// 			logrus.Errorf("error adding follower to database: %v", err)
// 			http.Error(w, "error adding follower to database", http.StatusInternalServerError)
// 			return
// 		}

// 		// Retrieve the new followers count
// 		var newFollowersCount int
// 		err = db.Get(&newFollowersCount, fmt.Sprintf(`SELECT COUNT(*) FROM followers WHERE followed = '%s'`, AddUser.Followed))
// 		if err != nil {
// 			logrus.Errorf("error retrieving followers count: %v", err)
// 			http.Error(w, "error retrieving followers count", http.StatusInternalServerError)
// 			return
// 		}

// 		// Return a success response
// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-Type", "application/json")
// 		response := map[string]interface{}{
// 			"message":           "Follower added successfully",
// 			"new_followers_count": newFollowersCount,
// 		}
// 		err = json.NewEncoder(w).Encode(response)
// 		if err != nil {
// 			logrus.Errorf("error writing response: %v", err)
// 			http.Error(w, "error writing response", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }


// func RemoveFollower(db *sqlx.DB, c *cache.Cache) http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {
		
// 	}
// }