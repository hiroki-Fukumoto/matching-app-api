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
	FindMessagesRequest() *FindMessagesRequest
	FindReceiveMessages(req *FindMessagesRequest) (messages []*model.Message, err error)
	SendMessageRequest() *SendMessageRequest
	SendMessage(req *SendMessageRequest) error
	ReadMessageRequest() *ReadMessageRequest
	ReadMessage(req *ReadMessageRequest) error
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{DB: db}
}

type FindMessagesRequest struct {
	ReceiverID string
}

func (mr *messageRepository) FindMessagesRequest() *FindMessagesRequest {
	return &FindMessagesRequest{}
}

func (mr *messageRepository) FindReceiveMessages(req *FindMessagesRequest) (messages []*model.Message, err error) {
	db := mr.DB

	if err := db.Where("receiver_id = ?", req.ReceiverID).
		Preload("Sender").
		Group("sender_id").
		Order("created_at desc").
		Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

type SendMessageRequest struct {
	SenderID   string
	ReceiverID string
	Message    string
}

func (mr *messageRepository) SendMessageRequest() *SendMessageRequest {
	return &SendMessageRequest{}
}

func (mr *messageRepository) SendMessage(req *SendMessageRequest) error {
	db := mr.DB

	message := &model.Message{
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
		Message:    html.EscapeString(req.Message),
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
