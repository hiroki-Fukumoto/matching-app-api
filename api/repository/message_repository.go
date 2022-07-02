package repository

import (
	"html"

	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"gorm.io/gorm"
)

type messageRepository struct {
	DB *gorm.DB
}

type MessageRepository interface {
	SendMessageRequest() *SendMessageRequest
	SendMessage(req *SendMessageRequest) error
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{DB: db}
}

type SendMessageRequest struct {
	SenderUserID   string
	ReceiverUserID string
	Message        string
}

func (mr *messageRepository) SendMessageRequest() *SendMessageRequest {
	return &SendMessageRequest{}
}

func (mr *messageRepository) SendMessage(req *SendMessageRequest) error {
	db := mr.DB

	message := &model.Message{
		SenderUserID:   req.SenderUserID,
		ReceiverUserID: req.ReceiverUserID,
		Message:        html.EscapeString(req.Message),
	}

	if err := db.Create(&message).Error; err != nil {
		return err
	}

	return nil
}
