package business

import (
	"context"
	"errors"
)

func (biz *restaurantBusiness) DeleteRestaurant(ctx context.Context, id int) error {

	if id <= 0 {
		return errors.New("id must be integer!")
	}
	restaurant, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if restaurant.Status == 0 {
		return errors.New("data has been deleted")
	}
	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return err
	}
	return nil
}
