package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Base
	Name       string    `gorm:"size:64;not null"`
	Email      string    `gorm:"size:128;not null;unique"`
	Sex        string    `gorm:"type: enum('male', 'female', 'other');default:'other';not null"` // TODO: type から読み込む方法
	Birthday   time.Time `gorm:"not null;type:date"`
	Prefecture uint16    `gorm:"not null;comment:都道府県コード(JIS X 0401 の規格に基づく) ex)1: 北海道"`
	Message    *string   `gorm:"size:255;"`
	Like       uint16    `gorm:"not null;default:0;comment:いいね数"`
	Avatar     string    `gorm:"size:255;comment:アバター画像URL"`
	Password   []byte    `gorm:"not null'"`
	DeletedAt  gorm.DeletedAt

	Hobbies []Hobby `gorm:"many2many:user_hobbies;"`
}
