package config

import (
	"log"
	"os"
	"time"

	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/seeds"
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
	err = dbConnect()
	if err != nil {
		log.Fatalln(err)
	}

	autoMigration()
	return db
}

func dbConnect() error {
	godotenv.Load(".env")
	godotenv.Load()

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	var count int = 10
	for count > 1 {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			time.Sleep(time.Second * 2)
			count--
			log.Printf("retry... count:%v\n", count)
			continue
		}
		break
	}

	return err
}

func autoMigration() {
	db.AutoMigrate(
		&model.User{},
		&model.Device{},
		&model.SendLike{},
		&model.Message{},
	)

	seeds.Seed(db)
}
