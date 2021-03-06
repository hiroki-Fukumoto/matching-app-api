package model

import "gorm.io/gorm"

type Device struct {
	Base
	UserID            string `gorm:"type:size:36;index"`
	DeviceToken       string `gorm:"size:64;not null;unique"`
	Os                string `gorm:"size:64;not null"`
	CurrentAppVersion string `gorm:"size:16;not null"`
	DeletedAt         gorm.DeletedAt

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
