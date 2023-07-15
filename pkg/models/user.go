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
	UserId          string    `json:"user_id" db:"user_id"`
	Username        *string    `json:"username" db:"username"`
	Email           *string    `json:"email" db:"email"`
	Password        string    `json:"user_password" db:"user_password"`
	FirstName       string    `json:"first_name" db:"first_name"`
	LastName        string    `json:"last_name" db:"last_name"`
	Bio             *string    `json:"bio" db:"bio" `
	Avatar          []byte    `json:"avatar" db:"avatar"` // Store photo as []byte or use a separate photo storage solution
	AccountRating   float64   `json:"account_rating" db:"account_rating" `
	FollowerCount   int       `json:"follower_count" db:"follower_count" `
	FollowingCount  int       `json:"following_count" db:"following_count"`
	Latitude        *string    `json:"latitude" db:"latitude"`  
	Longitude       *string    `json:"longitude" db:"longitude"` 
	Birthday        *string    `json:"birthday" db:"birthday"`
	City            *string    `json:"city" db:"city"`
	Country         *string    `json:"country" db:"country"`
	StateProvince   *string    `json:"state_province" db:"state_province"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	Active          bool      `json:"active" db:"active"`
	Verified        bool      `json:"verified" db:"verified"`
}

var CreateUserValidation = map[string]string{
	"UserId" : "required,uuid",
	"Username" : "required,max=225",
	"Email" : "required,email",
	"Password " : "required,min=10",
	"FirstName " : "required,max=225",
	"LastName" : "required,max=225",
	"UpdatedAt" : "required",
	"CreatedAt" : "required",
}

var UserLoginValidation = map[string]string{
	"Username" : "omitempty,max=225",
	"Email" : "omitempty,email",
	"Password" : "required,min=10",
}
