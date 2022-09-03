package restaurantservice

import (
	"context"
	"food_delivery/common"
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

type LikeStore interface {
	ListRestaurantLikes (ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantService struct {
	store ListRestaurantStore
	likeStore LikeStore
}
 
func NewListRestaurantService(store ListRestaurantStore, likeStore LikeStore) *listRestaurantService {
	return &listRestaurantService{store: store, likeStore: likeStore}
}

func (service *listRestaurantService) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := service.store.ListRestaurantByCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i, restaurant := range result {
		ids[i] = restaurant.Id
	}

	mapResLike, err := service.likeStore.ListRestaurantLikes(ctx, ids)

	if v := mapResLike; v != nil {
		for i, item := range result {
			result[i].LikedCount = mapResLike[item.Id]
		}
	}

	return result, nil
}
