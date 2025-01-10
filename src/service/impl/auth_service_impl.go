package impl

import (
	"res-gin/src/dto"
	"res-gin/src/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	db       *gorm.DB
	vaildate *validator.Validate
}

func NewAuthServiceImpl(db *gorm.DB) service.AuthService {
	return &AuthServiceImpl{
		db:       db,
		vaildate: validator.New(),
	}
}

// LoginUser implements service.AuthService.
func (a *AuthServiceImpl) LoginUser(loginDto *dto.LoginDTO) (*dto.LoginResponseDTO, error) {
	panic("unimplemented")
}
