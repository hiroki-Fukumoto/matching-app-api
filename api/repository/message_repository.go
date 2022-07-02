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
	ReadMessageRequest() *ReadMessageRequest
	ReadMessage(req *ReadMessageRequest) error
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

type ReadMessageRequest struct {
	ID string
}

func (mr *messageRepository) ReadMessageRequest() *ReadMessageRequest {
	return &ReadMessageRequest{}
}

func (mr *messageRepository) ReadMessage(req *ReadMessageRequest) error {
	db := mr.DB

	message := &model.Message{}

	if err := db.Model(&message).Where("id = ?", req.ID).Update("is_read", true).Error; err != nil {
		return err
	}

	return nil
}
