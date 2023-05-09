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
    Username       string      `json:"username" db:"username" required:"true"`
	Email          string      `json:"email" db:"email" required:"true"`
	Password       string      `json:"password" db:"password" required:"true"`
	FullName       null.String `json:"full_name" db:"fullname"`
	Bio            null.String `json:"bio" db:"bio"`
	Occupation     null.String `json:"occupation" db:"occupation"`
	Avatar         null.Int    `json:"avatar" db:"avatar"`
	Rating         float64     `json:"rating" db:"rating"`
	City           null.String `json:"city" db:"city"`
	SessionId      null.String `json:"session_id" db:"session_id"`
	Birthday       time.Time   `json:"birthday" db:"birthday" required:"true"`
	UpdatedAt      time.Time   `json:"last_access" db:"last_access"`
	CreatedAt      time.Time   `json:"created_at" db:"created_at"`
}
