package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
)

type SendLikeService interface {
	SendLike(senderID string, receiverID string) error
	CancelLike(senderID string, receiverID string) error
	FindSendLikes(senderID string) (res []*response.SendLikeResponse, err error)
	FindReceiveLikes(receiverID string) (res []*response.ReceiveLikeResponse, err error)
}

type sendLikeService struct {
	sendLikeRepository repository.SendLikeRepository
}

func NewSendLikeService(sr repository.SendLikeRepository) SendLikeService {
	return &sendLikeService{
		sendLikeRepository: sr,
	}
}

func (ss sendLikeService) SendLike(senderID string, receiverID string) error {
	req := ss.sendLikeRepository.SendLikeRequest()
	req.SenderID = senderID
	req.ReceiverID = receiverID
	err := ss.sendLikeRepository.SendLike(req)
	if err != nil {
		return err
	}
	return nil
}

func (ss sendLikeService) CancelLike(senderID string, receiverID string) error {
	req := ss.sendLikeRepository.CancelLikeRequest()
	req.SenderID = senderID
	req.ReceiverID = receiverID
	err := ss.sendLikeRepository.CancelLike(req)
	if err != nil {
		return err
	}
	return nil
}

func (ss sendLikeService) FindSendLikes(senderID string) (res []*response.SendLikeResponse, err error) {
	req := ss.sendLikeRepository.FindSendLikesRequest()
	req.SenderID = senderID
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

func (ss sendLikeService) FindReceiveLikes(receiverID string) (res []*response.ReceiveLikeResponse, err error) {
	req := ss.sendLikeRepository.FindReceiveLikesRequest()
	req.ReceiverID = receiverID
	likes, err := ss.sendLikeRepository.FindReceiveLikes(req)
	if err != nil {
		return nil, err
	}
	res = []*response.ReceiveLikeResponse{}
	for _, l := range likes {
		r := &response.ReceiveLikeResponse{}
		r.ToReceiveLikeResponse(l)
		res = append(res, r)
	}

	return res, nil
}
