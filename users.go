package main

import (
	"gopkg.in/guregu/null.v3"
	"time"
)

type UserLogin struct {
	Email    null.String `json:"email"`
	Username null.String `json:"username"`
	Password string      `json:"password"`
}

type UserSession struct {
	SessionId      string    `json:"session_id"`
	LastAccessTime time.Time `json:"last_access"`
	Username       string    `json:"username"`
}

type UserProfile struct {
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	FullName    null.String `json:"full_name"`
	Bio         null.String `json:"bio"`
	Links       null.String `json:"links"`
	Occupation  null.String `json:"occupation"`
	UserProfile null.Int    `json:"profile_pic"`
	Rating      string      `json:"rating"`
	City        string      `json:"city"`
	State       null.String `json:"state"`
	SessionId   null.String `json:"session_id"`
	Birthday    null.Time   `json:"birthday"`
}
