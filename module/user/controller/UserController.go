package controller

import (
	"FoodDelivery/common"
	usermodel "FoodDelivery/module/user/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCreateServive interface {
	CreateUser(ctx context.Context, dto *usermodel.UserDTO)
}
type UserUpdateService interface {
	UpdateUser(ctx context.Context, dto *usermodel.UserDTO) *usermodel.ResUser
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

func UpdateUser(service UserUpdateService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO usermodel.UserDTO
		if err := c.ShouldBind(&userDTO); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		res := service.UpdateUser(c.Request.Context(), &userDTO)
		c.JSON(http.StatusOK, common.SimpleSucessResponse(res))
	}
}
