package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/matching-app-api/api/config"
	"github.com/hiroki-Fukumoto/matching-app-api/api/enum"
	"github.com/hiroki-Fukumoto/matching-app-api/api/error_handler"
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

func CheckApiToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		messages := []string{enum.TokenFormatError.String()}

		authorization := c.Request.Header["Authorization"]
		if authorization == nil {
			err := error_handler.ApiErrorHandle(error_handler.ErrForbidden, error_handler.ErrorMessage(messages))
			c.AbortWithStatusJSON(err.Status, err)
			return
		}
		slice := strings.Split(authorization[0], " ")
		if slice[0] != "Bearer" {
			err := error_handler.ApiErrorHandle(error_handler.ErrForbidden, error_handler.ErrorMessage(messages))
			c.AbortWithStatusJSON(err.Status, err)
			return
		}
		jwt := slice[1]
		if jwt == "" {
			err := error_handler.ApiErrorHandle(error_handler.ErrForbidden, error_handler.ErrorMessage(messages))
			c.AbortWithStatusJSON(err.Status, err)
			return
		}

		auth, err := util.ParseToken(jwt)
		if err != nil {
			messages = []string{enum.TokenFormatError.String()}
			if strings.Contains(err.Error(), "expired") {
				messages = []string{enum.ExpiredToken.String()}
			}
			err := error_handler.ApiErrorHandle(error_handler.ErrForbidden, error_handler.ErrorMessage(messages))
			c.AbortWithStatusJSON(err.Status, err)
			return
		}

		ur := repository.NewUserRepository(config.Connect())

		loginUser, err := ur.FindByID(auth.UserID)

		c.Set("AuthorizedUser", loginUser)
		c.Next()
	}
}
