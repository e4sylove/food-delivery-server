package restaurantstorage

import (
	"context"
	"food_delivery/modules/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

func (storage *SQLStorage) ListDataByCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) error {
	


	return nil
}