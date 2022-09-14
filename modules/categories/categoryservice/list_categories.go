package categoryservice

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/categories/categorymodel"
)


type ListCategoriesStore interface {
	ListCategories(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) ([]categorymodel.Categories, error)
}

type listCategoriesService struct {
	store ListCategoriesStore
}


func NewCategoriesService(store ListCategoriesStore) *listCategoriesService {
	return &listCategoriesService{store: store}
}


func (service *listCategoriesService) ListCategories(ctx context.Context) ([]categorymodel.Categories, error) {
	categories, err:= service.store.ListCategories(ctx, nil) 

	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}

	return categories, nil
}