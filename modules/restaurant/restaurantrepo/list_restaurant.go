package restaurantrepo

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

type UserStore interface {
	GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error)
}

type LikeStore interface {
	ListRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store ListRestaurantStore
	userStore UserStore
	likeStore LikeStore
}
 
func NewListRestaurantRepo(store ListRestaurantStore, userStore UserStore, likeStore LikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store: store, 
		userStore: userStore,
		likeStore: likeStore,
	}
}

func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	restaurants, err := repo.store.ListRestaurantByCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	userIds := make([]int, len(restaurants))
	for i, restaurant := range restaurants {
		userIds[i] = restaurant.Owner_Id
	}

	users, err := repo.userStore.GetUsers(ctx, userIds)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	// for i := range result {
	// 	for j := range users {
	// 		if result[i].UserId == users[j].Id {
	// 			result[i].User = &users[j]
	// 			break
	// 		}
	// 	}
	// }

	mapUser := make(map[int]*common.SimpleUser)

	for i, user := range users {
		mapUser[user.Id] = &users[i]
	}

	for i, restaurant := range restaurants {
		restaurants[i].Owner = mapUser[restaurant.Owner_Id]
	}

	restaurantIds := make([]int, len(restaurants))

	for i, restaurant := range restaurants {
		restaurantIds[i] = restaurant.Id
	}

	mapResLike, err := repo.likeStore.ListRestaurantLikes(ctx, restaurantIds)

	if err != nil {
		return restaurants, nil
	}

	if v := mapResLike; v != nil {
		for i, item := range restaurants {
			restaurants[i].LikedCount = mapResLike[item.Id]
		}
	}

	return restaurants, nil
}
