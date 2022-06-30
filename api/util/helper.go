package util

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
)

// ミドルウェアからセットされたログイン中のユーザー情報を取得
func GetLoginUser(c *gin.Context) (user *model.User, err error) {
	authorizedUser, ok := c.Get("AuthorizedUser")
	if !ok {
		return nil, error_handler.ErrUnauthorized
	}

	return authorizedUser.(*model.User), nil
}
