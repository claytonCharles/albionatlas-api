package auth

type RegisterForm struct {
	Name     string `form:"name" validate:"required"`
	Mail     string `form:"mail" validate:"required|mail"`
	Password string `form:"password" validate:"required|min=8"`
}
