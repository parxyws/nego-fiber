package dto

type UserResponse struct {
	UserID      string   `json:"user_id"`
	Username    string   `json:"username"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	Avatar      string   `json:"avatar,omitempty"`
	PhoneNumber string   `json:"phone_number,omitempty"`
	Jwt         JwtToken `json:"jwt,omitempty"`
}

type JwtToken struct {
	AccessToken  string `json:"access_Token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRegisterResponse struct {
	ReferenceID string `json:"reference_id"`
}
