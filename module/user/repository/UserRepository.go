package userrepository

import (
	"FoodDelivery/common"
	usermodel "FoodDelivery/module/user/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (s *sqlStore) Create(ctx context.Context, user *usermodel.User) error {
	check, err := s.ExistsByEmail(ctx, user.Email)
	if err != nil {
		return common.ErrDB(err)
	}
	if check {
		return common.ErrCreateNewEntity(usermodel.EntityName, errors.New("email already exists"))
	}
	if err := s.db.Create(&user).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var cnt int64
	if err := s.db.WithContext(ctx).Table(usermodel.User{}.TableName()).Where("email = ?", email).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (s *sqlStore) Update(ctx context.Context, user *usermodel.User) (*usermodel.User, error) {
	userUpdate := s.db.WithContext(ctx).Where("id = ?", user.Id).Updates(user)
	if userUpdate.Error != nil {
		return nil, userUpdate.Error
	}
	var userDB usermodel.User
	if err := s.db.WithContext(ctx).Where("id = ?", user.Id).First(&userDB).Error; err != nil {
		return nil, err
	}
	return &userDB, nil

}
