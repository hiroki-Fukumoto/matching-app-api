package response

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

type ReceiveMessageResponse struct {
	ReceiveAt string       `json:"receive_at" validate:"required"` // 受信日時
	Message   string       `json:"message" validate:"required"`    // メッセージ
	Sender    UserResponse `json:"sender" validate:"required"`     // 送り手
}

func (mr *ReceiveMessageResponse) ToMessageResponse(m *model.Message) ReceiveMessageResponse {
	mr.ReceiveAt = util.FormatDateTime(*m.CreatedAt)
	mr.Message = m.Message
	mr.Sender = mr.Sender.ToUserResponse(&m.Sender)

	return *mr
}
