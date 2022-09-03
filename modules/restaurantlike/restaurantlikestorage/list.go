package restaurantlikestorage

import (
	"context"
	"fmt"
	"food_delivery/modules/restaurantlike/restaurantlikemodel"
)


func (storage *SQLStorage) ListRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {

	result := make(map[int]int)

	type sqlData struct {
		RestaurantId 		int	`json:"restaurant_id" gorm:"column:restaurant_id"`
		LikeCount 			int `json:"count" gorm:"column:count"`
	}

	var listLike []sqlData
	if err := storage.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id", "count(restaurant_id) as count").
		Where("restaurant_id in ?", ids). 
		Group("restaurant_id").
		Find(&listLike).Error; err != nil {
			return nil, err
		}
	
	fmt.Println(listLike)
	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
} 