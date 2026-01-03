package auth

import "net/http"

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	RefreshJWT(w http.ResponseWriter, r *http.Request)
}

type AuthService interface {
	CreateUser(form RegisterForm) error
}

type AuthRepository interface {
	CheckMailExists(mail string) bool
	CreateUser(form RegisterForm) error
}
