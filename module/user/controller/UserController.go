package controller

import (
	"FoodDelivery/common"
	usermodel "FoodDelivery/module/user/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCreateServive interface {
	CreateUser(ctx context.Context, dto *usermodel.UserDTO)
}

func CreateUser(service UserCreateServive) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO usermodel.UserDTO
		if err := c.ShouldBind(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Data user invalid",
				"details": err,
			})
		}
		service.CreateUser(c.Request.Context(), &userDTO)
		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
