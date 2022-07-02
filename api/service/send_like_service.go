package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
)

type SendLikeService interface {
	SendLike(senderUserId string, receiverUserId string) error
	CancelLike(senderUserId string, receiverUserId string) error
}

type sendLikeService struct {
	sendLikeRepository repository.SendLikeRepository
}

func NewSendLikeService(sr repository.SendLikeRepository) SendLikeService {
	return &sendLikeService{
		sendLikeRepository: sr,
	}
}

func (ss sendLikeService) SendLike(senderUserId string, receiverUserId string) error {
	req := ss.sendLikeRepository.SendLikeRequest()
	req.SenderUserId = senderUserId
	req.ReceiverUserId = receiverUserId
	err := ss.sendLikeRepository.SendLike(req)
	if err != nil {
		return err
	}
	return nil
}

func (ss sendLikeService) CancelLike(senderUserId string, receiverUserId string) error {
	req := ss.sendLikeRepository.CancelLikeRequest()
	req.SenderUserId = senderUserId
	req.ReceiverUserId = receiverUserId
	err := ss.sendLikeRepository.CancelLike(req)
	if err != nil {
		return err
	}
	return nil
}
