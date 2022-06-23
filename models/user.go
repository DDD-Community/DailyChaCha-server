package models

import "time"

type User struct {
	Id          int
	Email       string  `json:"email"`
	Password    *string `json:"password"`
	AccessToken *string
	ExpiredAt   *time.Time
}

type Auth struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   string `json:"expired_at"`
}
