package impl

import (
	"errors"
	"fmt"
	"log"
	"res-gin/src/dto"
	"res-gin/src/model"
	"res-gin/src/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

	userID := uuid.New().String()
	userQuery := `INSERT INTO "users" ("id", "email", "username", "password", "created_at", "updated_at") 
								VALUES (?, ?, ?, ?, ?, ?) RETURNING "id"`
	if err := tx.Raw(userQuery, userID, userDto.Email, userDto.Username, string(hashedPassword), time.Now(), time.Now()).Scan(&userID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	log.Println("User inserted successfully, ID:", userID)

	var roleID string
	roleQuery := `SELECT "id" FROM "roles" WHERE "role" = ? LIMIT 1`
	err = tx.Raw(roleQuery, userDto.Role).Scan(&roleID).Error

	if err != nil || roleID == "" {
		roleID = uuid.New().String()
		insertRoleQuery := `INSERT INTO "roles" ("id", "role", "created_at", "updated_at") 
												VALUES (?, ?, ?, ?) RETURNING "id"`
		if err := tx.Raw(insertRoleQuery, roleID, userDto.Role, time.Now(), time.Now()).Scan(&roleID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		log.Println("Role inserted successfully, ID:", roleID)
	} else {
		log.Println("Role already exists, ID:", roleID)
	}

	userRoleQuery := `INSERT INTO "user_roles" ("users_id", "role_id") 
										VALUES (?, ?) ON CONFLICT DO NOTHING`
	if err := tx.Raw(userRoleQuery, userID, roleID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	log.Println("User-role relationship inserted successfully")

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &model.Users{
		ID:       userID,
		Email:    userDto.Email,
		Username: userDto.Username,
		Password: string(hashedPassword),
		Roles:    []model.Role{{ID: roleID, Role: userDto.Role}},
	}, nil
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
