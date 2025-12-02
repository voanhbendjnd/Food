package business

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

func (biz *restaurantBusiness) FindAllRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindAll(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil

}
