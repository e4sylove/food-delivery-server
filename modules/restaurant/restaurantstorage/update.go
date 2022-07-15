package restaurantstorage

import (
	"context"
	"food_delivery/modules/restaurant/restaurantmodel"
)

func (storage *SQLStorage) UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	
	db := storage.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {

		return err
	}
	
	return nil
}