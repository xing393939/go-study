package models

import (
	_ "github.com/lib/pq"
)

type User struct {
	Id                  string `json:"id" binding:"required"`
	Username            string `json:"username" binding:"required"`
	Pic                 string `json:"pic"`
	TwitterAccessToken  string `json:"-" binding:"required"`
	TwitterAccessSecret string `json:"-" binding:"required"`
	TimeCreated         string `json:"timeCreated"`
	TimeUpdated         string `json:"timeUpdated"`
	IsAdmin             bool   `json:"isAdmin"`
}

func (db *DB) GetUserWithId(id string) (*User, error) {
	var user User
	err := db.SQLDB.QueryRow(`
        SELECT id, username, pic, twitter_access_token, twitter_access_secret,
               time_created, time_updated
        FROM auth_user WHERE id = ?`, id).Scan(
		&user.Id,
		&user.Username,
		&user.Pic,
		&user.TwitterAccessToken,
		&user.TwitterAccessSecret,
		&user.TimeCreated,
		&user.TimeUpdated,
	)
	return &user, err
}
