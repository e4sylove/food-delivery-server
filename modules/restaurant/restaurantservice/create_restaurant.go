package restaurantservice

import (
	"context"
	"errors"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantService struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantService(store CreateRestaurantStore) *createRestaurantService {
	return &createRestaurantService{ store: store }
}

func (service *createRestaurantService) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	
	if data.Name == "" {
		return errors.New("restaurant name can't be blank")
	}

	err := service.store.Create(ctx, data)
	return err
}