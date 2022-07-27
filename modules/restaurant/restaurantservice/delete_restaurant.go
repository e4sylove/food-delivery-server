package restaurantservice

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindRestaurantByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)

	SoftDelete(ctx context.Context, id int) error
}

type deleteRestaurantService struct {
	storage DeleteRestaurantStore
}

func NewDeleteRestaurantService(storage DeleteRestaurantStore) *deleteRestaurantService {
	return &deleteRestaurantService{storage: storage}
}

func (service *deleteRestaurantService) DeleteRestaurant(ctx context.Context, id int) error {

	oldData, err := service.storage.FindRestaurantByCondition(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if err := service.storage.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
