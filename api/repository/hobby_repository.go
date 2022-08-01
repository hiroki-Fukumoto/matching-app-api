package repository

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"gorm.io/gorm"
)

type hobbyRepository struct {
	DB *gorm.DB
}

type HobbyRepository interface {
	FindAll() (hobbies []*model.Hobby, err error)
}

func NewHobbyRepository(db *gorm.DB) HobbyRepository {
	return &hobbyRepository{DB: db}
}

func (hr *hobbyRepository) FindAll() (hobbies []*model.Hobby, err error) {
	db := hr.DB
	if err := db.Find(&hobbies).Error; err != nil {
		return nil, err
	}
	return hobbies, nil
}
