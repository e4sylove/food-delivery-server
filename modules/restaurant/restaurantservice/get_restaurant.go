package restaurantservice

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindRestaurantByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type getRestaurantService struct {
	store GetRestaurantStore
}

func NewGetRestaurantService(store GetRestaurantStore) *getRestaurantService {
	return &getRestaurantService{store: store}
}

func (service *getRestaurantService) GetRestaurantService(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {

	data, err := service.store.FindRestaurantByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return data, err
}
