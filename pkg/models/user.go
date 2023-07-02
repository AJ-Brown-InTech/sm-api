package models

import (
	"time"
	//"gopkg.in/guregu/null.v3"
)

type Session struct {
	Token string
	Expiration time.Time
}

type User struct {
	UserId          string    `json:"user_id" db:"user_id" validate:"required,uuid"`
	Username        string    `json:"username" db:"username" validate:"max=225"`
	Email           string    `json:"email" db:"email" validate:"email"`
	Password        string    `json:"user_password" db:"user_password" validate:"required,min=12"`
	FirstName       string    `json:"first_name" db:"first_name" validate:"max=255"`
	LastName        string    `json:"last_name" db:"last_name" validate:"max=255"`
	Bio             *string    `json:"bio" db:"bio" validate:"max=255"`
	Avatar          []byte    `json:"avatar" db:"avatar"` // Store photo as []byte or use a separate photo storage solution
	AccountRating   float64   `json:"account_rating" db:"account_rating" validate:"number"`
	FollowerCount   int       `json:"follower_count" db:"follower_count" validate:"number"`
	FollowingCount  int       `json:"following_count" db:"following_count" validate:"number"`
	Latitude        *string    `json:"latitude" db:"latitude"`  
	Longitude       *string    `json:"longitude" db:"longitude"` 
	Birthday        *string    `json:"birthday" db:"birthday" validate:"datetime"`
	City            *string    `json:"city" db:"city"`
	Country         *string    `json:"country" db:"country"`
	StateProvince   *string    `json:"state_province" db:"state_province"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	Active          bool      `json:"active" db:"active"`
	Verified        bool      `json:"verified" db:"verified"`
}
