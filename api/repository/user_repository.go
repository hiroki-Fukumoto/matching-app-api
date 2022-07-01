package repository

import (
	"errors"

	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Create(request *request.CreateUserRequest) (user *model.User, err error)
	FindByEmail(email string) (user *model.User, err error)
	FindByID(id string) (user *model.User, err error)
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
		Sex:      req.Sex,
		Birthday: util.ParseDate(req.Birthday),
		Password: passwordHash,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (up *userRepository) FindByEmail(email string) (user *model.User, err error) {
	db := up.DB

	err = db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error_handler.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (up *userRepository) FindByID(id string) (user *model.User, err error) {
	db := up.DB

	err = db.Where("id = ?", id).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error_handler.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
