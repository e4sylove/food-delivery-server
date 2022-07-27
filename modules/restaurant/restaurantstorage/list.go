package restaurantstorage

import (
	"context"
	common2 "food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

func (storage *SQLStorage) ListRestaurantByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common2.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	var result []restaurantmodel.Restaurant
	db := storage.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.CityID > 0 {
			db = db.Where("city_id = ?", v.CityID)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common2.ErrDB(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id DESC").
		Find(&result).Error; err != nil {
		return nil, common2.ErrDB(err)
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common2.ErrDB(err)
	}

	return result, nil
}
