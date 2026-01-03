package auth

import (
	"encoding/json"
	"net/http"

	"github.com/claytonCharles/albionatlas-api/pkg/validator"
)

type authHandlerImp struct {
	service AuthService
}

func NewHandler(repo AuthRepository) AuthHandler {
	service := NewService(repo)
	return &authHandlerImp{
		service: service,
	}
}

func (ahi *authHandlerImp) Login(w http.ResponseWriter, r *http.Request)  {}
func (ahi *authHandlerImp) Logout(w http.ResponseWriter, r *http.Request) {}

func (ahi *authHandlerImp) Register(w http.ResponseWriter, r *http.Request) {
	var form RegisterForm

	errs := validator.FormValidate(r, &form)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	err := ahi.service.CreateUser(form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Usuário criado com sucesso!"))
}

func (ahi *authHandlerImp) RefreshJWT(w http.ResponseWriter, r *http.Request) {}
