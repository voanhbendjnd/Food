package middleware

import (
	"FoodDelivery/common"
	appctx "FoodDelivery/component"
	"errors" // Cần import package errors
	"fmt"    // Cần import package fmt

	"github.com/gin-gonic/gin"
)

func Recover(app appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// Đổi tên biến err thành r (recover value)
			if r := recover(); r != nil {
				c.Header("Content-type", "application/json")

				// 1. Xử lý Lỗi Nghiệp vụ (*common.AppError)
				if appErr, ok := r.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return // ⬅️ KẾT THÚC, KHÔNG PANIC LẶP LẠI
				}

				// 2. Xử lý Lỗi Chuẩn (std error)
				if stdErr, ok := r.(error); ok {
					// Bao gói lỗi chuẩn thành 500
					appErr := common.ErrInternal(stdErr)
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return // ⬅️ KẾT THÚC, KHÔNG PANIC LẶP LẠI
				}

				// 3. Xử lý các Panic Khác (string, int, v.v.)
				// Chuyển đổi mọi thứ sang string và tạo lỗi 500
				appErr := common.ErrInternal(errors.New(fmt.Sprint(r)))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				return // ⬅️ KẾT THÚC, KHÔNG PANIC LẶP LẠI
			}
		}()
		c.Next()
	}
}
