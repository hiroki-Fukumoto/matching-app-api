package request

type CreateUserRequest struct {
	Name       string `json:"name" validate:"required" ja:"名前"`
	Email      string `json:"email" validate:"required,email,email_exists" ja:"メールアドレス"`
	Sex        string `json:"sex" validate:"required,sex" enums:"male,female,other" ja:"性別"`
	Birthday   string `json:"birthday" validate:"required,date_format" ja:"生年月日"`
	Prefecture int    `json:"prefecture" validate:"required,prefecture" ja:"都道府県"`
	Password   string `json:"password" validate:"required,gte=8,lt=64" ja:"パスワード"`
}

type SearchUserRequest struct {
	Page       int  `json:"page,omitempty"`
	Prefecture *int `json:"prefecture,omitempty"`
	FromAge    *int `json:"from_age,omitempty"`
	ToAge      *int `json:"to_age,omitempty"`
	Sort       *int `json:"sort,omitempty"` // TODO: 並び順の種類
}
