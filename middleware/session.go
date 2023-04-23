package session

import (
	"net/http"
	"time"

	"github.com/AJ-Brown-InTech/sm-api/models"
	"github.com/google/uuid"
	//"github.com/gorilla/securecookie"
	"go.uber.org/zap"
)


func CreateSession(w http.ResponseWriter, r *http.Request, user *models.UserLogin, l *zap.SugaredLogger) error {
	session_id := uuid.New().String()
	//set user cookie
	userCookie := &http.Cookie{
		Name:       "user",
		Value:      user.Username,
		Path:       "/",
		Expires:    time.Now().Add(time.Hour * 24),
	}
	http.SetCookie(w, userCookie)
	
	// set session cookie
	sessionCookie := &http.Cookie{
		Name:       "session_id",
		Value:      session_id,
		Path:       "/",
		Expires:    time.Now().Add(time.Hour * 24),
	}
	http.SetCookie(w, sessionCookie)
	
	//store session_id in database
	//TODO: store sessionId in database	
	
	l.Infof("Session created for user %s on %s", user.Username, string(time.Now().Format("2006-01-02 15:04:05 MST")))
	return nil
}

func GetSession(w http.ResponseWriter, r *http.Request, l *zap.SugaredLogger) (*models.UserSession, error) {
	//TODO: get both cookies in one call 
	//get user cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		l.Errorf("Not found user cookie")
		return nil, err
	}
	//get session cookie
	session, err := r.Cookie("session_id")
	if err != nil {
		l.Errorf("Not found session cookie")
		return nil, err
	}
	//if any of the values are empty throw an error
	if cookie.Value == "" || session.Value == "" {
		l.Errorf("Cookie/Session value is empty")
		return nil,nil
	}
	//create a user session instance and assign the found cookies for return
	user := &models.UserSession{}
	user.Username = cookie.Value
	user.SessionId = session.Value
	return user, nil
}

