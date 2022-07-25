package response

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/enum"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

func (p *PrefectureResponse) toPrefectureResponse(u *model.User) PrefectureResponse {
	code := int(u.Prefecture)
	name, _ := util.GetPrefectureName(code)
	p.Code = int(u.Prefecture)
	p.Name = *name

	return *p
}

type MeResponse struct {
	ID         string             `json:"id" validate:"required"`                            // ID
	Name       string             `json:"name" validate:"required"`                          // 名前
	Email      string             `json:"email" validate:"required"`                         // メールアドレス
	Sex        enum.Sex           `json:"sex" validate:"required" enums:"male,female,other"` // 性別
	Birthday   string             `json:"birthday" validate:"required"`                      // 生年月日
	Message    *string            `json:"message"`                                           // メッセージ
	Avatar     string             `json:"avatar" validate:"required"`                        // アバター
	Like       int                `json:"like" validate:"required"`                          // いいね数
	Prefecture PrefectureResponse `json:"prefecture" validate:"required"`                    // 都道府県
}

func (m *MeResponse) ToMeResponse(u *model.User) MeResponse {
	m.ID = u.Base.ID
	m.Name = u.Name
	m.Email = u.Email
	m.Sex = enum.Sex(u.Sex)
	m.Birthday = util.FormatDate(u.Birthday)
	m.Message = u.Message
	m.Avatar = u.Avatar
	m.Like = int(u.Like)
	m.Prefecture = m.Prefecture.toPrefectureResponse(u)

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

type UserResponse struct {
	ID         string             `json:"id" validate:"required"`                            // ID
	Name       string             `json:"name" validate:"required"`                          // 名前
	Sex        enum.Sex           `json:"sex" validate:"required" enums:"male,female,other"` // 性別
	Birthday   string             `json:"birthday" validate:"required"`                      // 生年月日
	Message    *string            `json:"message"`                                           // メッセージ
	Avatar     string             `json:"avatar" validate:"required"`                        // アバター
	Like       int                `json:"like" validate:"required"`                          // いいね数
	Prefecture PrefectureResponse `json:"prefecture" validate:"required"`                    // 都道府県
}

func (ur *UserResponse) ToUserResponse(u *model.User) UserResponse {
	ur.ID = u.Base.ID
	ur.Name = u.Name
	ur.Sex = enum.Sex(u.Sex)
	ur.Birthday = util.FormatDate(u.Birthday)
	ur.Message = u.Message
	ur.Avatar = u.Avatar
	ur.Like = int(u.Like)
	ur.Prefecture = ur.Prefecture.toPrefectureResponse(u)

	return *ur
}
