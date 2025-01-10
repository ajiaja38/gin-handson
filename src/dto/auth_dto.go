package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,"`
	Password string `json:"password" validate:"required,"`
}

type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
