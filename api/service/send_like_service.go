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
	err := ss.sendLikeRepository.SendLike(senderUserId, receiverUserId)
	if err != nil {
		return err
	}
	return nil
}

func (ss sendLikeService) CancelLike(senderUserId string, receiverUserId string) error {
	err := ss.sendLikeRepository.CancelLike(senderUserId, receiverUserId)
	if err != nil {
		return err
	}
	return nil
}
