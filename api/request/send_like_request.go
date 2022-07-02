package request

type SendLikeRequest struct {
	ReceiverID string `json:"receiver_id" validate:"required" ja:"受け取り手"` // いいねを受け取るユーザー
}
