package repository

import (
	"errors"

	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"gorm.io/gorm"
)

type sendLikeRepository struct {
	DB *gorm.DB
}

type SendLikeRepository interface {
	SendLike(senderUserId string, receiverUserId string) error
	CancelLike(senderUserId string, receiverUserId string) error
}

func NewSendLikeRepository(db *gorm.DB) SendLikeRepository {
	return &sendLikeRepository{DB: db}
}

func (sr *sendLikeRepository) SendLike(senderUserId string, receiverUserId string) error {
	db := sr.DB

	sendLike := &model.SendLike{
		SenderUserID:   senderUserId,
		ReceiverUserID: receiverUserId,
	}

	err := db.Where("sender_user_id = ?", senderUserId).Where("receiver_user_id = ?", receiverUserId).First(&sendLike).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Create(&sendLike).Error; err != nil {
			return err
		}
	}

	return error_handler.ErrBadRequest
}

func (sr *sendLikeRepository) CancelLike(senderUserId string, receiverUserId string) error {
	db := sr.DB

	sendLike := *&model.SendLike{}

	if err := db.Where("sender_user_id = ?", senderUserId).Where("receiver_user_id = ?", receiverUserId).Delete(&sendLike).Error; err != nil {
		return err
	}

	return nil
}
