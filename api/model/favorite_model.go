package model

type Favorite struct {
	Base
	SenderID   string `gorm:"size:36;index;not null"`
	ReceiverID string `gorm:"size:36;index;not null"`

	Sender   User `gorm:"foreignkey:SenderID"`
	Receiver User `gorm:"foreignkey:ReceiverID"`
}
