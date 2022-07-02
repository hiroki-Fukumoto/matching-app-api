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
	SendLikeRequest() *SendLikeRequest
	CancelLikeRequest() *CancelLikeRequest
	SendLike(req *SendLikeRequest) error
	CancelLike(req *CancelLikeRequest) error
}

func NewSendLikeRepository(db *gorm.DB) SendLikeRepository {
	return &sendLikeRepository{DB: db}
}

type SendLikeRequest struct {
	SenderUserId   string
	ReceiverUserId string
}

func (sr *sendLikeRepository) SendLikeRequest() *SendLikeRequest {
	return &SendLikeRequest{}
}

func (sr *sendLikeRepository) SendLike(req *SendLikeRequest) error {
	db := sr.DB

	sendLike := &model.SendLike{
		SenderUserID:   req.SenderUserId,
		ReceiverUserID: req.ReceiverUserId,
	}

	err := db.Where("sender_user_id = ?", req.SenderUserId).Where("receiver_user_id = ?", req.ReceiverUserId).First(&sendLike).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Create(&sendLike).Error; err != nil {
			return err
		} else {
			return nil
		}
	}

	return error_handler.ErrBadRequest
}

type CancelLikeRequest struct {
	SenderUserId   string
	ReceiverUserId string
}

func (sr *sendLikeRepository) CancelLikeRequest() *CancelLikeRequest {
	return &CancelLikeRequest{}
}

func (sr *sendLikeRepository) CancelLike(req *CancelLikeRequest) error {
	db := sr.DB

	sendLike := *&model.SendLike{}

	if err := db.Where("sender_user_id = ?", req.SenderUserId).Where("receiver_user_id = ?", req.ReceiverUserId).Delete(&sendLike).Error; err != nil {
		return err
	}

	return nil
}
