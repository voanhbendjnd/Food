package business

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

// CreateRestaurant login interface
func (biz *restaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCreateNewEntity(restaurantmodel.EntityName, err)
	}
	return nil

}
