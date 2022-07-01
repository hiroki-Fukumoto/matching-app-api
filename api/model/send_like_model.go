package model

type SendLike struct {
	Base
	SenderUserID   string `gorm:"size:36;index"`
	ReceiverUserID string `gorm:"size:36;index"`

	Sender   User `gorm:"foreignkey:SenderUserID"`
	Receiver User `gorm:"foreignkey:ReceiverUserID"`
}
