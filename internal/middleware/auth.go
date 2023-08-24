package middleware

import (
	"github.com/gin-gonic/gin"
	"julo-test/internal/dto"
	"julo-test/internal/model"
	"julo-test/pkg/helper"
	"net/http"
	"regexp"
)

func BearerToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		if len(header) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Common{
				Status: "failed",
				Data:   dto.Error{Error: "Unauthenticated"},
			})
			return
		}

		rep := regexp.MustCompile(`(Token)\s?`)
		bearerStr := rep.ReplaceAllString(header[0], "")

		token := helper.GetTokenByBearer(bearerStr)
		if token == (model.Token{}) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Common{
				Status: "failed",
				Data:   dto.Error{Error: "Unauthenticated"},
			})
			return
		}

		wallet, err := helper.OpenWalletFile(token.CustomerXid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Common{
				Status: "failed",
				Data:   dto.Error{Error: "Unauthenticated"},
			})
			return
		}

		c.Set("wallet", wallet)

		c.Next()
		return
	}
}
