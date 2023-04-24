package main

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
