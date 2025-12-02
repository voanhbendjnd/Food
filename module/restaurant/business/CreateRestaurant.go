package business

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
	"errors"
)

// CreateRestaurant login interface
func (biz *restaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name cannot be empty")
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil

}
