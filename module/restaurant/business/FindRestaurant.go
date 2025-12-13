package business

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
	"errors"
)

func (biz *restaurantBusiness) FindRestaurant(ctx context.Context, id int) *restaurantmodel.Restaurant {
	if id <= 0 {
		panic(common.ErrInvalidRequest(errors.New("id must be integer")))
	}
	restaurant, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.ResourceNotFound {
			panic(common.ResourceNotFound)
		}
	}
	if restaurant.Status == 0 {
		panic(common.ErrInvalidRequest(errors.New("data has already been marked")))
	}
	return restaurant

}
