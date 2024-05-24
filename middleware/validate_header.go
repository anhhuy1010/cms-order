package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/anhhuy1010/cms-order/helpers/respond"
	"github.com/anhhuy1010/cms-order/helpers/util"
)

type AppHeader struct {
	Platform string `header:"X-PLATFORM"`
	Lang     string `header:"X-LANG"`
}

func ValidateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !util.ShoudBindHeader(c) {
			c.JSON(http.StatusBadRequest, respond.MissingHeader())
			c.Abort()
			return
		}

		c.Next()
	}
}
