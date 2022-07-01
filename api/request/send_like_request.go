package request

type SendLikeRequest struct {
	ReceiverUserID string `json:"receiver_user_id" validate:"required" ja:"受け取り手"` // いいねを受け取るユーザー
}
