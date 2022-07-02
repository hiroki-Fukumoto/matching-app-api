package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
)

type MessageService interface {
	FindReceiveMessages(receiverID string) (res []*response.ReceiveMessageResponse, err error)
	SendMessage(senderID string, receiverID string, message string) error
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

func (ms messageService) FindReceiveMessages(receiverID string) (res []*response.ReceiveMessageResponse, err error) {
	req := ms.messageRepository.FindMessagesRequest()
	req.ReceiverID = receiverID
	messages, err := ms.messageRepository.FindReceiveMessages(req)
	if err != nil {
		return nil, err
	}

	res = []*response.ReceiveMessageResponse{}
	for _, m := range messages {
		r := &response.ReceiveMessageResponse{}
		r.ToMessageResponse(m)
		res = append(res, r)
	}
	return res, nil
}

func (ms messageService) SendMessage(senderID string, receiverID string, message string) error {
	req := ms.messageRepository.SendMessageRequest()
	req.SenderID = senderID
	req.ReceiverID = receiverID
	req.Message = message
	err := ms.messageRepository.SendMessage(req)
	if err != nil {
		return err
	}
	return nil
}

func (ms messageService) ReadMessage(ID string) error {
	req := ms.messageRepository.ReadMessageRequest()
	req.ID = ID
	err := ms.messageRepository.ReadMessage(req)
	if err != nil {
		return err
	}
	return nil
}
