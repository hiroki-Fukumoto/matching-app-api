package response

import (
	"github.com/google/uuid"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
)

type MeResponse struct {
	ID    uuid.UUID `json:"id" validate:"required"`    // ID
	Name  string    `json:"name" validate:"required"`  // 名前
	Email string    `json:"email" validate:"required"` // メールアドレス
}

func ToMeResponse(u *model.User) (res MeResponse) {
	res.ID = u.Base.ID
	res.Name = u.Name
	res.Email = u.Email

	return res
}

type AuthenticationResponse struct {
	ApiToken string `json:"api_token" validate:"required"` // IDトークン
}

type LoginUserResponse struct {
	Me             MeResponse             `json:"me" validate:"required"`
	Authentication AuthenticationResponse `json:"authentication" validate:"required"`
}

func ToLoginUserResponse(u *model.User, apiToken string) (res LoginUserResponse) {
	res.Me = ToMeResponse(u)
	res.Authentication.ApiToken = apiToken

	return res
}
