package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/enum"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/service"
)

type HobbyController interface {
	FindAll(c *gin.Context)
}

type hobbyController struct {
	hobbyService service.HobbyService
}

func NewHobbyController(hs service.HobbyService) HobbyController {
	return &hobbyController{
		hobbyService: hs,
	}
}

// @Summary 趣味マスター一覧を取得する
// @Description 趣味マスター
// @Tags hobby
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Success 200 {object} []response.HobbyResponse
// @Failure 403 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/hobbies [get]
func (hc hobbyController) FindAll(c *gin.Context) {

	res, err := hc.hobbyService.FindAll()
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}
