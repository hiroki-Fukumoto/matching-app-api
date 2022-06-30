package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
	"github.com/hiroki-Fukumoto/matching-app-api/api/service"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
	"github.com/hiroki-Fukumoto/matching-app-api/api/validator"
)

type UserController interface {
	Create(c *gin.Context)
	Me(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

// @Summary ユーザー新規作成
// @Description 新しいユーザーを作成する
// @Tags users
// @Accept json
// @Produce json
// @Param request body request.CreateUserRequest true "ユーザー情報"
// @Success 201 {object} response.LoginUserResponse{}
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/users [post]
func (uc userController) Create(c *gin.Context) {
	var req request.CreateUserRequest
	c.BindJSON(&req)
	if err := validator.Validate(&req); err != nil {
		errors := validator.GetErrorMessages(err)

		apiError := error_handler.ApiErrorHandle(error_handler.ErrBadRequest, error_handler.ErrorMessage(errors))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := uc.userService.Create(&req)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err)
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// @Summary ログインユーザー情報取得
// @Description ログイン中のユーザー情報を取得する
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Success 200 {object} response.MeResponse{}
// @Failure 401 {object} error_handler.ErrorResponse
// @Failure 403 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/users/me [get]
func (uc userController) Me(c *gin.Context) {
	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err)
		c.JSON(apiError.Status, apiError)
		return
	}

	res := &response.MeResponse{}
	res.ToMeResponse(user)

	c.JSON(http.StatusOK, res)
}
