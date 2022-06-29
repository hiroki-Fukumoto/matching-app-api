package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController interface {
	HealthCheck(c *gin.Context)
}

type healthCheckController struct{}

func NewHealthCheckController() HealthCheckController {
	return &healthCheckController{}
}

// @Summary health check
// @Accept json
// @Produce json
// @Success 200 {string} Health Check OK
// @Router /api/v1/health-check [get]
func (hc healthCheckController) HealthCheck(c *gin.Context) {

	c.JSON(http.StatusOK, "Health Check OK")
}
