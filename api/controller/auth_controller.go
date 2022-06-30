package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/service"
	"github.com/hiroki-Fukumoto/matching-app-api/api/validator"
)

type AuthController interface {
	Login(c *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(au service.AuthService) AuthController {
	return &authController{
		authService: au,
	}
}

// @Summary ログイン
// @Description ログイン処理を行う。JWTを新たに発行する
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "ログイン情報"
// @Success 200 {object} response.LoginUserResponse{}
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 401 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/login [post]
func (ac authController) Login(c *gin.Context) {
	var request request.LoginRequest
	c.BindJSON(&request)
	if err := validator.Validate(&request); err != nil {
		errors := validator.GetErrorMessages(err)

		apiError := error_handler.ApiErrorHandle(error_handler.ErrBadRequest, error_handler.ErrorMessage(errors))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := ac.authService.Login(&request)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(error_handler.ErrUnauthorized)
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}
