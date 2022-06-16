package models

import "time"

type User struct {
	Id          int        `json:"id"`
	Email       string     `json:"email"`
	Password    *string    `json:"password"`
	AccessToken *string    `json:"access_token"`
	ExpiredAt   *time.Time `json:"expired_at"`
}
