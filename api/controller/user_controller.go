package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/service"
	"github.com/hiroki-Fukumoto/matching-app-api/api/validator"
)

type UserController interface {
	Create(c *gin.Context)
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
