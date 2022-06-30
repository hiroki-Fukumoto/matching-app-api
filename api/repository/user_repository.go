package repository

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Create(request *request.CreateUserRequest) (user *model.User, err error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (up *userRepository) Create(req *request.CreateUserRequest) (user *model.User, err error) {
	db := up.DB

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	user = &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
