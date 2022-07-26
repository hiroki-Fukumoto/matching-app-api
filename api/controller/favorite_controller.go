package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/enum"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/service"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
	"github.com/hiroki-Fukumoto/matching-app-api/api/validator"
)

type FavoriteController interface {
	SendLike(c *gin.Context)
	CancelLike(c *gin.Context)
	FindSendLikes(c *gin.Context)
	FindReceiveLikes(c *gin.Context)
}

type favoriteController struct {
	favoriteService service.FavoriteService
}

func NewFavoriteController(us service.FavoriteService) FavoriteController {
	return &favoriteController{
		favoriteService: us,
	}
}

// @Summary いいねを送る
// @Description いいねを送る
// @Tags like
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param request body request.SendLikeRequest true "いいねを送る情報"
// @Success 204 {object} nil
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/likes [post]
func (sc favoriteController) SendLike(c *gin.Context) {
	var req request.SendLikeRequest
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

	if user.Base.ID == req.ReceiverID {
		apiError := error_handler.ApiErrorHandle("", error_handler.ErrBadRequest, error_handler.ErrorMessage([]string{"いいねを送るユーザーと受け取るユーザーが同一です"}))
		c.JSON(apiError.Status, apiError)
		return
	}

	err = sc.favoriteService.SendLike(user.Base.ID, req.ReceiverID)
	if err != nil {
		if err.Error() == error_handler.ErrBadRequest.Error() {
			apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrBadRequest, error_handler.ErrorMessage([]string{"既にいいねを送っているユーザーです"}))
			c.JSON(apiError.Status, apiError)
			return
		}

		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary いいねを取り消す
// @Description いいねを取り消す
// @Tags like
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param receiverID path string true "取り消しにするユーザーID"
// @Success 204 {object} nil
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/likes/{receiverID}/cancel [delete]
func (sc favoriteController) CancelLike(c *gin.Context) {
	receiverID := c.Param("receiverID")
	if receiverID == "" {
		apiError := error_handler.ApiErrorHandle("", error_handler.ErrBadRequest, error_handler.ErrorMessage([]string{"IDが指定されていません"}))
		c.JSON(apiError.Status, apiError)
	}

	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	err = sc.favoriteService.CancelLike(user.Base.ID, receiverID)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary 送信したいいね一覧を取得する
// @Description 登録日が新しいもの順で返す
// @Tags like
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Success 200 {object} []response.SendLikeResponse
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/likes/send [get]
func (sc favoriteController) FindSendLikes(c *gin.Context) {
	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := sc.favoriteService.FindSendLikes(user.Base.ID)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 受信したいいね一覧を取得する
// @Description 登録日が新しいもの順で返す
// @Tags like
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Success 200 {object} []response.ReceiveLikeResponse
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/likes/receive [get]
func (sc favoriteController) FindReceiveLikes(c *gin.Context) {
	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := sc.favoriteService.FindReceiveLikes(user.Base.ID)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}
