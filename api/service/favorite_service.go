package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
)

type FavoriteService interface {
	SendLike(senderID string, receiverID string) error
	CancelLike(senderID string, receiverID string) error
	FindSendLikes(senderID string) (res []*response.SendLikeResponse, err error)
	FindReceiveLikes(receiverID string) (res []*response.ReceiveLikeResponse, err error)
}

type favoriteService struct {
	favoriteRepository repository.FavoriteRepository
}

func NewFavoriteService(sr repository.FavoriteRepository) FavoriteService {
	return &favoriteService{
		favoriteRepository: sr,
	}
}

func (ss favoriteService) SendLike(senderID string, receiverID string) error {
	req := ss.favoriteRepository.SendLikeRequest()
	req.SenderID = senderID
	req.ReceiverID = receiverID
	err := ss.favoriteRepository.SendLike(req)
	if err != nil {
		return err
	}
	return nil
}

func (ss favoriteService) CancelLike(senderID string, receiverID string) error {
	req := ss.favoriteRepository.CancelLikeRequest()
	req.SenderID = senderID
	req.ReceiverID = receiverID
	err := ss.favoriteRepository.CancelLike(req)
	if err != nil {
		return err
	}
	return nil
}

func (ss favoriteService) FindSendLikes(senderID string) (res []*response.SendLikeResponse, err error) {
	req := ss.favoriteRepository.FindSendLikesRequest()
	req.SenderID = senderID
	likes, err := ss.favoriteRepository.FindSendLikes(req)
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

func (ss favoriteService) FindReceiveLikes(receiverID string) (res []*response.ReceiveLikeResponse, err error) {
	req := ss.favoriteRepository.FindReceiveLikesRequest()
	req.ReceiverID = receiverID
	likes, err := ss.favoriteRepository.FindReceiveLikes(req)
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
