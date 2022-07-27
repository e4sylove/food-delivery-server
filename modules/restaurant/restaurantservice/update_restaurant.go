package restaurantservice

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindRestaurantByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantService struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantService(store UpdateRestaurantStore) *updateRestaurantService {
	return &updateRestaurantService{store: store}
}

func (service *updateRestaurantService) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {

	oldData, err := service.store.FindRestaurantByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if err := service.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
