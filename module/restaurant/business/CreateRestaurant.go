package business

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

// CreateRestaurant login interface
func (biz *restaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) {

	if err := data.Validate(); err != nil {
		panic(common.ErrInvalidRequest(err))
	}
	if err := biz.store.Create(ctx, data); err != nil {
		panic(common.ErrCreateNewEntity(restaurantmodel.EntityName, err))
	}

}
