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
	Username    string      `json:"username" db:"username"`
	Email       string      `json:"email" db:"email"`
	Password    string      `json:"password" db:"password"`
	FullName    null.String `json:"full_name" db:"fullname"`
	Bio         null.String `json:"bio" db:"bio"`
	Links       null.String `json:"links" db:"links"`
	Occupation  null.String `json:"occupation" db:"occupation"`
	UserProfile null.Int    `json:"profile_pic" db:"user_profile"`
	Rating      string      `json:"rating" db:"rating"`
	City        null.String `json:"city" db:"city"`
	State       null.String `json:"state" db:"state"`
	SessionId   null.String `json:"session_id" db:"session_id"`
	Birthday    null.Time   `json:"birthday" db:"birthday"`
}
