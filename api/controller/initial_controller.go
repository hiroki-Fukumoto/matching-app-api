package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app/api/controller/service"
)

type InitialController interface {
	Initial(c *gin.Context)
}

type initialController struct {
	initialService service.InitialService
}

func NewInitialController(is service.InitialService) InitialController {
	return &initialController{initialService: is}
}

// @Summary アプリ起動時にコールする
// @Tags initial
// @Accept json
// @Produce json
// @Success 200 {object} response.InitialResponse{}
// @Router /api/v1/initial [get]
func (ic initialController) Initial(c *gin.Context) {
	// TODO
	res, _ := ic.initialService.Initial()
	c.JSON(http.StatusOK, res)
}
