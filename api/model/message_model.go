package model

import (
	"gorm.io/gorm"
)

type Message struct {
	Base
	SenderID   string `gorm:"size:36;index;not null"`
	ReceiverID string `gorm:"size:36;index;not null"`
	Message    string `gorm:"size:255;not null"`
	IsRead     bool   `gorm:"default:false;not null"`
	DeletedAt  gorm.DeletedAt

	Sender   User `gorm:"foreignkey:SenderID"`
	Receiver User `gorm:"foreignkey:ReceiverID"`
}
