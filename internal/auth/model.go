package auth

import "time"

type User struct {
	Name    string
	Mail    string
	Updated time.Time
}

type UserLogin struct {
	Name         string
	Mail         string
	PasswordHash string
	Updated      time.Time
}

type Session struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
