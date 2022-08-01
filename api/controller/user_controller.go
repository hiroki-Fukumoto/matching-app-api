package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/enum"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
	"github.com/hiroki-Fukumoto/matching-app-api/api/service"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
	"github.com/hiroki-Fukumoto/matching-app-api/api/validator"
)

type UserController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Me(c *gin.Context)
	PickupToday(c *gin.Context)
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
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
// @Tags user
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

		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrBadRequest, error_handler.ErrorMessage(errors))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := uc.userService.Create(&req)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// @Summary ユーザー情報更新
// @Description ユーザー情報を更新する
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param request body request.UpdateUserRequest true "更新内容"
// @Success 200 {object} response.MeResponse{}
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 403 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/users [patch]
func (uc userController) Update(c *gin.Context) {
	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	var req request.UpdateUserRequest
	c.BindJSON(&req)
	if err := validator.Validate(&req); err != nil {
		errors := validator.GetErrorMessages(err)

		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrBadRequest, error_handler.ErrorMessage(errors))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := uc.userService.Update(c, user.ID, &req)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary ログインユーザー情報取得
// @Description ログイン中のユーザー情報を取得する
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Success 200 {object} response.MeResponse{}
// @Failure 401 {object} error_handler.ErrorResponse
// @Failure 403 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/users/info/me [get]
func (uc userController) Me(c *gin.Context) {
	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	res := &response.MeResponse{}
	res.ToMeResponse(user)

	c.JSON(http.StatusOK, res)
}

// @Summary 本日のピックアップユーザー取得
// @Description ログインユーザーとは異なる性別のユーザーを20件返す
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Success 200 {object} []response.UserResponse{}
// @Failure 401 {object} error_handler.ErrorResponse
// @Failure 403 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/users/pickup/today [get]
func (uc userController) PickupToday(c *gin.Context) {
	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	var targetSex string
	if enum.Sex(user.Sex) == enum.SEX.MALE {
		targetSex = string(enum.SEX.FEMALE)
	} else {
		targetSex = string(enum.SEX.MALE)
	}

	res, err := uc.userService.PickupToday(targetSex)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary ユーザー情報全件取得
// @Description 50件ずつ取得。検索条件がない場合は登録日が新しい順に返す（今の所）
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param request body request.SearchUserRequest true "ユーザー情報"
// @Success 200 {object} []response.UserResponse{}
// @Failure 401 {object} error_handler.ErrorResponse
// @Failure 403 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/users/all [post]
func (uc userController) FindAll(c *gin.Context) {
	// req := request.SearchUserRequest{}
	var req request.SearchUserRequest
	c.BindJSON(&req)
	if err := validator.Validate(&req); err != nil {
		errors := validator.GetErrorMessages(err)

		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrBadRequest, error_handler.ErrorMessage(errors))
		c.JSON(apiError.Status, apiError)
		return
	}

	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := uc.userService.FindAll(&req, user.ID)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary ユーザー詳細情報取得
// @Description 指定したユーザーの詳細情報を取得する
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param id path string true "id"
// @Success 200 {object} response.UserResponse{}
// @Failure 401 {object} error_handler.ErrorResponse
// @Failure 403 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/users/{id} [get]
func (uc userController) FindByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		apiError := error_handler.ApiErrorHandle("", error_handler.ErrBadRequest, error_handler.ErrorMessage([]string{enum.IDNoFound.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := uc.userService.FindByID(id, user.ID)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}
