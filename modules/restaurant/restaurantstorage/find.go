package restaurantstorage

import (
	"context"
	"food_delivery/modules/restaurant/restaurantmodel"
)


func (storage *SQLStorage) FindRestaurantByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	
	var result restaurantmodel.Restaurant

	db := storage.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}