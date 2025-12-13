package ginrestaurant

import (
	"FoodDelivery/common"
	appctx "FoodDelivery/component"
	"FoodDelivery/module/restaurant/business"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	restaurantstorage "FoodDelivery/module/restaurant/storage"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RestaurantCreateService interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate)
}
type RestaurantFindService interface {
	FindRestaurant(ctx context.Context, id int) *restaurantmodel.Restaurant
}
type RestaurantUpdateService interface {
	UpdateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantDTO) *restaurantmodel.ResRestaurant
}
type RestaurantDeleteService interface {
	DeleteRestaurant(ctx context.Context, id int)
}

func DeleteRestaurantSoftDelete(service RestaurantDeleteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID invalid",
				"details": err,
			})
		}
		service.DeleteRestaurant(c.Request.Context(), id)
		c.JSON(http.StatusOK, common.SimpleSucessResponse("Delete successful"))
	}
}

func UpdateRestaurant(restaurantService RestaurantUpdateService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto restaurantmodel.RestaurantDTO
		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"details": err,
			})
		}
		result := restaurantService.UpdateRestaurant(c.Request.Context(), &dto)
		c.JSON(http.StatusOK, common.SimpleSucessResponse(result))
	}
}

func FindRestaurant(restaurantService RestaurantFindService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID must be number!",
				"details": err,
			})
		}
		restaurant := restaurantService.FindRestaurant(c.Request.Context(), id)
		if restaurant == nil {
			panic(common.ResourceNotFound.Error())
		}
		c.JSON(http.StatusOK, common.SimpleSucessResponse(restaurant))
	}
}

// func DeleteRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error":   "ID must be number!",
// 				"details": err,
// 			})
// 		}
// 		store := restaurantstorage.NewSQLStore(ctx.GetMaiDBConnection())
// 		biz := business.RestaurantBusiness(store)
// 		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
// 			panic(err)
// 		}
// 		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
// 	}
// }

// CreateRestaurant Controller
func CreateRestaurant(restaurantService RestaurantCreateService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Input data invalid!",
				"details": err.Error(),
			})
		}
		restaurantService.CreateRestaurant(c.Request.Context(), &data)
		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
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
