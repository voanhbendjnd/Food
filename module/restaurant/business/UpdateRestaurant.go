package business

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

func (biz *restaurantBusiness) UpdateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantDTO) *restaurantmodel.ResRestaurant {
	res, err := biz.store.Update(ctx, data)
	if err != nil {
		if err == common.ResourceNotFound {
			panic(common.ErrEntityNotFound("restaurant", err))
		}
		panic(common.ErrInternal(err))
	}
	return res
}
