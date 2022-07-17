package restaurantservice

import (
	"context"
	"food_delivery/modules/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)


type ListRestaurantStore interface {
	ListRestaurantByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantService struct {
	store ListRestaurantStore
}

func NewListRestaurantService(store ListRestaurantStore) *listRestaurantService {
	return &listRestaurantService{ store : store}
}

func (service *listRestaurantService) ListRestaurant(
	ctx context.Context, 
	filter *restaurantmodel.Filter, 
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := service.store.ListRestaurantByCondition(ctx, nil, filter, paging)

	return result, err
}