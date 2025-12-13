package userservice

import (
	"FoodDelivery/common"
	usermodel "FoodDelivery/module/user/model"
	"context"

	"gorm.io/gorm"
)

type userRepository interface {
	Create(ctx context.Context, user *usermodel.User) error
	Update(ctx context.Context, user *usermodel.User) (*usermodel.User, error)
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

func (biz *userService) UpdateUser(ctx context.Context, dto *usermodel.UserDTO) *usermodel.ResUser {
	// if err := dto.Validate(); err != nil {
	// 	panic(common.ErrInvalidRequest(err))
	// }

	userEntity := usermodel.User{
		Id:      dto.Id,
		Name:    dto.Name,
		Address: dto.Address,
	}
	user, err := biz.store.Update(ctx, &userEntity)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(common.ErrEntityNotFound(usermodel.EntityName, err))
		}
		panic(common.ErrInternal(err))
	}
	result := usermodel.ResUser{
		Id:      user.Id,
		Email:   user.Email,
		Name:    user.Name,
		Address: user.Address,
	}
	return &result

}
