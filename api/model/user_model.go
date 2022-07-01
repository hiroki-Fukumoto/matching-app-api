package model

import "time"

type User struct {
	Base
	Name     string    `gorm:"size:255;not null"`
	Email    string    `gorm:"size:255;not null;unique"`
	Sex      string    `gorm:"type: enum('male', 'female', 'other'); default: 'other'; not null"` // TODO: type から読み込む方法
	Birthday time.Time `gorm:"not null;type:date"`
	Password []byte    `gorm:"not null'"`
}
