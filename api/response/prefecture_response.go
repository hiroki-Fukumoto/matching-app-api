package response

type PrefectureResponse struct {
	Code int    `json:"code" validate:"required"` // 都道府県コード
	Name string `json:"name" validate:"required"` // 都道府県名
}
