package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
)

type SendLikeService interface {
	SendLike(senderUserID string, receiverUserID string) error
	CancelLike(senderUserID string, receiverUserID string) error
	FindSendLikes(senderUserID string) (res []*response.SendLikeResponse, err error)
	FindReceiveLikes(receiverUserID string) (res []*response.ReceiveLikeResponse, err error)
}

type sendLikeService struct {
	sendLikeRepository repository.SendLikeRepository
}

func NewSendLikeService(sr repository.SendLikeRepository) SendLikeService {
	return &sendLikeService{
		sendLikeRepository: sr,
	}
}

func (ss sendLikeService) SendLike(senderUserID string, receiverUserID string) error {
	req := ss.sendLikeRepository.SendLikeRequest()
	req.SenderUserID = senderUserID
	req.ReceiverUserID = receiverUserID
	err := ss.sendLikeRepository.SendLike(req)
	if err != nil {
		return err
	}
	return nil
}

func (ss sendLikeService) CancelLike(senderUserID string, receiverUserID string) error {
	req := ss.sendLikeRepository.CancelLikeRequest()
	req.SenderUserID = senderUserID
	req.ReceiverUserID = receiverUserID
	err := ss.sendLikeRepository.CancelLike(req)
	if err != nil {
		return err
	}
	return nil
}

func (ss sendLikeService) FindSendLikes(senderUserID string) (res []*response.SendLikeResponse, err error) {
	req := ss.sendLikeRepository.FindSendLikesRequest()
	req.SenderUserID = senderUserID
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

func (ss sendLikeService) FindReceiveLikes(receiverUserID string) (res []*response.ReceiveLikeResponse, err error) {
	req := ss.sendLikeRepository.FindReceiveLikesRequest()
	req.ReceiverUserID = receiverUserID
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
