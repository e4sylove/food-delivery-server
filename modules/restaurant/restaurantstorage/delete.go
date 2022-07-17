package restaurantstorage

import (
	"context"
	"food_delivery/modules/restaurant/restaurantmodel"
)

func (storage *SQLStorage) SoftDelete(
	ctx context.Context,
	id int,
) (error) {
	db := storage.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{ "status": 0 }).
		Error; err != nil {
			
		return err
	}

	return nil
}