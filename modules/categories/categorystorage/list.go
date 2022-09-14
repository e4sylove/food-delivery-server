package categorystorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/categories/categorymodel"
)


func (storage *SQLStorage) ListCategories(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) ([]categorymodel.Categories, error){
	
	var categories []categorymodel.Categories

	db := storage.db

	db = db.Table(categorymodel.Categories{}.TableName()).
		Where(conditions).Where("status in (1)")

	
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	
	if err := db.Find(&categories).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return categories, nil
}