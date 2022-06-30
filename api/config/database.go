package config

import (
	"os"

	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() *gorm.DB {
	godotenv.Load(".env")
	godotenv.Load()

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	autoMigration()

	return db
}

func autoMigration() {
	db.AutoMigrate(
		&model.User{},
	)
}
