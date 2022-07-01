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
	Create(req *request.CreateUserRequest) (user *model.User, err error)
	FindAll(req *request.SearchUserRequest) (users []*model.User, err error)
	FindByEmail(email string) (user *model.User, err error)
	FindByID(id string) (user *model.User, err error)
	FindPickUpToday(sex string) (users []*model.User, err error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (ur *userRepository) Create(req *request.CreateUserRequest) (user *model.User, err error) {
	db := ur.DB

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	user = &model.User{
		Name:       req.Name,
		Email:      req.Email,
		Sex:        req.Sex,
		Birthday:   util.ParseDate(req.Birthday),
		Prefecture: uint16(req.Prefecture),
		Password:   passwordHash,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindAll(req *request.SearchUserRequest) (users []*model.User, err error) {
	db := ur.DB

	q := db
	if req.Prefecture != nil {
		q = q.Where("prefecture = ?", req.Prefecture)
	}
	if req.FromAge != nil {
		b := util.CalcBirthdayMonthFromAge(*req.FromAge)
		q = q.Where("birthday <= ?", b)
	}
	if req.ToAge != nil {
		b := util.CalcBirthdayMonthFromAge(*req.ToAge)
		q = q.Where("birthday >= ?", b)
	}
	if req.Sort != nil {
		q = q.Order("created_at desc")
	}
	l := 20
	if err := q.Limit(l).Offset((req.Page - 1) * l).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) FindByEmail(email string) (user *model.User, err error) {
	db := ur.DB

	err = db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error_handler.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindByID(id string) (user *model.User, err error) {
	db := ur.DB

	err = db.Where("id = ?", id).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error_handler.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindPickUpToday(sex string) (users []*model.User, err error) {
	db := ur.DB

	if err := db.Where("sex = ?", sex).Limit(20).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
