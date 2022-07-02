package models

type Auth struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   string `json:"expired_at"`
}
