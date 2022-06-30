package model

import "github.com/google/uuid"

type Device struct {
	Base
	UserID            uuid.UUID `gorm:"type:char(36);index"`
	DeviceToken       string    `gorm:"size:64;not null;unique"`
	Os                string    `gorm:"size:64;not null"`
	CurrentAppVersion string    `gorm:"size:16;not null"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
