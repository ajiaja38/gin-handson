package impl

import (
	"errors"
	"os"
	"res-gin/src/dto"
	"res-gin/src/model"
	"res-gin/src/service"
	"res-gin/src/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	db       *gorm.DB
	vaildate *validator.Validate
}

var (
	accessTokenSecretKey  []byte = []byte(os.Getenv("SECRET_ACCESS_TOKEN"))
	refreshTokenSecretKey []byte = []byte(os.Getenv("SECRET_REFRESH_TOKEN"))
	ErrUserNotFound       error  = errors.New("Invalid Username or Password")
)

func NewAuthServiceImpl(db *gorm.DB) service.AuthService {
	return &AuthServiceImpl{
		db:       db,
		vaildate: validator.New(),
	}
}

func (s *AuthServiceImpl) LoginUser(loginDto *dto.LoginDTO) (*dto.LoginResponseDTO, error) {
	err := s.vaildate.Struct(loginDto)

	if err != nil {
		return nil, err
	}

	var user model.Users

	if err := s.db.Where("email =?", loginDto.Email).First(&user).Error; err != nil {
		return nil, ErrUserNotFound
	}

	if err := s.verifyPassword(user.Password, loginDto.Password); err != nil {
		return nil, ErrUserNotFound
	}

	accessToken, err := utils.GenerateToken(user.ID, time.Hour, accessTokenSecretKey)

	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(user.ID, time.Hour*24*7, refreshTokenSecretKey)

	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *AuthServiceImpl) RefreshToken(refreshTokenDto *dto.RefreshTokenDTO) (*dto.RefreshTokenResponseDTO, error) {
	err := s.vaildate.Struct(refreshTokenDto)

	if err != nil {
		return nil, err
	}

	claims := &model.UserClaims{}

	err = utils.ValidateToken(refreshTokenDto.RefreshToken, refreshTokenSecretKey, claims)

	if err != nil {
		return nil, err
	}

	newAccessToken, err := utils.GenerateToken(claims.UserID, time.Hour, accessTokenSecretKey)

	if err != nil {
		return nil, err
	}

	response := &dto.RefreshTokenResponseDTO{
		AccessToken: newAccessToken,
	}

	return response, nil
}

func (s *AuthServiceImpl) verifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
