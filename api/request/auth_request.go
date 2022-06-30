package request

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" ja:"メールアドレス"`
	Password string `json:"password" validate:"required" ja:"パスワード"`
}
