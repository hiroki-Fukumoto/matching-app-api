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

type MessageController interface {
	FindReceiveMessages(c *gin.Context)
	SendMessage(c *gin.Context)
	ReadMessage(c *gin.Context)
}

type messageController struct {
	messageService service.MessageService
}

func NewMessageController(us service.MessageService) MessageController {
	return &messageController{
		messageService: us,
	}
}

// @Summary 受信済みメッセージを取得
// @Description 送信者別の受信メッセージ(最新の1通のみ)を受信日が最新のもの順に返す。
// @Tags messages
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Success 200 {object} []response.ReceiveMessageResponse
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/messages [get]
func (mc messageController) FindReceiveMessages(c *gin.Context) {
	user, err := util.GetLoginUser(c)

	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	res, err := mc.messageService.FindReceiveMessages(user.Base.ID)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary メッセージを送る
// @Description メッセージを送る
// @Tags messages
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param request body request.SendMessageRequest true "メッセージを送る情報"
// @Success 204 {object} nil
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/messages [post]
func (mc messageController) SendMessage(c *gin.Context) {
	var req request.SendMessageRequest
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
		apiError := error_handler.ApiErrorHandle("", error_handler.ErrBadRequest, error_handler.ErrorMessage([]string{"メッセージを送るユーザーと受け取るユーザーが同一です"}))
		c.JSON(apiError.Status, apiError)
		return
	}

	err = mc.messageService.SendMessage(user.Base.ID, req.ReceiverID, req.Message)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary メッセージを既読にする
// @Description メッセージを既読にする
// @Tags messages
// @Accept json
// @Produce json
// @Param Authorization header string true "ログイン時に取得したIDトークン(Bearer)"
// @Param id path string true "既読にするメッセージID"
// @Success 204 {object} nil
// @Failure 400 {object} error_handler.ErrorResponse
// @Failure 500 {object} error_handler.ErrorResponse
// @Router /api/v1/messages/{id}/read [put]
func (mc messageController) ReadMessage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		apiError := error_handler.ApiErrorHandle("", error_handler.ErrBadRequest, error_handler.ErrorMessage([]string{"IDが指定されていません"}))
		c.JSON(apiError.Status, apiError)
	}

	err := mc.messageService.ReadMessage(id)
	if err != nil {
		apiError := error_handler.ApiErrorHandle(err.Error(), error_handler.ErrInternalServerError, error_handler.ErrorMessage([]string{enum.InternalServerError.String()}))
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
