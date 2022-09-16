package restaurantlikeservice

import (
	"context"
	"food_delivery/modules/restaurantlike/restaurantlikemodel"
)

type UserDislikeRestaurantStore interface{
	Delete(context context.Context, userId, restaurantId int) error
	FindUserLike(context context.Context, userId, restaurantId int) (*restaurantlikemodel.Like, error)
}

type userDislikeRestaurantService struct {
	store UserDislikeRestaurantStore
}

func NewUserDislikeRestaurantService(store UserDislikeRestaurantStore) *userDislikeRestaurantService {
	return &userDislikeRestaurantService{store: store}
}

func (service *userDislikeRestaurantService) DislikeRestaurant(ctx context.Context, userId, restaurantId int) error {

	oldData, err := service.store.FindUserLike(ctx, userId, restaurantId)

	if oldData == nil {
		return restaurantlikemodel.ErrCannotDidNotlikeRestaurant(err)
	}

	err = service.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	return nil
}