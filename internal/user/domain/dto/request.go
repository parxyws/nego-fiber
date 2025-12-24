package dto

type UserRegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,max=100,alpha"`
	LastName  string `json:"last_name" validate:"required,max=100"`
	Username  string `json:"username" validate:"required,max=100,alpha"`
	Email     string `json:"email" validate:"required,max=100,email"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,max=100,alpha"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}
