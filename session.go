package main

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func CreateSession(w http.ResponseWriter, r *http.Request, user string, l *zap.SugaredLogger) (string, error) {
	// ! add error handling
	session_id := uuid.New().String()
	sessionCookie := &http.Cookie{
		Name:    "session_id",
		Value:   session_id,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
	}
	http.SetCookie(w, sessionCookie)
	l.Infof("Session created for user %s on %s", user, string(time.Now().Format("2006-01-02 15:04:05 MST")))
	return session_id, nil
}

func GetSession(w http.ResponseWriter, r *http.Request, l *zap.SugaredLogger) (*UserSession, error) {
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
		return nil, nil
	}
	//create a user session instance and assign the found cookies for return
	user := &UserSession{}
	user.Username = cookie.Value
	user.SessionId = session.Value
	return user, nil
}
