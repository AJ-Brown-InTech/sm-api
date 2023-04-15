package session

import (
	"net/http"
	"time"
	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	"go.uber.org/zap"
)

type Session struct{
	UserId string `json:"user_id"`
	LastAccessTime time.Time `json:"last_access"`
}

type UserInfo struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Phone string `json:"phone"`

}

func CreateSession(w http.ResponseWriter, r *http.Request, user *UserInfo, log *zap.SugaredLogger) error {
	
		var instance = securecookie.New([]byte(uuid.New().String()),securecookie.GenerateRandomKey(32))
		userData := make(map[string]string)
		userData["username"] = user.Username
		userData["email"] = user.Email
		cookieValue, _ := instance.Encode("user", userData)
		//Todo: throw error and do error handling
	cookie := &http.Cookie{
		    Name: "user",
            Value: cookieValue,
            Path: "/",
            Expires: time.Now().Add(time.Hour * 24),
            HttpOnly: true,
            SameSite: http.SameSiteStrictMode,
        }
        http.SetCookie(w, cookie)
        return nil
}


func GetSession(w http.ResponseWriter, r *http.Request, log *zap.SugaredLogger) (*UserInfo, error) {
	    cookie, err := r.Cookie("user")
        if err!= nil {
          //Todo: throw error and do error handling
        }
        if cookie.Value == "" {
           //Todo: throw error and do error handling
        }
		var instance = securecookie.New([]byte(uuid.New().String()), nil)
		var userData = &UserInfo{}
		 err = instance.Decode("user",cookie.Value,userData)
		if err != nil{
			 //Todo: throw error and do error handling
		} 
		user := userData
		return user , nil
}
