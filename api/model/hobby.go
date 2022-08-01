package model

type Hobby struct {
	Base
	Name string `gorm:"size:64;not null;unique"`
}
