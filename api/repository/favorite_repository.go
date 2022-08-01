package repository

import (
	"errors"
	"fmt"

	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"gorm.io/gorm"
)

type favoriteRepository struct {
	DB *gorm.DB
}

type FavoriteRepository interface {
	SendLikeRequest() *SendLikeRequest
	CancelLikeRequest() *CancelLikeRequest
	SendLike(req *SendLikeRequest) error
	CancelLike(req *CancelLikeRequest) error
	FindSendLikesRequest() *FindSendLikesRequest
	FindSendLikes(req *FindSendLikesRequest) (likes []*model.Favorite, err error)
	FindReceiveLikesRequest() *FindReceiveLikesRequest
	FindReceiveLikes(req *FindReceiveLikesRequest) (likes []*model.Favorite, err error)
}

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository {
	return &favoriteRepository{DB: db}
}

type SendLikeRequest struct {
	SenderID   string
	ReceiverID string
}

func (sr *favoriteRepository) SendLikeRequest() *SendLikeRequest {
	return &SendLikeRequest{}
}

func (sr *favoriteRepository) SendLike(req *SendLikeRequest) error {
	db := sr.DB

	sendLike := &model.Favorite{
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
	}

	err := db.Where("sender_id = ?", req.SenderID).Where("receiver_id = ?", req.ReceiverID).First(&sendLike).Error

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
	SenderID   string
	ReceiverID string
}

func (sr *favoriteRepository) CancelLikeRequest() *CancelLikeRequest {
	return &CancelLikeRequest{}
}

func (sr *favoriteRepository) CancelLike(req *CancelLikeRequest) error {
	db := sr.DB

	sendLike := *&model.Favorite{}

	if err := db.Where("sender_id = ?", req.SenderID).Where("receiver_id = ?", req.ReceiverID).Delete(&sendLike).Error; err != nil {
		return err
	}

	return nil
}

type FindSendLikesRequest struct {
	SenderID string
}

func (sr *favoriteRepository) FindSendLikesRequest() *FindSendLikesRequest {
	return &FindSendLikesRequest{}
}

func (sr *favoriteRepository) FindSendLikes(req *FindSendLikesRequest) (likes []*model.Favorite, err error) {
	db := sr.DB
	if err := db.Where("sender_id = ?", req.SenderID).
		Preload("Receiver").
		Order("created_at desc").
		Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

type FindReceiveLikesRequest struct {
	ReceiverID string
}

func (sr *favoriteRepository) FindReceiveLikesRequest() *FindReceiveLikesRequest {
	return &FindReceiveLikesRequest{}
}

func (sr *favoriteRepository) FindReceiveLikes(req *FindReceiveLikesRequest) (likes []*model.Favorite, err error) {
	db := sr.DB
	if err := db.Where("receiver_id = ?", req.ReceiverID).
		Preload("Sender").
		Order("created_at desc").
		Find(&likes).Error; err != nil {
		return nil, err
	}
	fmt.Println(likes)
	return likes, nil
}
