package model

import (
	"gorm.io/gorm"
)

type Message struct {
	Base
	SenderUserID   string `gorm:"size:36;index;not null"`
	ReceiverUserID string `gorm:"size:36;index;not null"`
	Message        string `gorm:"size:255;not null"`
	IsRead         bool   `gorm:"default:false;not null"`
	DeletedAt      gorm.DeletedAt

	Sender   User `gorm:"foreignkey:SenderUserID"`
	Receiver User `gorm:"foreignkey:ReceiverUserID"`
}
