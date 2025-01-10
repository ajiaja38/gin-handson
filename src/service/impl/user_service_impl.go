package impl

import (
	"errors"
	"fmt"
	"res-gin/src/dto"
	"res-gin/src/model"
	"res-gin/src/service"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func NewUserService(db *gorm.DB) service.UserService {
	return &UserServiceImpl{
		db:       db,
		validate: validator.New(),
	}
}

func (s *UserServiceImpl) CreateUser(userDto *dto.CreateUserDTO) (*model.Users, error) {
	err := s.validate.Struct(userDto)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &model.Users{
		Email:    userDto.Email,
		Username: userDto.Username,
		Password: string(hashedPassword),
		Role:     userDto.Role,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]model.Users, error) {
	var users []model.Users

	err := s.db.Raw("SELECT * FROM users").Scan(&users).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(users)

	return users, nil
}

func (s *UserServiceImpl) GetUserById(id string) (*model.Users, error) {
	var user model.Users

	result := s.db.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)

	if result.RowsAffected == 0 {
		return nil, errors.New("User Not Found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *UserServiceImpl) VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
