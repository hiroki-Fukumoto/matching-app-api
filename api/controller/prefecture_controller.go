package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/service"
)

type PrefectureController interface {
	FindAll(c *gin.Context)
}

type prefectureController struct {
	prefectureService service.PrefectureService
}

func NewPrefectureController(ps service.PrefectureService) PrefectureController {
	return &prefectureController{
		prefectureService: ps,
	}
}

// @Summary 都道府県リスト取得
// @Description 都道府県のコードと名前のリストを取得する
// @Tags prefectures
// @Accept json
// @Produce json
// @Success 200 {object} []response.PrefectureResponse{}
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/prefectures [get]
func (pc prefectureController) FindAll(c *gin.Context) {
	res := pc.prefectureService.FindPrefectures()

	c.JSON(http.StatusOK, res)
}
