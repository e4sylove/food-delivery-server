package restaurantstorage

import (
	"context"
	common2 "food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"

	"gorm.io/gorm"
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
		if err == gorm.ErrRecordNotFound {
			return nil, common2.RecordNotFound
		}

		return nil, common2.ErrDB(err)
	}

	return &result, nil
}
