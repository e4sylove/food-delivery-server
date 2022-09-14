package restaurantlikeservice

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurantlike/restaurantlikemodel"
)

type UserLikeRestaurantStore interface{
	Create(context context.Context, data *restaurantlikemodel.Like) error
	CheckUserLike(context context.Context, userId, restaurantId int) (bool, error)
}

type UserLikeRestaurantService struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantService(store UserLikeRestaurantStore) *UserLikeRestaurantService {
	return &UserLikeRestaurantService{store: store}
}

func (service *UserLikeRestaurantService) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {

	liked, err := service.store.CheckUserLike(ctx, data.UserId, data.RestaurantId)

	if err != nil && err != common.ErrRecordNotFound {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	if liked {
		return restaurantlikemodel.ErrUserAlreadyLikedRestaurant(nil)
	}	

	err = service.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}