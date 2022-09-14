package restaurantlikestorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurantlike/restaurantlikemodel"

	"gorm.io/gorm"
)

func (storage *SQLStorage) FindUserLike(context context.Context, userId, restaurantId int) (
	*restaurantlikemodel.Like, error) {
	
	var data restaurantlikemodel.Like
	db := storage.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = (?) and restaurant_id = (?)", userId, restaurantId).
		First(data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, common.ErrRecordNotFound
			}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}

func (storage *SQLStorage) CheckUserLike(context context.Context, userId, restaurantId int) (bool, error) {
	var data restaurantlikemodel.Like

	db := storage.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = (?) and restaurant_id = (?)", userId, restaurantId).
		First(data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return false, common.ErrRecordNotFound
			}
		
		return false, common.ErrDB(err)
	}

	return true, nil
}