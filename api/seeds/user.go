package seeds

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hiroki-Fukumoto/matching-app-api/api/enum"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateDummyUser(db *gorm.DB) error {
	users := []model.User{}

	bmin := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	bmax := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	for i := 0; i < 100; i++ {
		u := model.User{}
		u.Name = fmt.Sprintf("dummy user%s", strconv.Itoa(i))

		if err := db.Where("name = ?", u.Name).First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			u.Email = fmt.Sprintf("dummy%s@test.com", strconv.Itoa(i))
			if i%2 == 0 {
				u.Sex = enum.SEX.MALE
			} else {
				u.Sex = enum.SEX.FEMALE
			}
			delta := bmax - bmin
			sec := rand.Int63n(delta) + bmin
			tm := time.Unix(sec, 0)
			u.Birthday = tm
			rand.Seed(time.Now().UnixNano())
			pmin := 1
			pmax := 47
			u.Prefecture = uint16(rand.Intn(pmax-pmin+1) + pmin)
			lmin := 1
			lmax := 500
			u.Like = uint16(rand.Intn(lmax-lmin+1) + lmin)
			var ava string = "https://placehold.jp/300x300.png"
			u.Avatar = &ava
			var m string = "初めまして。よろしくお願いします。"
			u.Message = &m
			passwordHash, _ := bcrypt.GenerateFromPassword([]byte("12345678"), 14)
			u.Password = passwordHash
			users = append(users, u)
		}
	}

	db.Create(&users)

	return nil
}
