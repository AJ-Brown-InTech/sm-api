package models

type Followers struct {
	Follower string `json:"follower" db:"follower"`
	Followed string `json:"followed" db:"followed"`
}
