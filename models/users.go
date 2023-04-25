package maodels

import (
	"time"
)

type UserLogin struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	SessionId      string    `json:"session_id"`
	LastAccessTime time.Time `json:"last_access"`
	Username       string    `json:"username"`
}

type UserProfile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName  string `json:"full_name"`
	Bio string `json:"bio"`
	Links string `json:"links"`
	Occupation string `json:"occupation"`
	UserProfile string `json:"profile_pic"`
	Rating string `json:"rating"`
	City string `json:"city"`
	State string `json:"state"`
	SessionId string `json:"session_id"`
	Birthday string `json:"birthday"`
}