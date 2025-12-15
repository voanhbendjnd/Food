package controller

import (
	"FoodDelivery/common"
	usermodel "FoodDelivery/module/user/model"
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserCreateServive interface {
	CreateUser(ctx context.Context, dto *usermodel.UserDTO)
}
type UserUpdateService interface {
	UpdateUser(ctx context.Context, dto *usermodel.UserDTO) *usermodel.ResUser
}
type UserFindByIDService interface {
	FindById(ctx context.Context, id int) *usermodel.ResUser
}
type FetchAllService interface {
	FetchAll(ctx context.Context, filter *usermodel.Filter, paging *common.Paging, moreKey ...string) []usermodel.ResUser
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

func FindUserByID(service UserFindByIDService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(errors.New("ID must be a positive integer")))
		}
		res := service.FindById(c.Request.Context(), id)
		c.JSON(200, common.SimpleSucessResponse(res))
	}
}

func FetchAll(service FetchAllService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		pagingData.Fulfill()
		var filter usermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		res := service.FetchAll(c.Request.Context(), &filter, &pagingData)
		c.JSON(http.StatusOK, common.NewSuccessResponse(res, pagingData, filter))
	}
}
