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

type SendLikeController interface {
	SendLike(c *gin.Context)
	CancelLike(c *gin.Context)
	FindSendLikes(c *gin.Context)
}

type sendLikeController struct {
	sendLikeService service.SendLikeService
}

func NewSendLikeController(us service.SendLikeService) SendLikeController {
	return &sendLikeController{
		sendLikeService: us,
	}
}

// @Summary いいねを送る
// @Description いいねを送る
// @Tags send like
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param request body request.SendLikeRequest true "いいねを送る情報"
// @Success 204 nil nil
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/likes [post]
func (sc sendLikeController) SendLike(c *gin.Context) {
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

	if user.Base.ID == req.ReceiverUserID {
		apiError := error_handler.ApiErrorHandle("", error_handler.ErrBadRequest, error_handler.ErrorMessage([]string{"いいねを送るユーザーと受け取るユーザーが同一です"}))
		c.JSON(apiError.Status, apiError)
		return
	}

	err = sc.sendLikeService.SendLike(user.Base.ID, req.ReceiverUserID)
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
// @Tags send like
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param request body request.SendLikeRequest true "いいねを取り消す情報"
// @Success 204 nil nil
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/likes/cancel [delete]
func (sc sendLikeController) CancelLike(c *gin.Context) {
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

	err = sc.sendLikeService.CancelLike(user.Base.ID, req.ReceiverUserID)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary 送信したいいね一覧を取得する
// @Description 登録日が新しいもの順で返す
// @Tags send like
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Success 200 {object} []response.SendLikeResponse
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/likes [get]
func (sc sendLikeController) FindSendLikes(c *gin.Context) {
	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := sc.sendLikeService.FindSendLikes(user.Base.ID)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}
