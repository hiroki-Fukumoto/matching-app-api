package model

type SendLike struct {
	Base
	SenderUserID   string `gorm:"size:36;index;not null"`
	ReceiverUserID string `gorm:"size:36;index;not null"`

	Sender   User `gorm:"foreignkey:SenderUserID"`
	Receiver User `gorm:"foreignkey:ReceiverUserID"`
}
