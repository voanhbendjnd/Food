package business

import (
	"FoodDelivery/common"
	"context"
	"errors"
)

func (biz *restaurantBusiness) DeleteRestaurant(ctx context.Context, id int) {

	if id <= 0 {
		panic(common.ErrInvalidRequest(errors.New("id must be a positive integer")))
	}
	restaurant, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.ResourceNotFound {
			panic(common.ResourceNotFound)
		}
	}
	if restaurant.Status == 0 {
		panic(common.ErrInvalidRequest(errors.New("data has already been marked as deleted")))
	}
	if err := biz.store.SoftDelete(ctx, id); err != nil {
		panic(common.ErrInternal(err))
	}
	// return nil
}
