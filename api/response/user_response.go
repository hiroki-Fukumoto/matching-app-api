package response

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

type Prefecture struct {
	Code int
	Name string
}

type MeResponse struct {
	ID         string     `json:"id" validate:"required"`         // ID
	Name       string     `json:"name" validate:"required"`       // 名前
	Email      string     `json:"email" validate:"required"`      // メールアドレス
	Sex        string     `json:"sex" validate:"required"`        // 性別
	Birthday   string     `json:"birthday" validate:"required"`   // 生年月日
	Message    *string    `json:"message"`                        // メッセージ
	Avatar     *string    `json:"avatar"`                         // アバター
	Like       int        `json:"like" validate:"required"`       // いいね数
	Prefecture Prefecture `json:"prefecture" validate:"required"` // 都道府県
}

func (m *MeResponse) ToMeResponse(u *model.User) MeResponse {
	m.ID = u.Base.ID
	m.Name = u.Name
	m.Email = u.Email
	m.Sex = u.Sex
	m.Birthday = util.FormatDate(u.Birthday)
	m.Message = u.Message
	m.Avatar = u.Avatar
	m.Like = int(u.Like)

	code := int(u.Prefecture)
	name, _ := util.GetPrefectureName(code)
	m.Prefecture = Prefecture{Code: code, Name: *name}

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
