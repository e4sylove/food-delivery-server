package restaurantlikemodel

import (
	"fmt"
	"food_delivery/common"
	"time"
)


type Like struct {
	RestaurantId 	int `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId 			int `json:"user_id" gorm:"column:user_id"`
	CreatedAt 		*time.Time `json:"created_at" gorm:"column:created_at"`
	User			*common.SimpleUser `json:"user" gorm:"preload:false;"`
}


func (like Like) TableName() string{
	return "restaurant_likes"
}

func (like *Like) GetRestaurantId() int {
	return like.RestaurantId
}

func (like *Like) GetUserId() int {
	return like.UserId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotDislikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot dislike this restaurant"),
		fmt.Sprintf("ErrCannotDislikeRestaurant"),
	)
}

func ErrCannotDidNotlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("You have not liked this restaurant"),
		fmt.Sprintf("ErrCannotDidNotlikeRestaurant"),
	)
}

func ErrUserAlreadyLikedRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("you've already liked this restaurant"),
		fmt.Sprintf("ErrUserAlreadyLikedRestaurant"),
	)
}
