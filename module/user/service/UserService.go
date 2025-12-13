package userservice

import (
	"FoodDelivery/common"
	usermodel "FoodDelivery/module/user/model"
	"context"
)

type userRepository interface {
	Create(ctx context.Context, user *usermodel.User) error
}
type userService struct {
	store userRepository
}

func UserService(store userRepository) *userService {
	return &userService{store: store}
}

func (biz *userService) CreateUser(ctx context.Context, dto *usermodel.UserDTO) {
	if err := dto.Validate(); err != nil {
		panic(common.ErrInvalidRequest(err))
	}
	var userEntity usermodel.User
	userEntity = usermodel.User{
		Name:    dto.Name,
		Address: dto.Address,
		Email:   dto.Email,
	}

	if err := biz.store.Create(ctx, &userEntity); err != nil {
		panic(common.ErrCreateNewEntity(usermodel.EntityName, err))
	}
}
