package response

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

type SendLikeResponse struct {
	SentAt   string       `json:"sen_at" validate:"required"`   // 送信日時
	Receiver UserResponse `json:"receiver" validate:"required"` // 受け取り手
}

func (s *SendLikeResponse) ToSendLikeResponse(l *model.SendLike) SendLikeResponse {
	s.SentAt = util.FormatDateTime(*l.CreatedAt)
	s.Receiver = s.Receiver.ToUserResponse(&l.Receiver)

	return *s
}
