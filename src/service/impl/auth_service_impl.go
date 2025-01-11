package impl

import (
	"errors"
	"os"
	"res-gin/src/dto"
	"res-gin/src/model"
	"res-gin/src/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	db       *gorm.DB
	vaildate *validator.Validate
}

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewAuthServiceImpl(db *gorm.DB) service.AuthService {
	return &AuthServiceImpl{
		db:       db,
		vaildate: validator.New(),
	}
}

func (s *AuthServiceImpl) LoginUser(loginDto *dto.LoginDTO) (*dto.LoginResponseDTO, error) {
	accessTokenSecretKey := []byte(os.Getenv("SECRET_ACCESS_TOKEN"))
	refreshTokenSecretKey := []byte(os.Getenv("SECRET_REFRESH_TOKEN"))

	err := s.vaildate.Struct(loginDto)

	if err != nil {
		return nil, err
	}

	var user model.Users

	if err := s.db.Where("email =?", loginDto.Email).First(&user).Error; err != nil {
		return nil, errors.New("Invalid Username or Password")
	}

	if err := s.verifyPassword(user.Password, loginDto.Password); err != nil {
		return nil, errors.New("Invalid Username or Password")
	}

	accessToken, err := s.generateToken(user.ID, time.Hour, accessTokenSecretKey)

	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateToken(user.ID, time.Hour*24*7, refreshTokenSecretKey)

	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *AuthServiceImpl) generateToken(id string, duration time.Duration, secret []byte) (string, error) {
	claims := &UserClaims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func (s *AuthServiceImpl) verifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
