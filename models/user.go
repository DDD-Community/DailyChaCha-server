package models

import "time"

type User struct {
	Id                    int        `json:"id"`
	Email                 string     `json:"email"`
	IsOnboardingCompleted bool       `json:"is_onboarding_completed"`
	Password              *string    `json:"password"`
	AccessToken           *string    `json:"access_token"`
	ExpiredAt             *time.Time `json:"expired_at"`
}

type Auth struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   string `json:"expired_at"`
}
