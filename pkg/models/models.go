package main

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Login struct {
	Email    null.String `json:"email"`
	Username null.String `json:"username"`
	Password string      `json:"password"`
}

type User struct {
	Key        string       `json:"share_key" db:"share_key"`
	Username   string       `json:"username" db:"username" required:"true"`
	Email      string       `json:"email" db:"email" required:"true"`
	Password   string       `json:"password" db:"password" required:"true"`
	FullName   null.String  `json:"fullname" db:"fullname"`
	Bio        null.String  `json:"bio" db:"bio"`
	WhoAreYou  null.String  `json:"whoareyou" db:"whoareyou"`
	Avatar     null.Int     `json:"avatar" db:"avatar"`
	AccountRating float64   `json:"account_rating" db:"account_rating"`
	PostRating float64      `json:"post_rating" db:"post_rating"`
	FollowerCount int       `json:"follower_count" db:"follower_count"`
	FollowingCount int     `json:"following_count" db:"following_count"`
	PostCount  int          `json:"post_count" db:"post_count"`
	Location   null.String `json:"location" db:"location"`
	SessionId  null.String `json:"session_id" db:"session_id"`
	Birthday   string      `json:"birthday" db:"birthday" required:"true"`
	UpdatedAt  string      `json:"updated_at" db:"updated_at"`
	CreatedAt  string      `json:"created_at" db:"created_at"`
	Active     bool         `json:"active" db:"active"`
}

type Followers struct {
	Follower string `json:"follower" db:"follower"`
	Followed string `json:"followed" db:"followed"`
}

type Post struct {
	ID        int          `json:"id" db:"id"`
	UserID    int          `json:"user_id" db:"user_id"`
	Title     string       `json:"title" db:"title"`
	Content   string       `json:"content" db:"content"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
}

type Comment struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	PostID    int       `json:"post_id" db:"post_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Like struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	PostID    int       `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Message struct {
	ID        int       `json:"id" db:"id"`
	SenderID  int       `json:"sender_id" db:"sender_id"`
	ReceiverID int      `json:"receiver_id" db:"receiver_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

