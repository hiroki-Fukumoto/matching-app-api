package seeds

import (
	"errors"

	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"gorm.io/gorm"
)

func CreateHobby(db *gorm.DB) error {
	hobbies := []model.Hobby{}

	hs := [...]string{"マンガ", "お酒", "スポーツ", "読書", "音楽", "アニメ", "アウトドア", "映画"}

	for _, h := range hs {
		if err := db.Where("name = ?", h).First(&model.Hobby{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			m := model.Hobby{Name: h}
			hobbies = append(hobbies, m)
		}
	}

	db.Create(&hobbies)

	return nil
}
