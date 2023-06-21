package models

import (
	"time"
	//"github.com/go-playground/validator/v10"
	"gopkg.in/guregu/null.v3"
)

type Login struct {
	Email    null.String `json:"email"`
	Username null.String `json:"username"`
	Password string      `json:"password"`
}

type Request struct {
	TraceId string
	From string
}

type Session struct {
	Token string
	Expiration string
}

type User struct {
	UserId 				string		 `json:"user_id" db:"user_id" validate:"required,uuid"` 
	Username   			string       `json:"username" db:"username" validate:max=225`
	Email      			string       `json:"email" db:"email" validate:"email"`
	Password   			string       `json:"user_password" db:"User_password" validate:"required"`
	FirstName  			string 		 `json:"first_name" db:"first_name" validate:"max=255"`
	LastName   			string  	 `json:"last_name" db:"last_name" validate:"max=255"`
	Bio        			string  	 `json:"bio" db:"bio" validate:"max=255"`
	Avatar     			float64      `json:"avatar" db:"avatar" validate:"binary"` // TODO: double check if this is allowed
	AccountRating 		float64   	 `json:"account_rating" db:"account_rating" validate:"number"`
	RunningPointCount 	float64      `json:"running_point_count" db:"running_point_count" validate:"number"` // TODO: double check if this is allowed
	FollowerCount  		int       	 `json:"follower_count" db:"follower_count" validate:"number"` // TODO: double check if this is allowed
	FollowingCount 		int   		 `json:"following_count" db:"following_count" validate:"number"` // TODO: double check if this is allowed
	Location     		string 		 `json:"location" db:"location"`
	Birthday   			string       `json:"birthday" db:"birthday" validate:"datetime"`
	City    			string 		 `json:"city" db:"city"`
	Country    			string 		 `json:"country" db:"country"`
	StateProvince    	string 		 `json:"state_province" db:"state_province"`
	UpdatedAt  			time.Time    `json:"updated_at" db:"updated_at"`
	CreatedAt  			time.Time    `json:"created_at" db:"created_at"`
	Active     			bool         `json:"active" db:"active"`
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

