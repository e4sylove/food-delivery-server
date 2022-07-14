package restaurantstorage

import (
	"context"
	"food_delivery/modules/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

func (storage *SQLStorage) ListDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	
	var result []restaurantmodel.Restaurant
	db := storage.db
	
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)
	
	if v := filter; v != nil {
		if v.CityID > 0 {
			db = db.Where("city_id = ?", v.CityID)
		}
	}
	
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id DESC").
		Find(&result).Error; err != nil {
			return nil, err
		}

	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	 
	return result, nil
}