package restaurantservice

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantService struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantService(store CreateRestaurantStore) *createRestaurantService {
	return &createRestaurantService{store: store}
}

func (service *createRestaurantService) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}

	err := service.store.Create(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil	
}
