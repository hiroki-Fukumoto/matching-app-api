package enum

type Message string

const (
	BadRequest          Message = "不正なリクエストです"
	Unauthorized        Message = "認証が必要です"
	Forbidden           Message = "実行権限がありません"
	NotFound            Message = "指定された情報は存在しません"
	InternalServerError Message = "サーバーエラーが発生しました"
	IDNoFound           Message = "IDが指定されていません"
	TokenFormatError    Message = "トークンが不正です"
	ExpiredToken        Message = "トークンの有効期限が切れています"
)

func (v Message) String() string {
	return string(v)
}
