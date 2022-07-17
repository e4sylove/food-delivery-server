package restaurantservice

import (
	"context"
	"errors"
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

func NewGetRestaurantService(store GetRestaurantStore) * getRestaurantService {
	return &getRestaurantService{ store: store }
}

func (service *getRestaurantService) GetRestaurantService(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	
	data, err := service.store.FindRestaurantByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	if data.Status == 0 {
		return nil, errors.New("data has been deleted")
	}
	
	return data, err
}