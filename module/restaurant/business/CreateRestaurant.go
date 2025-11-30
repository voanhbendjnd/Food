package business

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
	"errors"
)

type createRestaurantStore interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBusiness struct {
	store createRestaurantStore
}

func NewCreateRestaurantBusiness(store createRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (biz *createRestaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name cannot be empty")
	}
	if err := biz.store.CreateRestaurant(ctx, data); err != nil {
		return err
	}
	return nil

}
