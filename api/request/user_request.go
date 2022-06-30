package request

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required" ja:"名前"`
	Email    string `json:"email" validate:"required,email,email_exists" ja:"メールアドレス"`
	Password string `json:"password" validate:"required,gte=8,lt=64" ja:"パスワード"`
}
