package restaurantservice

import (
	"context"
	"errors"
	"food_delivery/modules/restaurant/restaurantmodel"
)


type DeleteRestaurantStore interface {
	FindRestaurantByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)

	SoftDelete(ctx context.Context, id int) (error)
}

type deleteRestaurantService struct {
	storage DeleteRestaurantStore
}

func NewDeleteRestaurantService(storage DeleteRestaurantStore) *deleteRestaurantService {
	return &deleteRestaurantService{ storage: storage}
}

func (service *deleteRestaurantService) DeleteRestaurant(ctx context.Context, id int) (error) {
	
	oldData, err := service.storage.FindRestaurantByCondition(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}
	
	if err := service.storage.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}	