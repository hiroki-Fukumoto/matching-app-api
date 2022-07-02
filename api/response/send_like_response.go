package response

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

type SendLikeResponse struct {
	SentAt   string       `json:"sent_at" validate:"required"`  // 送信日時
	Receiver UserResponse `json:"receiver" validate:"required"` // 受け取り手
}

func (s *SendLikeResponse) ToSendLikeResponse(l *model.SendLike) SendLikeResponse {
	s.SentAt = util.FormatDateTime(*l.CreatedAt)
	s.Receiver = s.Receiver.ToUserResponse(&l.Receiver)

	return *s
}

type ReceiveLikeResponse struct {
	ReceivedAt string       `json:"received_at" validate:"required"` // 受信日時
	Sender     UserResponse `json:"sender" validate:"required"`      // 送り手
}

func (s *ReceiveLikeResponse) ToReceiveLikeResponse(l *model.SendLike) ReceiveLikeResponse {
	s.ReceivedAt = util.FormatDateTime(*l.CreatedAt)
	s.Sender = s.Sender.ToUserResponse(&l.Sender)

	return *s
}
