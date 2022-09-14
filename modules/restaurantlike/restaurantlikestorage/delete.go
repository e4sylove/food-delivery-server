package restaurantlikestorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurantlike/restaurantlikemodel"
)


func (storage *SQLStorage) Delete(context context.Context, userId, restaurantId int) error {

	db := storage.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).Where("restaurant_id = (?) and user_id = (?)", restaurantId, userId).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}