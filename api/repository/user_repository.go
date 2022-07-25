package repository

import (
	"errors"
	"time"

	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	CreateRequest() *CreateRequest
	Create(req *CreateRequest) (user *model.User, err error)
	FindAllRequest() *FindAllRequest
	FindAll(req *FindAllRequest) (users []*model.User, err error)
	FindByEmailRequest() *FindByEmailRequest
	FindByEmail(req *FindByEmailRequest) (user *model.User, err error)
	FindByIDRequest() *FindByIDRequest
	FindByID(req *FindByIDRequest) (user *model.User, err error)
	FindPickUpTodayRequest() *FindPickUpTodayRequest
	FindPickUpToday(req *FindPickUpTodayRequest) (users []*model.User, err error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

type CreateRequest struct {
	Name       string
	Email      string
	Sex        string
	Birthday   time.Time
	Prefecture uint16
	Password   []byte
}

func (ur *userRepository) CreateRequest() *CreateRequest {
	return &CreateRequest{}
}

func (ur *userRepository) Create(req *CreateRequest) (user *model.User, err error) {
	db := ur.DB

	user = &model.User{
		Name:       req.Name,
		Email:      req.Email,
		Sex:        req.Sex,
		Birthday:   req.Birthday,
		Prefecture: req.Prefecture,
		Password:   req.Password,
		Avatar:     "https://placehold.jp/300x300.png",
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

type FindAllRequest struct {
	Page       int
	FromAge    *int
	ToAge      *int
	Prefecture *int
	Sort       *int
}

func (ur *userRepository) FindAllRequest() *FindAllRequest {
	return &FindAllRequest{}
}

func (ur *userRepository) FindAll(req *FindAllRequest) (users []*model.User, err error) {
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

type FindByEmailRequest struct {
	Email string
}

func (ur *userRepository) FindByEmailRequest() *FindByEmailRequest {
	return &FindByEmailRequest{}
}

func (ur *userRepository) FindByEmail(req *FindByEmailRequest) (user *model.User, err error) {
	db := ur.DB

	err = db.Where("email = ?", req.Email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error_handler.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

type FindByIDRequest struct {
	ID string
}

func (ur *userRepository) FindByIDRequest() *FindByIDRequest {
	return &FindByIDRequest{}
}

func (ur *userRepository) FindByID(req *FindByIDRequest) (user *model.User, err error) {
	db := ur.DB

	err = db.Where("id = ?", req.ID).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error_handler.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

type FindPickUpTodayRequest struct {
	Sex string
}

func (ur *userRepository) FindPickUpTodayRequest() *FindPickUpTodayRequest {
	return &FindPickUpTodayRequest{}
}

func (ur *userRepository) FindPickUpToday(req *FindPickUpTodayRequest) (users []*model.User, err error) {
	db := ur.DB

	if err := db.Where("sex = ?", req.Sex).Limit(20).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
