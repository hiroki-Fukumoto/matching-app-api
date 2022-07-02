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
	FindSendLikesRequest() *FindSendLikesRequest
	FindSendLikes(req *FindSendLikesRequest) (likes []*model.SendLike, err error)
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

type FindSendLikesRequest struct {
	SenderUserId string
}

func (sr *sendLikeRepository) FindSendLikesRequest() *FindSendLikesRequest {
	return &FindSendLikesRequest{}
}

func (sr *sendLikeRepository) FindSendLikes(req *FindSendLikesRequest) (likes []*model.SendLike, err error) {
	db := sr.DB
	if err := db.Where("sender_user_id = ?", req.SenderUserId).
		Preload("Receiver").
		Order("created_at desc").
		Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}
