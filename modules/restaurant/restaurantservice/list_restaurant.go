package restaurantservice

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type listRestaurantRepo interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantService struct {
	repo listRestaurantRepo
}
 
func NewListRestaurantService(repo listRestaurantRepo) *listRestaurantService {
	return &listRestaurantService{
		repo: repo, 
	}
}

func (service *listRestaurantService) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := service.repo.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
