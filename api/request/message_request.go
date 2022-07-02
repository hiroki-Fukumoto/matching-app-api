package request

type SendMessageRequest struct {
	ReceiverID string `json:"receiver_id" validate:"required" ja:"受け取り手"` // メッセージを受け取るユーザー
	Message    string `json:"message" validate:"required" ja:"メッセージ内容"`   // メッセージ
}
