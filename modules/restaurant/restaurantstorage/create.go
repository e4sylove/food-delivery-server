package restaurantstorage

import (
	"context"
	"food_delivery/modules/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

func (storage *SQLStorage) Create(
	ctx context.Context, 
	data *restaurantmodel.RestaurantCreate) error {
	
	db := storage.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}