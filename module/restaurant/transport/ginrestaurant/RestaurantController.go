package ginrestaurant

import (
	"FoodDelivery/common"
	appctx "FoodDelivery/component"
	"FoodDelivery/module/restaurant/business"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	restaurantstorage "FoodDelivery/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRestaurant Controller
func CreateRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSQLStore(ctx.GetMaiDBConnection())
		biz := business.RestaurantBusiness(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

func DeleteRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Id must be number!",
			})
			return
		}
		store := restaurantstorage.NewSQLStore(ctx.GetMaiDBConnection())
		biz := business.RestaurantBusiness(store)
		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

func FindAll(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := restaurantstorage.NewSQLStore(ctx.GetMaiDBConnection())
		biz := business.RestaurantBusiness(store)
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		pagingData.Fulfill()
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		res, err := biz.FindAllRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(res, pagingData, filter))
	}
}
