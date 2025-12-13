package business

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

// RestaurantRepository
type restaurantRepository interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
	SoftDelete(ctx context.Context, id int) error
	FindWithCondition(ctx context.Context, cdt map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	FindAll(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
	Update(ctx context.Context, data *restaurantmodel.RestaurantDTO) (*restaurantmodel.ResRestaurant, error)
}

// (Cầu nối)
type restaurantBusiness struct {
	store restaurantRepository
}

// NewCreateRestaurantBusiness use interface
func RestaurantBusiness(store restaurantRepository) *restaurantBusiness {
	return &restaurantBusiness{store: store}
}
