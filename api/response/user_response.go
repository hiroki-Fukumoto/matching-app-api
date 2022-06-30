package response

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
)

type MeResponse struct {
	ID    string `json:"id" validate:"required"`    // ID
	Name  string `json:"name" validate:"required"`  // 名前
	Email string `json:"email" validate:"required"` // メールアドレス
}

func (m *MeResponse) ToMeResponse(u *model.User) MeResponse {
	m.ID = u.Base.ID
	m.Name = u.Name
	m.Email = u.Email

	return *m
}

type AuthenticationResponse struct {
	ApiToken string `json:"api_token" validate:"required"` // IDトークン
}

type LoginUserResponse struct {
	Me             MeResponse             `json:"me" validate:"required"`
	Authentication AuthenticationResponse `json:"authentication" validate:"required"`
}

func (l *LoginUserResponse) ToLoginUserResponse(u *model.User, apiToken string) LoginUserResponse {
	l.Me = l.Me.ToMeResponse(u)
	l.Authentication.ApiToken = apiToken

	return *l
}
