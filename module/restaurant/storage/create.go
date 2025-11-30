package restaurantstorage

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

// CreateRestaurant Kế thừa từ bên business
func (s *sqlStore) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
