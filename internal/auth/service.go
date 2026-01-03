package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type authServiceImp struct {
	repository AuthRepository
}

func NewService(rp AuthRepository) AuthService {
	return &authServiceImp{
		repository: rp,
	}
}

func (rp *authServiceImp) CreateUser(form RegisterForm) error {
	if ok := rp.repository.CheckMailExists(form.Mail); !ok {
		return errors.New("Mail already exists!")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)
	if err != nil {
		return ErrFail
	}

	form.Password = string(password)
	if err := rp.repository.CreateUser(form); err != nil {
		return err
	}

	return nil
}
