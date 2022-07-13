package restaurantservice

import (
	"context"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestauranService struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantService(store CreateRestaurantStore) *createRestauranService {
	return &createRestauranService{ store: store }
}

func (service *createRestauranService) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	err := service.store.Create(ctx, data)
	return err
}