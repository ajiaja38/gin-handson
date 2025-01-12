package impl

import (
	"errors"
	"fmt"
	"log"
	"res-gin/src/dto"
	"res-gin/src/model"
	"res-gin/src/service"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db          *gorm.DB
	validate    *validator.Validate
	roleService service.RoleService
}

func NewUserService(db *gorm.DB) service.UserService {
	roleService := NewRoleServiceImpl(db)

	return &UserServiceImpl{
		db:          db,
		validate:    validator.New(),
		roleService: roleService,
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

	tx := s.db.Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user := &model.Users{
		Email:    userDto.Email,
		Username: userDto.Username,
		Password: string(hashedPassword),
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	role, err := s.roleService.GetOrSaveRole(userDto.Role)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&user).Association("Roles").Append(role); err != nil {
		log.Println("Error appending role:", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]model.Users, error) {
	var users []model.Users

	if err := s.db.Raw("SELECT * FROM users").Scan(&users).Error; err != nil {
		return nil, err
	}

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

func (s *UserServiceImpl) DeleteUsers(id string) error {
	result := s.db.Delete(&model.Users{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("User with id %s not found", id)
	}

	return nil
}
