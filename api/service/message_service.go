package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
)

type MessageService interface {
	SendMessage(senderUserID string, receiverUserID string, message string) error
	ReadMessage(ID string) error
}

type messageService struct {
	messageRepository repository.MessageRepository
}

func NewMessageService(sr repository.MessageRepository) MessageService {
	return &messageService{
		messageRepository: sr,
	}
}

func (ss messageService) SendMessage(senderUserID string, receiverUserID string, message string) error {
	req := ss.messageRepository.SendMessageRequest()
	req.SenderUserID = senderUserID
	req.ReceiverUserID = receiverUserID
	req.Message = message
	err := ss.messageRepository.SendMessage(req)
	if err != nil {
		return err
	}
	return nil
}

func (ss messageService) ReadMessage(ID string) error {
	req := ss.messageRepository.ReadMessageRequest()
	req.ID = ID
	err := ss.messageRepository.ReadMessage(req)
	if err != nil {
		return err
	}
	return nil
}