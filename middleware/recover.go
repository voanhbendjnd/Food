package middleware

import (
	"FoodDelivery/common"
	appctx "FoodDelivery/component"

	"github.com/gin-gonic/gin"
)

func Recover(app appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// panic(err)
					return
				}
				appErr := common.ErrINternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// panic(err)
				return
			}
		}()
		c.Next()
	}
}
