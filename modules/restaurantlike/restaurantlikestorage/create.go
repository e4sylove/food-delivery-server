package restaurantlikestorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurantlike/restaurantlikemodel"
)


func (storage *SQLStorage) Create(ctx context.Context, data *restaurantlikemodel.Like) error {

	db := storage.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Create(data).Error; err != nil {
			return common.ErrDB(err)
	}
	
	return nil
} 