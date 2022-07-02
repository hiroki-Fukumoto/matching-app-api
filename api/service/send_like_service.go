package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
)

type SendLikeService interface {
	SendLike(senderUserId string, receiverUserId string) error
	CancelLike(senderUserId string, receiverUserId string) error
	FindSendLikes(senderUserId string) (res []*response.SendLikeResponse, err error)
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

func (ss sendLikeService) FindSendLikes(senderUserId string) (res []*response.SendLikeResponse, err error) {
	req := ss.sendLikeRepository.FindSendLikesRequest()
	req.SenderUserId = senderUserId
	likes, err := ss.sendLikeRepository.FindSendLikes(req)
	if err != nil {
		return nil, err
	}
	res = []*response.SendLikeResponse{}
	for _, l := range likes {
		r := &response.SendLikeResponse{}
		r.ToSendLikeResponse(l)
		res = append(res, r)
	}

	return res, nil
}
