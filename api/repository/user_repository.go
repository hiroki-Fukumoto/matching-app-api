package repository

import (
	"errors"
	"log"
	"strconv"
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
	WithTrx(*gorm.DB) *userRepository
	CreateRequest() *CreateRequest
	Create(req *CreateRequest) (user *model.User, err error)
	UpdateRequest() *UpdateRequest
	Update(id string, req *UpdateRequest) (user *model.User, err error)
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

func (u *userRepository) WithTrx(trxHandle *gorm.DB) *userRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
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

type UpdateRequest struct {
	Name       string
	Prefecture uint16
	Hobbies    []string
}

func (ur *userRepository) UpdateRequest() *UpdateRequest {
	return &UpdateRequest{}
}

func (ur *userRepository) Update(id string, req *UpdateRequest) (user *model.User, err error) {
	db := ur.DB

	err = db.Where("id = ?", id).First(&user).
		Updates(model.User{Name: req.Name, Prefecture: req.Prefecture}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, error_handler.ErrRecordNotFound
	}

	if len(req.Hobbies) > 0 {
		if err := db.Where("user_id = ?", id).Delete(&model.UserHobby{}).Error; err != nil {
			return nil, err
		}
	}

	for _, h := range req.Hobbies {
		m := &model.UserHobby{
			UserID:  id,
			HobbyID: h,
		}

		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
	}

	err = db.Preload("Hobbies").First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

type FindAllRequest struct {
	Page        int
	FromAge     *int
	ToAge       *int
	Prefectures *[]int
	Sort        *int
}

func (ur *userRepository) FindAllRequest() *FindAllRequest {
	return &FindAllRequest{}
}

func (ur *userRepository) FindAll(req *FindAllRequest) (users []*model.User, err error) {
	db := ur.DB

	q := db
	if req.Prefectures != nil && len(*req.Prefectures) > 0 {
		var prefs []string
		for _, p := range *req.Prefectures {
			prefs = append(prefs, strconv.Itoa(p))
		}
		q = q.Where("prefecture IN ?", prefs)
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
	if err := q.Preload("Hobbies").Limit(l).Offset((req.Page - 1) * l).Find(&users).Error; err != nil {
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

	err = db.Preload("Hobbies").Where("email = ?", req.Email).First(&user).Error

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

	err = db.Preload("Hobbies").Where("id = ?", req.ID).First(&user).Error

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

	if err := db.Preload("Hobbies").Where("sex = ?", req.Sex).Limit(20).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
