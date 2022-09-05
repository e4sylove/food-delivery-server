package userstorage

import (
	"context"
	"food_delivery/common"
)


func (storage *SQLStorage) GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error) {

	var result []common.SimpleUser

	if err := storage.db.
		Table(common.SimpleUser{}.TableName()).
		Where("id in (?)", ids).
		Find(&result).Error; err != nil {
			return nil, common.ErrDB(err)
	}

	return result, nil
} 