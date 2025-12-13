package restaurantstorage

import (
	"FoodDelivery/common"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (s *sqlStore) Update(ctx context.Context, dto *restaurantmodel.RestaurantDTO) (*restaurantmodel.ResRestaurant, error) {
	var entity restaurantmodel.Restaurant
	if err := s.db.WithContext(ctx).Where("id = ?", dto.Id).First(&entity).Error; err != nil {
		return nil, err
	}
	entity = restaurantmodel.Restaurant{
		Name:        dto.Name,
		Address:     dto.Address,
		PhoneNumber: dto.PhoneNumber,
	}
	restaurantUpdate := s.db.WithContext(ctx).Where("id = ?", dto.Id).Updates(&entity)
	if restaurantUpdate.Error != nil {
		return nil, restaurantUpdate.Error
	}
	if restaurantUpdate.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var res restaurantmodel.Restaurant
	if err := s.db.WithContext(ctx).Where("id = ?", dto.Id).First(&res).Error; err != nil {
		return nil, err
	}
	lastRes := restaurantmodel.ResRestaurant{
		Id:          res.Id,
		Name:        res.Name,
		Address:     res.Address,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
		PhoneNumber: res.PhoneNumber,
		Type:        res.Type,
		Rating:      res.Rating,
	}
	return &lastRes, nil
}

// CreateRestaurant Kế thừa từ bên business (repository)
func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
func (s *sqlStore) SoftDelete(ctx context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

// *restaurant.Restaurant vì muốn không tìm thấy gì thì mình cần return nil, nếu không có thì sẽ return ra struct rỗng nhưng vẫn chứa dữ liệu
func (s *sqlStore) FindWithCondition(
	ctx context.Context,
	cdt map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(cdt).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ResourceNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}

func (s *sqlStore) FindAll(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var data []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	offset := (paging.Page - 1) * paging.Limit
	if err := db.Offset(offset).Limit(paging.Limit).Order("id desc").Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
