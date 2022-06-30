package model

type User struct {
	Base
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null;unique"`
	Password []byte `gorm:"not null'"`
}
