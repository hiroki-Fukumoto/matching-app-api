package seeds

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	if db.Migrator().HasTable(&model.User{}) {
		err := CreateDummyUser(db)
		if err != nil {
			panic(err)
		}
	}
}
